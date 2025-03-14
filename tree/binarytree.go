package tree

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) Insert(val int) {
	t.Root = insertNode(t.Root, val)
}

func insertNode(Root *Node, val int) *Node {

	if Root == nil {
		return &Node{Data: val}
	}

	if Root.Data > val {
		Root.Left = insertNode(Root.Left, val)
	} else {
		Root.Right = insertNode(Root.Right, val)
	}

	return Root
}

func (t *BinaryTree) InOrderTraversal() {
	inorderTraversal(t.Root)
	fmt.Println()
}

func inorderTraversal(Root *Node) {
	if Root != nil {
		inorderTraversal(Root.Left)
		fmt.Print(Root.Data, " ")
		inorderTraversal(Root.Right)
	}
}
