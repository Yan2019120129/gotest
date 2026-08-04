package main

import (
	"reflect"
	"testing"
)

func TestFindPrimeNodesWithWorkers(t *testing.T) {
	node3 := &TreeNode{Value: 3}
	node7 := &TreeNode{Value: 7}
	node9 := &TreeNode{Value: 9}
	node11 := &TreeNode{Value: 11}
	node2 := &TreeNode{Value: 2, Children: []*TreeNode{node3, node9}}
	node4 := &TreeNode{Value: 4, Children: []*TreeNode{node7}}
	node5 := &TreeNode{Value: 5, Children: []*TreeNode{node11}}
	root := &TreeNode{Value: 1, Children: []*TreeNode{node2, node4, node5}}
	want := []*TreeNode{node2, node5, node3, node7, node11}

	for _, workers := range []int{1, 2, 8} {
		if got := FindPrimeNodesWithWorkers(root, workers); !reflect.DeepEqual(got, want) {
			t.Fatalf("workers=%d: got values %v, want %v", workers, values(got), values(want))
		}
	}
}

func TestFindPrimeNodesWithWorkersHandlesEmptyCases(t *testing.T) {
	if got := FindPrimeNodesWithWorkers(nil, 2); got != nil {
		t.Fatalf("nil tree: got %v, want nil", values(got))
	}

	root := &TreeNode{Value: 1, Children: []*TreeNode{{Value: 0}, {Value: 4}}}
	if got := FindPrimeNodesWithWorkers(root, 0); len(got) != 0 {
		t.Fatalf("non-prime tree: got values %v, want none", values(got))
	}
}

func TestBuildNaryTree(t *testing.T) {
	root := BuildNaryTree(3, 3, 7)
	if len(root.Children) != 7 {
		t.Fatalf("root child count = %d, want 7", len(root.Children))
	}

	for _, child := range root.Children {
		if len(child.Children) != 3 {
			t.Fatalf("second-level child count = %d, want 3", len(child.Children))
		}
		for _, leaf := range child.Children {
			if len(leaf.Children) != 0 {
				t.Fatalf("leaf has %d children, want 0", len(leaf.Children))
			}
		}
	}
}

func values(nodes []*TreeNode) []int {
	values := make([]int, len(nodes))
	for index, node := range nodes {
		values[index] = node.Value
	}
	return values
}
