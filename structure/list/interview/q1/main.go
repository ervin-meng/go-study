package main

import "go-study/structure/list"

func main() {
	//校验输入的字符串中括号是否闭合匹配
	str := "abc{{dddaaddd]cc}"

	stack := list.NewStack()

	symbolMap := map[string]string{"{": "}", "[": "]", "(": ")"}

	for _, char := range str {
		switch string(char) {
		case "{", "[", "(":
			stack.Push(string(char))
		case "}", "]", ")":
			symbol, err := stack.Pop()
			if err != nil {
				panic("缺少" + string(char) + "对应的左括号")
			}
			if string(char) != symbolMap[symbol.(string)] {
				panic("左括号和右括号不匹配")
			}
		}
	}

	if !stack.IsEmpty() {
		panic("括号数量不匹配")
	}
}
