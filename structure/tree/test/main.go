package main

import (
	"go-study/structure/tree"
)

func main() {
	t := tree.NewBinaryTree("a")
	t.AddLeftNode("b")
	t.AddRightNode("c")
	t.Left.AddLeftNode("d")
	//tree.PreOrderTraversal(t)
	tree.PreOrderTraversal(t)
	//tree.InOrderTraversal(t)
	//tree.PostOrderTraversal(t)
}
