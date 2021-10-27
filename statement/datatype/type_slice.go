package main

import (
	"fmt"
	"reflect"
)

func StudySlice() {
	fmt.Println("-------------学习切片-----------------")
	//切片不需要指定长度
	var slice1 []int
	//使用make创建slice
	slice2 := make([]int, 10, 100)
	fmt.Println(slice1, slice2)
	//切片是对数组的抽象，数组的长度不可改变，在特定场景中就不太适用，因此提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	var arr = []interface{}{5, 6, 7, 8, 9}
	var slice3 = arr[1:4]
	var slice4 = arr[0:4]
	var slice5 = arr[1:]

	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))
	fmt.Println(slice5, len(slice5), cap(slice5))

	//slice6 := make([]int, 0 , 4)

	//slice6 = append(slice6, arr)

	slice7 := sliceType{1, 2, 3, 4}
	slice8 := slice7.Insert(1, "a", "b")

	var slice9 [][]int

	slice9 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(slice8, slice9)
}

type sliceType []interface{}

func (s sliceType) Insert(index int, elements ...interface{}) sliceType {
	result := make([]interface{}, len(s)+len(elements))
	at := copy(result, s[:index])
	at += copy(result[at:], elements)
	copy(result[at:], s[index:])
	return result
}

func (s sliceType) Del() {

}

func (s sliceType) Replace() {

}

func (s sliceType) Get() {

}

func (s sliceType) valueExists() {

}

func (s sliceType) indexExists() {

}

func (s sliceType) Add(start int, elements ...interface{}) sliceType {
	ss := s[:start]
	last := s[start:]

	elementsValue := reflect.ValueOf(elements)

	switch elementsValue.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		sss := reflect.ValueOf(&ss).Elem()
		for i := 0; i < elementsValue.Len(); i += 1 {
			sss = reflect.Append(sss, elementsValue.Index(i))
		}
		ss = sss.Interface().(sliceType)
	default:
	}

	ss = append(ss, last...)

	return ss
}
