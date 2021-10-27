package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

//服务注册

//put http://127.0.0.1:8500/v1/agent/service/register

//header content-type application/json

//body  {
// 	"Name": "api-user",
//	"ID": "api-user",
//  "Tags": ["api", "user"],
// "Address": "127.0.0.1",
// "Port":8501
//}

//服务注销
//put http://127.0.0.1:8500/v1/agent/service/deregister/api-user

func consulRegister(address string, port int, name string, tags []string, id string) error {

	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.0.102:8501/v1/user/list",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	reg := new(api.AgentServiceRegistration)
	reg.Name = name
	reg.ID = id
	reg.Port = port
	reg.Tags = tags
	reg.Address = address
	reg.Check = check

	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, _ := api.NewClient(cfg)

	err := client.Agent().ServiceRegister(reg)

	return err
}

func consulDiscover() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, _ := api.NewClient(cfg)

	services, _ := client.Agent().Services()
	//services,_ := client.Agent().ServicesWithFilter(`Service == "api-user"`)

	for key, value := range services {
		fmt.Println(key, value)
	}
}
