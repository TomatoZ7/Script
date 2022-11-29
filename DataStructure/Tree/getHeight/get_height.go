package getHeight

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type linkedNode struct {
	data *TreeNode
	next *linkedNode
}

type queue struct {
	head *linkedNode
	tail *linkedNode
	size int
}

// 广度优先
func BreadthFirst(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 树的高度
	height := 0
	// 链表节点
	rootNode := &linkedNode{ data: root }
	// 队列
	q := queue{
		head: rootNode,
		tail: rootNode,
		size: 1,
	}
	// 每一层的个数
	numPile := 1
	// 记录下一层的节点个数
	numNextPile := 0

	for q.size != 0 {
		node := q.head.data
		//fmt.Println(node, height)

		if node.left != nil {
			q.tail.next = &linkedNode{
				data: node.left,
			}
			q.tail = q.tail.next
			numNextPile++
			q.size++
		}

		if node.right != nil {
			q.tail.next = &linkedNode{
				data: node.right,
			}
			q.tail = q.tail.next
			numNextPile++
			q.size++
		}

		q.head = q.head.next
		q.size--
		numPile--

		if numPile == 0 {
			height++
			numPile = numNextPile
			numNextPile = 0
		}
	}

	return height
}

// 深度优先
func DepthFirst(root *TreeNode, height int) int {
	if root == nil {
		return height
	}

	leftTreeHeight := DepthFirst(root.left, height)
	rightTreeHeight := DepthFirst(root.right, height)

	if leftTreeHeight > rightTreeHeight {
		return leftTreeHeight + 1
	} else {
		return rightTreeHeight + 1
	}
}
