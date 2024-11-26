package tree_t

import (
	"fmt"
	"testing"
)

// SearchManyTreeTest 多叉数查找值
// func BenchmarkSearchManyTree(b *testing.B) {
func TestSearchManyTree(b *testing.T) {
	manyTree := NewManyTree()
	val := manyTree.SearchNode(5)
	fmt.Printf("val：%d\n", val)
}

// TestInorderManyTree 多叉数迭代查找值
// func BenchmarkInorderManyTree(b *testing.B) {
func TestInorderManyTree(b *testing.T) {
	manyTree := NewManyTree()
	val := manyTree.IterateNode(7)
	fmt.Printf("val：%d\n", val)
}

// TestInserterManyTree 链表转换多叉树
func BenchmarkInserterManyTree(b *testing.B) {
	//func TestInserterManyTree(b *testing.T) {
	manyTree := NewManyTree().ToStack()
	fmt.Println("start")
	for i := 0; i < b.N; i++ {
		tmpMap := map[int]*ManyTree{}
		for _, tree := range manyTree {
			index := tree.Id
			tmpMap[index] = tree
		}
		rootId := 0
		for _, tree := range manyTree {
			index := tree.ParentId
			if p, ok := tmpMap[index]; ok {
				if p.ParentId == 0 {
					rootId = p.Id
				}
				p.Client = append(p.Client, tree)
			}
		}
		fmt.Println(rootId)
	}
}
