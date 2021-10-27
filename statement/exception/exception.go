package main

import "fmt"

func main() {
	defer func() {
		info := recover()
		if info != nil {
			fmt.Println(info)
		}
	}()

	panic("exception")
}
