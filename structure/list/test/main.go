package main

import (
	"fmt"
	"go-study/structure/list"
)

func main() {
	//testQueue()
	testStack()
}

func testSeqList() {
	l := list.NewSeqList(1, 2, 3)
	fmt.Println(l)
	fmt.Println(l.IsEmpty())
	l.Append("four")
	fmt.Println(l)
	err := l.Insert(3, "five")
	fmt.Println(err, l)
	err = l.Delete(1)
	fmt.Println(err, l)
	fmt.Println(l.Get(0))
}

func testLinkedList() {
	l := list.NewLinkedList(1, 2, 3)
	fmt.Println(l.IsEmpty())
	l.Append("four")
	fmt.Println(l.Scan(), l.Length())
	err := l.Insert(3, "five")
	fmt.Println(err, l.Scan())
	err = l.Delete(1)
	fmt.Println(err, l.Scan())
	err = l.Update(0, 1.1)
	fmt.Println(err, l.Scan())
	fmt.Println(l.Get(1))
}

func testQueue() {
	q := list.NewQueue(1, 2)
	fmt.Println(q)
	fmt.Println(q.Length())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q)
}

func testStack() {
	s := list.NewStack(1, 2)
	fmt.Println(s)
	fmt.Println(s.Length())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	s.Clear()
	fmt.Println(s)
}
