package main

import "fmt"

func StudyArray() {
	//数组长度是固定的，数组的长度需要在编译阶段确定，数组的零值为每个元素的对应类型的零值
	//声明一维数组
	var arr1 [2]int
	//声明二维数组
	var arr11 [2][2]int
	//赋值一维数组
	var arr2 = [2]int{1, 2}
	//赋值二维数组
	var arr22 = [2][2]int{{1, 2}, {3, 4}}
	var arr3 = [...]int{1, 2, 3}
	//指定索引赋值
	arr4 := [...]int{1: 1, 0: 0}
	arr5 := [...]string{4: "4444"}

	fmt.Println(arr1, arr11, arr2, arr22, arr3, arr4, arr5)
	//取值
	fmt.Println(arr5[4])
	fmt.Println(arr22[1][0])
	//修改
	arr22[1][0] = 33
	fmt.Println(arr22)
	//比较
	if arr1 == arr2 {

	}
}
