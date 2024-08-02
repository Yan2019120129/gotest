package tree_t

import "fmt"

// Key 接口定义了键的比较方法
type Key interface {
	Less(Key) bool // 判断当前键是否小于给定键
	Eq(Key) bool   // 判断当前键是否等于给定键
}

// Node 结构体定义了B+树的节点
type Node struct {
	isLeaf   bool    // 是否是叶子节点
	keys     []Key   // 键
	children []*Node // 子节点
	next     *Node   // 叶子节点的下一个节点
}

// BTree 结构体定义了B+树
type BTree struct {
	degree int   // 树的度
	root   *Node // 树的根节点
}

// NewBTree 函数创建一个新的B+树
func NewBTree(degree int) *BTree {
	return &BTree{
		degree: degree,
		root:   &Node{isLeaf: true},
	}
}

// Insert 方法插入一个键到B+树中
func (t *BTree) Insert(k Key) {
	root := t.root
	if len(root.keys) == (2*t.degree - 1) {
		temp := &Node{}
		t.root = temp
		temp.children = append(temp.children, root)
		t.splitChild(temp, 0)
		t.insertNonFull(temp, k)
	} else {
		t.insertNonFull(root, k)
	}
}

// insertNonFull 方法在一个非满节点中插入一个键
func (t *BTree) insertNonFull(x *Node, k Key) {
	i := len(x.keys) - 1
	if x.isLeaf {
		// 如果是叶子节点，直接插入键
		x.keys = append(x.keys, nil)
		x.children = append(x.children, nil)
		for i >= 0 && k.Less(x.keys[i]) {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
	} else {
		// 如果是内部节点，找到合适的子节点递归插入
		for i >= 0 && k.Less(x.keys[i]) {
			i--
		}
		i++
		if len(x.children[i].keys) == (2*t.degree - 1) {
			// 如果子节点满了，先分裂
			t.splitChild(x, i)
			if k.Eq(x.keys[i]) {
				i++
			}
		}
		t.insertNonFull(x.children[i], k)
	}
}

// splitChild 方法分裂一个满的子节点
func (t *BTree) splitChild(x *Node, i int) {
	tDegree := t.degree
	y := x.children[i]
	z := &Node{isLeaf: y.isLeaf}
	x.children = append(x.children, nil)
	copy(x.children[i+2:], x.children[i+1:])
	x.children[i+1] = z
	x.keys = append(x.keys, nil)
	copy(x.keys[i+1:], x.keys[i:])
	x.keys[i] = y.keys[tDegree-1]
	z.keys = append(z.keys, y.keys[tDegree:]...)
	y.keys = y.keys[:tDegree-1]
	if !y.isLeaf {
		z.children = append(z.children, y.children[tDegree:]...)
		y.children = y.children[:tDegree]
	}
}

// Print 方法打印B+树的所有键
func (t *BTree) Print() {
	t.printInOrder(t.root, 0)
}

// printInOrder 方法按顺序打印树的所有键
func (t *BTree) printInOrder(x *Node, l int) {
	fmt.Printf("Level \"%v\" ", l)
	if x == nil {
		fmt.Println("NIL")
	} else {
		fmt.Printf("keys: ")
		for _, k := range x.keys {
			fmt.Printf("%v ", k)
		}
		fmt.Println()
		if !x.isLeaf {
			l++
			for _, c := range x.children {
				t.printInOrder(c, l)
			}
		}
	}
}

// InorderTraversal 二叉树 Morris 遍历算法中序遍历 时间复杂度 O(n)，空间复杂度O(1)
func InorderTraversal(root *TowTree) (res []int) {
	for root != nil {
		if root.Left != nil {

			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {

				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {

				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root

				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)

				// 恢复原样
				predecessor.Right = nil

				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)

			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}
