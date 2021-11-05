package branch

func StudySwitch(a int) {
	normalSwitch(a)
	tagLessSwitch(a)
	typeSwitch()
}

func normalSwitch(a int) {
	//case 默认带break语句
	three := 3
	switch a {
	case 1:
		println(1)
		fallthrough //fallthrough 会强制执行后面的case语句，不会判断下一条case表达式的结果
	case 2:
		println(2)
	case three:
		println(3)
	default:
		println(0)
	}
}

func tagLessSwitch(a int) {
	switch {
	case a == 1:
		println(1)
	case a == 2:
		println(2)
	case a == 3:
		println(3)
	default:
		println(0)
	}
}

func typeSwitch() {
	var b interface{}
	b = 2
	//只能针对接口类型
	switch b.(type) {
	case int:
		println("int")
	}
}
