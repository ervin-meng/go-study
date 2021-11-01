package main

import (
	"go-study/sdk/os/file"
	"os"
)

func main() {
	file.Info("sdk/os/file/file.go")
}

//获取当前工作目录
func getPwd() string {
	pwd, _ := os.Getwd()

	return pwd
}
