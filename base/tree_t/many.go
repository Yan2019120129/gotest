package tree_t

// ManyTree 多叉树节点
type ManyTree struct {
	Val      int
	Id       int
	ParentId int
	Client   []*ManyTree
}

func NewManyTree() *ManyTree {
	return &ManyTree{
		Id:       1,
		ParentId: 0,
		Val:      0,
		Client: []*ManyTree{
			{Id: 2, ParentId: 1, Val: 2,
				Client: []*ManyTree{
					{Id: 3, ParentId: 2, Val: 3},
					{Id: 4, ParentId: 2, Val: 4},
				},
			},
			{Id: 5, ParentId: 1, Val: 5,
				Client: []*ManyTree{
					{Id: 6, ParentId: 5, Val: 6},
					{Id: 7, ParentId: 5, Val: 7},
				},
			},
			{Id: 8, ParentId: 1, Val: 8,
				Client: []*ManyTree{
					{Id: 9, ParentId: 8, Val: 9},
					{Id: 10, ParentId: 8, Val: 10},
				},
			},
		},
	}
}

// SearchNode 查找节点
func (m *ManyTree) SearchNode(k int) int {
	if m.Id == k {
		return m.Val
	}
	for _, tree := range m.Client {
		val := tree.SearchNode(k)
		if val != -1 {
			return val
		}
	}
	return -1
}

// IterateNode 查找节点
func (m *ManyTree) IterateNode(k int) int {
	stack := []*ManyTree{m}
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		if node.Id == k {
			return node.Val
		}
		stack = stack[:index]
		for _, tree := range node.Client {
			stack = append(stack, tree)
		}
	}
	return -1
}

// SelectClientNode 查找子节点
func (m *ManyTree) SelectClientNode(k int) *ManyTree {
	if m == nil {
		return nil
	}
	stack := []*ManyTree{m}
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		if node.Id == k {
			return node
		}
		for _, tree := range node.Client {
			stack = append(stack, tree)
		}
	}
	return nil
}

// ToStack 转换为盏
func (m *ManyTree) ToStack() []*ManyTree {
	stack := []*ManyTree{m}
	tmpStack := make([]*ManyTree, 0)
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		tmpStack = append(tmpStack, &ManyTree{
			Val:      node.Val,
			Id:       node.Id,
			ParentId: node.ParentId,
			Client:   nil,
		})
		for _, tree := range node.Client {
			stack = append(stack, tree)
		}
	}
	return tmpStack
}

// InserterClientNode 插入子节点
func (m *ManyTree) InserterClientNode(node *ManyTree) {
	nodeTmp := m.SelectClientNode(node.ParentId)
	if nodeTmp == nil {
		m.Client = append(m.Client, node)
		return
	}
	node.Client = append(node.Client, node)
	return
}

// ListToTree 数组转换为树
func ListToTree(rootKey int, listNode []*ManyTree) *ManyTree {
	tmp := map[int][]*ManyTree{}
	for _, tree := range listNode {
		if p, ok := tmp[tree.ParentId]; ok {
			p = append(p, tree)
			continue
		}
		tmp[tree.ParentId] = append(tmp[tree.ParentId], tree)
	}
	//rootNode := &ManyTree{}
	//if len(tmp[rootKey]) > 0 {
	//	rootNode = tmp[rootKey][0]
	//}

	return nil
}
