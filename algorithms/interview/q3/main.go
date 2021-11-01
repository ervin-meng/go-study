package main

import "fmt"

//反转字符串，连续的字符或‘.’当成一个字符，如$a = ‘ab...asd.asg..dsf’;输出为‘dsf..asg.asd...ab’
func main() {
	str := "ab...asd.asg..dsf"
	strLength := len(str)
	leftWord := ""
	leftStr := ""
	rightWord := ""
	rightStr := ""

	for i, j := 0, strLength-1; i <= j; i, j = i+1, j-1 {

		leftChar := string(str[i])

		if leftChar == "." {
			if leftWord != "" {
				leftChar = leftChar + leftWord
				leftWord = ""
			}
			leftStr = leftChar + leftStr
		} else {
			leftWord = leftWord + leftChar
		}

		if i == j {
			break
		}

		rightChar := string(str[j])

		if rightChar == "." {
			if rightWord != "" {
				rightChar = rightWord + rightChar
				rightWord = ""
			}
			rightStr = rightStr + rightChar
		} else {
			rightWord = rightChar + rightWord
		}
	}

	fmt.Println(rightStr + rightWord + leftWord + leftStr)
}
