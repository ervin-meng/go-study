package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func StudyString() {

	//字符串底层是字节数组，是固定不允许修改
	fmt.Println("------------学习字符串--------------")
	var _ string
	var str = "abc\n"  //会有转义操作
	var str1 = `abc\n` //不带有转义，是字面量
	fmt.Println(str, str1)

	//连接字符串
	fmt.Println(str + str1) //每次调用加号底层都会分配新内存

	buffer := bytes.Buffer{}
	buffer.WriteString("a")
	buffer.WriteString("b")

	fmt.Println(buffer.String())

	//统计字符数
	fmt.Println(utf8.RuneCountInString(str1))

	str3 := "中文"

	//统计字节数
	fmt.Println(len(str3))

	//取值
	fmt.Println(string(str1[1]))

	//修改 strVal1[0] = "a" //这是不允许的,只可以给变量赋值新的字符串值

	str2 := &str1

	fmt.Println(*str2)
}
