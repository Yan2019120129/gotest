package tree_t_test

import (
	"fmt"
	"gotest/base/tree_t"
	"testing"
)

// TestSearchNode 查找节点
func TestSearchTowTree(t *testing.T) {
	tt := tree_t.NewTwoTree()
	val := tt.SearchNode(28)
	fmt.Println("val:", val)
}

// TestInorderTraversal 迭代的方式遍历二叉树
func TestInorderTraversal(t *testing.T) {
	tt := tree_t.NewTwoTree()
	val := tt.IterateNode(28)
	fmt.Println("val:", val)
}
