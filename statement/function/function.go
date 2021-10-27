package main

type sumInt struct {
	args []int
}

func (s sumInt) invoke() int {
	result := 0

	for _, num := range s.args {
		result += num
	}

	return result
}

func main() {
	sum := intSum(1, 2)
	println(sum)
	sum1 := intSum1([]int{1, 2, 3, 4, 5}...)
	println(sum1)
	sum2, _ := intSum2([]int{1, 2, 3, 4, 5}...)
	println(sum2)
	//闭包函数
	var sum3 func(int, int) int
	sum3 = intSum3()
	println(sum3(1, 2))
	//函数值
	sum4 := intSum4(sum3)
	println(sum4)
	//匿名函数
	sum5 := func(a int, b int) int {
		return a + b
	}
	println(sum5(1, 2))
	//方法
	ss := sumInt{[]int{1, 2, 3, 4, 5}}
	println(ss.invoke())
	sss := ss.invoke
	println(sss())
}

//返回值列表带名字
func intSum(a int, b int) (result int) {
	result = a + b
	return
}

//可变参数列表
func intSum1(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}

	return result
}

//多个返回值
func intSum2(nums ...int) (int, error) {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result, nil
}

//闭包函数
func intSum3() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}

//使用函数值
func intSum4(fn func(int, int) int) int {
	return fn(1, 2)
}
