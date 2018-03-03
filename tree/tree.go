package tree

// Node 节点结构
type Node struct {
	left  *Node
	right *Node
	value int
}

// Root 根结点
type Root struct {
	root *Node
	size int
}

// BinaryTree 二叉树结构
type BinaryTree struct {
	root *Node
	size int
}

// Size 二叉树的size
func (bt *BinaryTree) Size() int {
	return bt.size
}

// Root 二叉树的根
func (bt *BinaryTree) Root() *Node {
	return bt.root
}

// NewBt 初始化二叉树
func NewBt() *BinaryTree {
	bt := new(BinaryTree)
	bt.size = 0
	return bt
}

// Insert 节点插入
func (root *Node) Insert(node *Node) {
	if node.value > root.value {
		if root.right == nil {
			root.right = node
		} else {
			root.right.Insert(node)
		}
	} else if node.value < root.value {
		if root.left == nil {
			root.left = node
		} else {
			root.left.Insert(node)
		}
	}
}

// Insert 树插入
func (bt *BinaryTree) Insert(value int) {
	if bt.root == nil {
		bt.root = &Node{nil, nil, value}
	}
	bt.size++
	bt.root.Insert(&Node{nil, nil, value})
}
