package tree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// 前序遍历（中-左-右）
func preOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	preOrder(root.left)
	preOrder(root.right)
}

// 中序遍历（左-中-右）
func inOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Println(root.data)
	inOrder(root.right)
}

// 后序遍历（左-右-中）
func postOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	inOrder(root.right)
	fmt.Println(root.data)
}
