package main

import (
	"encoding/base64"
	"fmt"
)

func StudyBase64() {
	rawStr := "zheshiyigezifuchuan"
	base64Str := base64.StdEncoding.EncodeToString([]byte(rawStr))
	fmt.Println(base64Str)
	raw1Str, _ := base64.StdEncoding.DecodeString(base64Str)
	fmt.Println(string(raw1Str))
	//如果在url中传输要使用UrlEncoding
}
