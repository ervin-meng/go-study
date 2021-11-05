package main

import (
	"fmt"
	"go-study/statement/control/branch"
	"go-study/statement/control/loop"
)

func init() {
	fmt.Println("学习流程控制")
}

func main() {
	branch.StudyIf(1, 2, 3)
	branch.StudySwitch(1)
	loop.StudyFor()
}
