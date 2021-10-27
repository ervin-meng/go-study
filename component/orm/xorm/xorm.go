package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

type TbUser struct {
	Id         int64     `json:"id" xorm:"'id' not null pk comment('用户ID') INT"`
	Name       string    `json:"name" xorm:"'name' not null default '' comment('用户名') VARCHAR(255)"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' not null default current_timestamp() updated comment('更新时间') DATETIME"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' not null default current_timestamp() created comment('创建时间') DATETIME"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3305)/go_study")

	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(TbUser))

	if err != nil {
		log.Fatal(err)
	}
}
