package main

import "fmt"

//a-z 26个字母相当于26进制，写一个函数输出字符串，比如0为a 25为z 26为ba 52为ca 676为baa
func main() {

	charMap := map[int]string{
		0:  "a",
		1:  "b",
		2:  "c",
		3:  "d",
		4:  "e",
		5:  "f",
		6:  "g",
		7:  "h",
		8:  "i",
		9:  "j",
		10: "k",
		11: "l",
		12: "m",
		13: "n",
		14: "o",
		15: "p",
		16: "q",
		17: "r",
		18: "s",
		19: "t",
		20: "u",
		21: "v",
		22: "w",
		23: "x",
		24: "y",
		25: "z",
	}

	number := 676
	result := ""

	for {
		round := number / 26

		remind := number % 26

		if round == 0 {
			result = charMap[remind] + result
			break
		}

		if round < 26 {
			result = charMap[round] + charMap[remind] + result
			break
		}

		result = charMap[remind]
		number = round
	}

	fmt.Println(result)
}
