package tree_test_test

import (
	"fmt"
	"gotest/base/tree_test"
	"testing"
)

// TestSearchNode 查找节点
func TestSearchNode(t *testing.T) {
	tree := &tree_test.TreeNode{
		K: 30,
		Left: &tree_test.TreeNode{
			K: 28,
			V: "q",
			Left: &tree_test.TreeNode{
				K: 27,
				V: "w",
			},
			Right: &tree_test.TreeNode{
				K: 26,
				V: "e",
			},
		},
		Right: &tree_test.TreeNode{
			K: 31,
			V: "r",
			Left: &tree_test.TreeNode{
				K: 32,
				V: "t",
			},
			Right: &tree_test.TreeNode{
				K: 35,
				V: "y",
			},
		},
	}
	i := 0
	tree_test.SearchNode(27, tree, &i)
	fmt.Println("count:", i)
}
