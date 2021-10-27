package main

import "fmt"

type study struct {
	typeName string
}

func (s study) printTypeName() {
	fmt.Println(s.typeName)
}

func StudyStruct() {
	fmt.Println("-----------学习结构体--------------")

	//匿名成员
	type studyWrapper struct {
		study
		name string
	}

	study1 := studyWrapper{study{typeName: "Wrapper"}, "ha"}

	fmt.Println(study1.name)

	study1.printTypeName()

	//非匿名成员
	type studyCombination struct {
		s    study
		name string
	}

	study2 := studyCombination{s: study{typeName: "Combination"}, name: "你猜"}

	study2.s.printTypeName()

	//结构体是值赋值，如果这样赋值下面的语句时不能改变study3的typeName属性的
	study3 := study{"t3"}
	//如果赋值的时候使用指针，则下面的语句会修改study3的typename
	//study3 := &study{"t3"}
	study4 := study3

	study4.typeName = "t4"

	fmt.Println(study3.typeName, study4.typeName)
}
