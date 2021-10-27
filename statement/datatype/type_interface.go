package main

import (
	"fmt"
	"time"
)

type IConnector interface {
	connect() bool
}

type Connector struct {
	ip      string
	port    int
	timeout time.Duration
}

func (c Connector) connect() bool {
	fmt.Println(c.port)
	return false
}

func StudyInterface() {
	fmt.Println("--------------学习接口-------------")

	var connector = new(Connector)

	if _, ok := interface{}(connector).(IConnector); ok {
		fmt.Println("实现了IConnector接口")
	}

	connector.connect()
}
