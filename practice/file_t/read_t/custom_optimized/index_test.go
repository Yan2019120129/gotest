package custom_optimized

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestTopKNodesProcessEachLineExactlyOnce 验证跨越原始分片边界的日志不会重复或丢失。
func TestTopKNodesProcessEachLineExactlyOnce(t *testing.T) {
	baseTime := time.Date(2026, 7, 29, 12, 32, 0, 0, time.UTC)
	expectedCounts := map[string]int64{}
	lines := make([]string, 0, 101)
	for index := 0; index < 101; index++ {
		nodeID := "node-" + string(rune('a'+index%4))
		expectedCounts[nodeID]++
		lines = append(lines, accessLogLine(
			baseTime.Add(time.Duration(index)*time.Second).Format(time.RFC3339Nano),
			nodeID,
			strings.Repeat("x", index%41),
		))
	}
	logFile := writeTestLog(t, strings.Join(lines, "\n"))

	processor := newLogProcessor()
	fileInfo, err := os.Stat(logFile)
	if err != nil {
		t.Fatalf("获取测试文件信息失败：%v", err)
	}
	partitions, err := processor.createPartitions(logFile, fileInfo.Size())
	if err != nil {
		t.Fatalf("创建分片失败：%v", err)
	}
	assertPartitionsAreLineAligned(t, logFile, fileInfo.Size(), partitions)

	nodes, err := TopKNodes(
		context.Background(),
		logFile,
		baseTime.Add(-time.Second).Unix(),
		baseTime.Add(102*time.Second).Unix(),
		10,
	)
	if err != nil {
		t.Fatalf("TopKNodes 返回错误：%v", err)
	}
	if len(nodes) != len(expectedCounts) {
		t.Fatalf("节点数量 = %d，期望 %d", len(nodes), len(expectedCounts))
	}
	for _, node := range nodes {
		if node.Count != expectedCounts[node.NodeID] {
			t.Fatalf("节点 %s 数量 = %d，期望 %d", node.NodeID, node.Count, expectedCounts[node.NodeID])
		}
	}
}

// TestTopKNodesHandlesSmallFileAndFinalLine 验证小文件和无换行结尾的最后一行。
func TestTopKNodesHandlesSmallFileAndFinalLine(t *testing.T) {
	baseTime := time.Date(2026, 7, 29, 12, 32, 0, 0, time.UTC)
	logFile := writeTestLog(t, accessLogLine(baseTime.Format(time.RFC3339Nano), "node-a", ""))

	nodes, err := TopKNodes(
		context.Background(),
		logFile,
		baseTime.Add(-time.Second).Unix(),
		baseTime.Add(time.Second).Unix(),
		1,
	)
	if err != nil {
		t.Fatalf("TopKNodes 返回错误：%v", err)
	}
	if len(nodes) != 1 || nodes[0] != (NodeCount{NodeID: "node-a", Count: 1}) {
		t.Fatalf("统计结果 = %#v，期望 node-a:1", nodes)
	}
}

// TestTopKNodesCanceled 验证调用方取消上下文后会立即返回取消错误。
func TestTopKNodesCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := TopKNodes(ctx, "not-used.log", 0, time.Now().Unix(), 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("错误 = %v，期望 context.Canceled", err)
	}
}

// TestTopKNodesRejectsLongLine 验证异常长行不会导致无界内存分配。
func TestTopKNodesRejectsLongLine(t *testing.T) {
	logFile := writeTestLog(t, strings.Repeat("x", maxLineSize+1))

	_, err := TopKNodes(context.Background(), logFile, 0, time.Now().Unix(), 1)
	if err == nil {
		t.Fatal("超长日志行未返回错误")
	}
}

// TestSelectTopK 验证小根堆只保留正确的 Top-K，并稳定处理相同计数。
func TestSelectTopK(t *testing.T) {
	counts := map[string]NodeCount{
		"node-z": {NodeID: "node-z", Count: 10},
		"node-b": {NodeID: "node-b", Count: 8},
		"node-a": {NodeID: "node-a", Count: 8},
		"node-c": {NodeID: "node-c", Count: 7},
		"node-d": {NodeID: "node-d", Count: 6},
	}

	nodes := selectTopK(counts, 3)
	expected := []NodeCount{
		{NodeID: "node-z", Count: 10},
		{NodeID: "node-a", Count: 8},
		{NodeID: "node-b", Count: 8},
	}
	assertNodeCounts(t, nodes, expected)

	assertNodeCounts(t, selectTopK(counts, 10), []NodeCount{
		{NodeID: "node-z", Count: 10},
		{NodeID: "node-a", Count: 8},
		{NodeID: "node-b", Count: 8},
		{NodeID: "node-c", Count: 7},
		{NodeID: "node-d", Count: 6},
	})
	assertNodeCounts(t, selectTopK(nil, 1), []NodeCount{})
}

// assertPartitionsAreLineAligned 验证相邻分片连续且所有边界都位于完整日志行之间。
func assertPartitionsAreLineAligned(t *testing.T, filePath string, fileSize int64, partitions []filePartition) {
	t.Helper()
	if len(partitions) != partitionCount {
		t.Fatalf("分片数量 = %d，期望 %d", len(partitions), partitionCount)
	}
	if partitions[0].start != 0 || partitions[len(partitions)-1].end != fileSize {
		t.Fatalf("首尾分片边界不正确：%#v", partitions)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("读取测试文件失败：%v", err)
	}
	for index, partition := range partitions {
		if partition.start > partition.end {
			t.Fatalf("分片 %d 的范围无效：%#v", index, partition)
		}
		if index > 0 && partitions[index-1].end != partition.start {
			t.Fatalf("分片 %d 与前一分片不连续", index)
		}
		if partition.start > 0 && partition.start < fileSize && content[partition.start-1] != '\n' {
			t.Fatalf("分片 %d 起点未对齐换行符", index)
		}
		if partition.end > 0 && partition.end < fileSize && content[partition.end-1] != '\n' {
			t.Fatalf("分片 %d 终点未对齐换行符", index)
		}
	}
}

// assertNodeCounts 验证节点统计结果与预期完全一致。
func assertNodeCounts(t *testing.T, actual []NodeCount, expected []NodeCount) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("结果数量 = %d，期望 %d：%#v", len(actual), len(expected), actual)
	}
	for index := range expected {
		if actual[index] != expected[index] {
			t.Fatalf("第 %d 项 = %#v，期望 %#v", index, actual[index], expected[index])
		}
	}
}

// writeTestLog 创建指定内容的临时日志文件。
func writeTestLog(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "access.log")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("写入测试日志失败：%v", err)
	}
	return path
}

// accessLogLine 生成包含可变填充内容的访问日志行。
func accessLogLine(accessTime string, nodeID string, padding string) string {
	return accessTime + " INFO [access] access_time=" + accessTime +
		" node_id=" + nodeID +
		" request_id=request-1 remote_addr=127.0.0.1 method=GET path=/health status=200 bytes=1 duration_ms=1 padding=" + padding
}
