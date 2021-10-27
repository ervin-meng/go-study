package main

import (
	"fmt"
	"reflect"
)

type author struct {
	name string
	age  int
	sex  int
}

func (a author) GetName() string {
	return a.name
}

func (a author) GetAge() int {
	return a.age
}

func (a author) GetSex() int {
	return a.sex
}

func main() {
	author := author{"ervin", 32, 1}
	authorType := reflect.TypeOf(author)
	authorValue := reflect.ValueOf(author)
	fmt.Println(authorType)
	fmt.Println(authorValue)
	authorType1 := reflect.New(authorType)
	fmt.Println(authorType1)
	fmt.Println(authorType.NumField())
	fmt.Println(authorType.NumMethod()) //定义导出的方法数量

	for i := 0; i < authorValue.NumMethod(); i += 1 {
		fmt.Println(authorValue.Method(i).Call([]reflect.Value{})[0].Interface())
	}

	authorAgeValue := reflect.ValueOf(&author.age)
	authorAgeValue.Elem().SetInt(28)

	fmt.Println(author)
}
