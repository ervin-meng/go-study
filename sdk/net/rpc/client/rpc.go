package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	var reply string
	client, _ := rpc.Dial("tcp", ":9001")
	client.Call("HelloService.Hello", "rpc", &reply)
	fmt.Println(reply)
}
