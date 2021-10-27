package control

func StudyFor() {
	for i := 1; i <= 5; i += 1 {
		println(i)
	}

	i := 1
	for i <= 5 {
		println(i)
		i += 1
	}

	i = 1
	for {
		if i > 5 {
			break
		}
		println(i)
		i += 1
	}

	//遍历切片
	for index, value := range []int{1, 2, 3, 4, 5} {
		println(index, value)
	}
	//遍历字典
	for key, value := range map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5} {
		println(key, value)
	}
	//TODO 遍历通道
}
