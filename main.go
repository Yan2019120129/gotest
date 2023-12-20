package main

import "fmt"

// TreeNode 数节点
type TreeNode struct {
	k     int // 节点存储的值
	v     string
	Left  *TreeNode // 左子树
	Right *TreeNode // 右子树
}

func main() {
	tree := &TreeNode{
		k: 30,
		Left: &TreeNode{
			k: 28,
			v: "q",
			Left: &TreeNode{
				k: 27,
				v: "w",
			},
			Right: &TreeNode{
				k: 26,
				v: "e",
			},
		},
		Right: &TreeNode{
			k: 31,
			v: "r",
			Left: &TreeNode{
				k: 32,
				v: "t",
			},
			Right: &TreeNode{
				k: 35,
				v: "y",
			},
		},
	}
	i := 0
	SearchNode(27, tree, &i)
	fmt.Println("count:", i)
}

// SearchNode 查找节点
func SearchNode(val int, t *TreeNode, count *int) {
	if t == nil {
		return
	}
	*count++
	if t.k == val {
		fmt.Printf("Found Node: (%d, %s)\n", t.k, t.v)
		return
	}
	//if val < t.k {
	SearchNode(val, t.Left, count)
	//} else {
	SearchNode(val, t.Right, count)
	//}
}
