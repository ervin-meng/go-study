package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	md5Obj := md5.New()

	_, _ = io.WriteString(md5Obj, "这一个字符串")

	hexStr := hex.EncodeToString(md5Obj.Sum(nil))

	fmt.Println(hexStr)
}
