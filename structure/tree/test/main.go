package main

import (
	"fmt"
	"go-study/structure/tree"
)

func main() {
	//testBinaryTree()
	testBinarySearchTree()
}

func testBinaryTree() {
	t := tree.NewBinaryTree("a")
	t.InsertLeft("b")
	t.InsertRight("c")
	t.Left.InsertLeft("d")
	//tree.PreOrderTraversal(t)
	tree.PreOrderTraversal(t)
	//tree.InOrderTraversal(t)
	//tree.PostOrderTraversal(t)
}

func testBinarySearchTree() {
	t := tree.NewBinarySearchTree(2, 1, 3, 4)
	//fmt.Println(t.Root.Data)
	//fmt.Println(t.Root.Left.Data)
	//fmt.Println(t.Root.Right.Data)
	//fmt.Println(t.Root.Right.Right.Data)
	t.Delete(2)
	fmt.Println(t.Root.Data)
	fmt.Println(t.Root.Left.Data)
	fmt.Println(t.Root.Right.Data)
}
