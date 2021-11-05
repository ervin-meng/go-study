package branch

import "fmt"

func StudyIf(a int, b int, c int) {
	if a == b {
		fmt.Println("a == b")
	} else if a == c {
		fmt.Println("a == c")
	} else {
		fmt.Println("a")
	}
}
