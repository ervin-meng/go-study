package main

import (
	"fmt"
	"go-study/structure/list"
)

func main() {
	//l := list.NewSeqList(1,2,3)
	//fmt.Println(l)
	//fmt.Println(l.IsEmpty())
	//l.Append("four")
	//fmt.Println(l)
	//err := l.Insert(3, "five")
	//fmt.Println(err, l)
	//err = l.Delete(1)
	//fmt.Println(err, l)
	//fmt.Println(l.Get(0))

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
