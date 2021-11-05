package main

import "fmt"

func init() {
	fmt.Println("学习常量")
}

func main() {
	const A = 1

	const (
		B = 2
		B1
		C = 3
		C1
	)

	const (
		D = iota + C
		E
		F
	)

	fmt.Println(A, B, B1, C, C1, D, E, F)
}
