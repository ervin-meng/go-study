package tree

import (
	"fmt"
	"go-study/structure/list"
)

type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTree(data interface{}) *BinaryTree {
	return newBinaryTree(data)
}

func newBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (t *BinaryTree) AddLeftNode(data interface{}) {
	t.Left = newBinaryTree(data)
}

func (t *BinaryTree) AddRightNode(data interface{}) {
	t.Right = newBinaryTree(data)
}

func PreOrderTraversal(t *BinaryTree) {
	fmt.Println(t.Data)

	if t.Left != nil {
		PreOrderTraversal(t.Left)
	}

	if t.Right != nil {
		PreOrderTraversal(t.Right)
	}
}

func PreOrderTraversalPro(t *BinaryTree) {
	s := list.NewStack(t)
	for !s.IsEmpty() {
		node := s.Pop().(*BinaryTree)
		fmt.Println(node.Data)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

func InOrderTraversal(t *BinaryTree) {
	if t.Left != nil {
		InOrderTraversal(t.Left)
	}

	fmt.Println(t.Data)

	if t.Right != nil {
		InOrderTraversal(t.Right)
	}
}

func PostOrderTraversal(t *BinaryTree) {
	if t.Left != nil {
		PostOrderTraversal(t.Left)
	}

	if t.Right != nil {
		PostOrderTraversal(t.Right)
	}

	fmt.Println(t.Data)
}
