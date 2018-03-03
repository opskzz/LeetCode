package tree

import (
	"container/list"
	"math"
)

// BinaryTreeNode 二叉树节点结构
type BinaryTreeNode struct {
	data   interface{}     // 数据域
	parent *BinaryTreeNode // 父节点
	lChild *BinaryTreeNode // 左孩子
	rChild *BinaryTreeNode // 右孩子
	height int             // 以该节点为根的子树的高度
	size   int             // 该节点的子孙数(包括节点本身)
}

// NewBinaryTreeNode 初始化二叉树节点
func NewBinaryTreeNode(e interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{data: e, size: 1}
}

// GetData 获取节点数据
func (btn *BinaryTreeNode) GetData() interface{} {
	if btn == nil {
		return nil
	}
	return btn.data
}

// SetData 设置节点数据
func (btn *BinaryTreeNode) SetData(e interface{}) {
	btn.data = e
}

// GetSize 取以该结点为根的树的结点数
func (btn *BinaryTreeNode) GetSize() int {
	return btn.size
}

// SetSize 更新当前结点及其祖先的子孙数
func (btn *BinaryTreeNode) SetSize() {
	btn.size = 1 //初始化为1,结点本身
	if btn.HasLChild() {
		btn.size += btn.GetLChild().GetSize() //加上左子树规模
	}
	if btn.HasRChild() {
		btn.size += btn.GetRChild().GetSize() //加上右子树规模
	}

	if btn.HasParent() {
		btn.parent.SetSize() //递归更新祖先的规模
	}

}

// GetHeight 取结点的高度,即以该结点为根的树的高度
func (btn *BinaryTreeNode) GetHeight() int {
	return btn.height
}

// SetHeight 更新当前结点及其祖先的高度
func (btn *BinaryTreeNode) SetHeight() {
	newH := 0 //新高度初始化为0,高度等于左右子树高度加1中的大者
	if btn.HasLChild() {
		newH = int(math.Max(float64(newH), float64(1+btn.GetLChild().GetHeight())))
	}
	if btn.HasRChild() {
		newH = int(math.Max(float64(newH), float64(1+btn.GetRChild().GetHeight())))
	}
	if newH == btn.height {
		//高度没有发生变化则直接返回
		return
	}

	btn.height = newH //否则更新高度
	if btn.HasParent() {
		btn.GetParent().SetHeight() //递归更新祖先的高度
	}
}

// HasParent 判断是否有父亲
func (btn *BinaryTreeNode) HasParent() bool {
	return btn.parent != nil
}

// GetParent 获取父节点
func (btn *BinaryTreeNode) GetParent() *BinaryTreeNode {
	if !btn.HasParent() {
		return nil
	}
	return btn.parent
}

// SetParent 设置父节点
func (btn *BinaryTreeNode) SetParent(p *BinaryTreeNode) {
	btn.parent = p
}

// HasLChild 判断是否有左孩子
func (btn *BinaryTreeNode) HasLChild() bool {
	return btn.lChild != nil
}

// SetLChild 设置当前结点的左孩子,返回原左孩子
func (btn *BinaryTreeNode) SetLChild(lc *BinaryTreeNode) *BinaryTreeNode {
	oldLC := btn.lChild
	if btn.HasLChild() {
		btn.lChild.CutOffParent() // 断开当前左孩子与结点的关系
	}
	if lc != nil {
		lc.CutOffParent() // 断开lc与其父结点的关系
		btn.lChild = lc   // 确定父子关系
		lc.SetParent(btn)
		btn.SetHeight() // 更新当前结点及其祖先高度
		btn.SetSize()   // 更新当前结点及其祖先规模
	}
	return oldLC
}

// GetLChild 获得左孩子节点
func (btn *BinaryTreeNode) GetLChild() *BinaryTreeNode {
	if !btn.HasLChild() {
		return nil
	}
	return btn.lChild
}

// HasRChild 判断是否有右孩子
func (btn *BinaryTreeNode) HasRChild() bool {
	return btn.rChild != nil
}

// SetRChild 设置当前结点的右孩子,返回原右孩子
func (btn *BinaryTreeNode) SetRChild(rc *BinaryTreeNode) *BinaryTreeNode {
	oldRC := btn.rChild
	if btn.HasRChild() {
		btn.rChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if rc != nil {
		rc.CutOffParent() //断开rc与其父结点的关系
		btn.rChild = rc   //确定父子关系
		rc.SetParent(btn)
		btn.SetHeight() //更新当前结点及其祖先高度
		btn.SetSize()   //更新当前结点及其祖先规模
	}
	return oldRC
}

// GetRChild 获得右孩子节点
func (btn *BinaryTreeNode) GetRChild() *BinaryTreeNode {
	if !btn.HasRChild() {
		return nil
	}
	return btn.rChild
}

// IsLeaf 判断是否为叶子结点
func (btn *BinaryTreeNode) IsLeaf() bool {
	return !btn.HasLChild() && !btn.HasRChild()
}

// IsLChild 判断是否为某结点的左孩子
func (btn *BinaryTreeNode) IsLChild() bool {
	return btn.HasParent() && btn == btn.parent.lChild
}

// IsRChild 判断是否为某结点的右孩子
func (btn *BinaryTreeNode) IsRChild() bool {
	return btn.HasParent() && btn == btn.parent.rChild
}

// CutOffParent 断开与父亲的关系
func (btn *BinaryTreeNode) CutOffParent() {
	if !btn.HasParent() {
		return
	}
	if btn.IsLChild() {
		btn.parent.lChild = nil // 断开该节点与父节点的连接
	} else {
		btn.parent.rChild = nil // 断开该节点与父节点的连接
	}
	btn.parent = nil       // 断开该节点与父节点的连接
	btn.parent.SetHeight() // 更新父结点及其祖先高度
	btn.parent.SetSize()   // 更新父结点及其祖先规模
}

// BinaryTree 二叉树结构
type BinaryTree struct {
	root   *BinaryTreeNode // 根结点
	height int
	size   int
}

// NewBinaryTree 初始化二叉树
func NewBinaryTree(root *BinaryTreeNode) *BinaryTree {
	return &BinaryTree{root: root}
}

//GetSize 获得二叉树总结点数
func (bt *BinaryTree) GetSize() int {
	return bt.root.size
}

//IsEmpty 判断二叉树是否为空
func (bt *BinaryTree) IsEmpty() bool {
	return bt.root != nil
}

//GetRoot 获得二叉树根节点
func (bt *BinaryTree) GetRoot() *BinaryTreeNode {
	return bt.root
}

//GetHeight 获得二叉树高度，根节点层为0
func (bt *BinaryTree) GetHeight() int {
	return bt.root.height
}

//Find 获得第一个与数据e相等的节点
func (bt *BinaryTree) Find(e interface{}) *BinaryTreeNode {
	if bt.root == nil {
		return nil
	}
	p := bt.root
	return isEqual(e, p)
}

func isEqual(e interface{}, node *BinaryTreeNode) *BinaryTreeNode {
	if e == node.GetData() {
		return node
	}
	if node.HasLChild() {
		lp := isEqual(e, node.GetLChild())
		if lp != nil {
			return lp
		}
	}
	if node.HasRChild() {
		rp := isEqual(e, node.GetRChild())
		if rp != nil {
			return rp
		}
	}
	return nil
}

//PreOrder 前序遍历，并保存在链表里
func (bt *BinaryTree) PreOrder() *list.List {
	traversal := list.New()
	preOrder(bt.root, traversal)
	return traversal
}

func preOrder(rt *BinaryTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	l.PushBack(rt)
	preOrder(rt.GetLChild(), l)
	preOrder(rt.GetRChild(), l)
}

//InOrder 中序遍历，并保存在链表里
func (bt *BinaryTree) InOrder() *list.List {
	traversal := list.New()
	inOrder(bt.root, traversal)
	return traversal
}

func inOrder(rt *BinaryTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	inOrder(rt.GetLChild(), l)
	l.PushBack(rt)
	inOrder(rt.GetRChild(), l)
}

//PostOrder 后序遍历，并保存在链表里
func (bt *BinaryTree) PostOrder() *list.List {
	traversal := list.New()
	postOrder(bt.root, traversal)
	return traversal
}

func postOrder(rt *BinaryTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	postOrder(rt.GetLChild(), l)
	postOrder(rt.GetRChild(), l)
	l.PushBack(rt)
}
