package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func StudyChannel() {
	fmt.Println("--------------学习通道-------------")
	var chan1 chan int
	fmt.Println(chan1)
	//无缓冲通道
	chan2 := make(chan int)
	chan2 = make(chan int, 0)
	fmt.Println(chan2)
	//有缓冲通道
	chan3 := make(chan int, 3)
	fmt.Println(chan3)
	//单方向通道，发送
	chan4 := make(chan<- int, 3)
	fmt.Println(chan4)
	//单方向通道，接收
	chan5 := make(<-chan int, 3)
	fmt.Println(chan5)
	//fmt.Println(recv())
	recv()
	println(runtime.NumGoroutine())
}

func recv() string {
	recv := make(chan string, 5)

	go func() {
		recv <- request("http://www.baidu.com")
	}()

	//会造成goroutine泄露
	go func() {
		recv <- request("http://www.baidu.com")
	}()

	return <-recv
}

func request(hostname string) string {
	res, err := http.Get(hostname)
	if err != nil {
		return fmt.Sprintln(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintln(err)
	}
	return string(content)
}
