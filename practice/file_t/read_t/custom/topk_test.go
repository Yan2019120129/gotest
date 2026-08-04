package custom

import "testing"

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
