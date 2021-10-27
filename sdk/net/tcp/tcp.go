package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")
	lis, _ := net.ListenTCP("tcp", addr)
	defer lis.Close()
	port := lis.Addr().(*net.TCPAddr).Port
	fmt.Println(addr, port)
}
