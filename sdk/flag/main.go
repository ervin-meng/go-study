package main

import (
	"flag"
	"fmt"
)

func main() {
	ip := flag.String("ip", "0.0.0.0", "")
	flag.Parse()
	fmt.Printf("%s", *ip)
}
