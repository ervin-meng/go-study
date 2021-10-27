package main

import "fmt"

func main() {
	fmt.Println("var a int = 1")

	var a int = 1

	fmt.Println("省略变量类型")
	fmt.Println("var b = 2")
	var b = 2

	fmt.Println("省略变量值")
	fmt.Println("var c int \n c = 3")

	var c int

	c = 3

	fmt.Println("使用短变量声明")
	fmt.Println("d := 4")
	d := 4

	var (
		e = 6
		f = 7
	)

	fmt.Println("使用new函数")
	fmt.Println("var g = new(int) \n *g = 8")
	var g = new(int)

	*g = 8

	fmt.Println("使用下划线丢弃一个变量")
	fmt.Println("var _ = 9")

	var _ = 9

	fmt.Println("fmt.Println(a,b,c,d,e,f,g)")

	fmt.Println(a, b, c, d, e, f, *g)
}
