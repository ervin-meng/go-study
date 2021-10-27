package main

import "fmt"

func StudyMap() {
	fmt.Println("-----------学习字典----------------")
	//使用零值初始化
	var map1 map[int]int
	fmt.Println(map1)
	//使用make函数初始化
	map2 := make(map[int]int)
	fmt.Println(map2)
	//使用字面量初始化
	map3 := map[int]int{1: 0, 2: 1}
	fmt.Println(map3)
	//取值,查找失败返回valueType对应的零值
	v, ok := map3[1]
	fmt.Println(v, ok)
	v, ok = map3[0]
	fmt.Println(v, ok)
	//添加，对一个nil值进行添加，报panic异常
	map3[3] = 3
	fmt.Println(map3)
	//删除
	delete(map3, 3)
	fmt.Println(map3)
	//修改
	map3[1] = 1
	map3[2] = 2
	fmt.Println(map3)
	//遍历是随机的，字典的Key值可以先放到切片实现Key排序，再遍历
	for key, value := range map3 {
		fmt.Println(key, value)
	}
}
