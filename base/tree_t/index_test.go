package tree_t_test

import (
	"fmt"
	"gotest/base/tree_t"
	"testing"
)

// TestSearchNode 查找节点
func TestSearchNode(t *testing.T) {
	tree := &tree_t.TreeNode{
		K: 30,
		Left: &tree_t.TreeNode{
			K:   28,
			Val: 'q',
			Left: &tree_t.TreeNode{
				K:   27,
				Val: 'w',
			},
			Right: &tree_t.TreeNode{
				K:   26,
				Val: 'e',
			},
		},
		Right: &tree_t.TreeNode{
			K:   31,
			Val: 'r',
			Left: &tree_t.TreeNode{
				K:   32,
				Val: 't',
			},
			Right: &tree_t.TreeNode{
				K:   35,
				Val: 'y',
			},
		},
	}
	i := 0
	tree_t.SearchNode(27, tree, &i)
	fmt.Println("count:", i)
}

// TestInorderTraversal 迭代的方式遍历二叉树
func TestInorderTraversal(t *testing.T) {
	sum := tree_t.InorderTraversal(&tree_t.TreeNode{
		Val: 1,
		Left: &tree_t.TreeNode{
			Val: 2,
			Left: &tree_t.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &tree_t.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &tree_t.TreeNode{
			Val: 3,
			Left: &tree_t.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &tree_t.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	})
	fmt.Println(sum)
}