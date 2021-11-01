package file

import (
	"fmt"
	"os"
)

//判断文件路径是否正确，如果正确返回文件信息
func Info(file string) {
	info, _ := os.Stat(file)
	pwd, _ := os.Getwd()
	fmt.Println(pwd, info)
}
