package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func studyNacos() {

	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "fe281935-c397-40f7-8fb2-4de3f88c6357",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "component/distribute/config/tmp/nacos/log",
		CacheDir:            "component/distribute/config/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: "api.yml",
		Group:  "dev",
	})

	fmt.Println(content)

	_ = client.ListenConfig(vo.ConfigParam{
		DataId: "api.ymp",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
}
