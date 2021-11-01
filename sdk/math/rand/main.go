package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 10; i = i + 1 {
		//返回随机数
		fmt.Println(time.Duration(rand.Uint64()%20) * time.Millisecond)
	}
}
