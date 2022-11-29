package tree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// 前序遍历（中-左-右）
func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// 中序遍历（左-中-右）
func InOrder(root *Node) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Println(root.data)
	InOrder(root.right)
}

// 后序遍历（左-右-中）
func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.data)
}
