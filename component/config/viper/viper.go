package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type StudyServerConfig struct {
	Name  string      `mapstructure:"name"`
	Host  string      `mapstructure:"host"`
	Port  int         `mapstructure:"port"`
	Mysql MysqlConfig `mapstructure:"mysql"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("../config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(v.Get("name"))

	c := StudyServerConfig{}
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}

	fmt.Println(c.Mysql.Port)

	viper.AutomaticEnv()

	fmt.Println(viper.GetString("HOME"))

	//动态监控文件变化
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&c)
	})
}
