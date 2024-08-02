package tree_t

// TowTree 数节点
type TowTree struct {
	K     int // 节点存储的值
	Val   int
	Left  *TowTree // 左子树
	Right *TowTree // 右子树
}

func NewTwoTree() *TowTree {
	return &TowTree{
		K:   30,
		Val: 90,
		Left: &TowTree{
			K:   28,
			Val: 119,
			Left: &TowTree{
				K:   27,
				Val: 113,
			},
			Right: &TowTree{
				K:   26,
				Val: 101,
			},
		},
		Right: &TowTree{
			K:   31,
			Val: 114,
			Left: &TowTree{
				K:   32,
				Val: 121,
			},
			Right: &TowTree{
				K:   35,
				Val: 165,
			},
		},
	}
}

// SearchNode 查找节点
func (t *TowTree) SearchNode(k int) int {
	if t.Left != nil {
		val := t.Left.SearchNode(k)
		if val != -1 {
			return val
		}
	}

	if t.K == k {
		return t.Val
	}

	if t.Right != nil {
		val := t.Right.SearchNode(k)
		if val != -1 {
			return val
		}
	}

	return -1
}

// IterateNode 迭代查找节点
func (t *TowTree) IterateNode(k int) int {
	stack := []*TowTree{t}
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		if node.K == k {
			return node.Val
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return -1
}
