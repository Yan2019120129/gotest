package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestCommandRunnerRun 验证三种读取方法都能处理同一份日志。
func TestCommandRunnerRun(t *testing.T) {
	baseTime := time.Date(2026, 7, 29, 13, 32, 0, 0, time.UTC)
	logFile := writeMainTestLog(t, strings.Join([]string{
		mainAccessLogLine(baseTime, "node-a"),
		mainAccessLogLine(baseTime.Add(time.Second), "node-b"),
		mainAccessLogLine(baseTime.Add(2*time.Second), "node-a"),
	}, "\n")+"\n")

	for _, method := range []string{"custom", "custom_optimized", "traverse"} {
		t.Run(method, func(t *testing.T) {
			runner := commandRunner{config: commandConfig{
				method:    method,
				filePath:  logFile,
				startTime: baseTime.Add(-time.Second),
				endTime:   baseTime.Add(3 * time.Second),
				topK:      10,
			}}

			nodes, err := runner.Run(context.Background())
			if err != nil {
				t.Fatalf("Run 返回错误：%v", err)
			}
			actual := make(map[string]int64, len(nodes))
			for _, node := range nodes {
				actual[node.nodeID] = node.count
			}
			if actual["node-a"] != 2 || actual["node-b"] != 1 {
				t.Fatalf("统计结果 = %#v，期望 node-a:2、node-b:1", actual)
			}
		})
	}
}

// TestParseCommandConfig 验证参数解析、默认时间范围和非法参数检查。
func TestParseCommandConfig(t *testing.T) {
	now := time.Date(2026, 7, 31, 8, 0, 0, 0, time.Local)
	config, err := parseCommandConfig([]string{"-method", "custom", "-file", "test.log"}, now)
	if err != nil {
		t.Fatalf("解析默认参数失败：%v", err)
	}
	if config.startTime.Format(time.RFC3339) != "2026-07-29T13:32:00Z" || config.endTime.Format(time.RFC3339) != "2026-07-29T13:55:00Z" {
		t.Fatalf("默认时间范围错误：%s ~ %s", config.startTime, config.endTime)
	}

	invalidArguments := [][]string{
		{"-method", "unknown"},
		{"-start", "invalid"},
		{"-start", "2026-07-29T13:55:00Z", "-end", "2026-07-29T13:32:00Z"},
		{"-k", "0"},
	}
	for _, arguments := range invalidArguments {
		if _, err := parseCommandConfig(arguments, now); err == nil {
			t.Fatalf("参数 %v 未返回错误", arguments)
		}
	}
}

// writeMainTestLog 创建供命令行分派测试使用的临时日志文件。
func writeMainTestLog(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "access.log")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("写入测试日志失败：%v", err)
	}
	return path
}

// mainAccessLogLine 生成可被三种实现共同解析的访问日志行。
func mainAccessLogLine(accessTime time.Time, nodeID string) string {
	formattedTime := accessTime.Format(time.RFC3339Nano)
	return formattedTime + " INFO [access] access_time=" + formattedTime +
		" node_id=" + nodeID +
		" request_id=request-1 remote_addr=127.0.0.1 method=GET path=/health status=200 bytes=1 duration_ms=1 user_agent=test"
}
