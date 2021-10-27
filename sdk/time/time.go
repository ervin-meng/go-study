package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("当前时间：", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("当前时间戳：", currentTime.Unix())
	fmt.Println("耗时：", time.Since(currentTime))
}
