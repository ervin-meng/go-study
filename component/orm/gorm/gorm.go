package main

import (
	"fmt"
	"go-study/project/ms/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

//https://gorm.io/zh_CN/docs/create.html

type TbUser struct {
	Id         int64     `json:"id" xorm:"'id' not null pk comment('用户ID') INT"`
	Name       string    `json:"name" xorm:"'name' not null default '' comment('用户名') VARCHAR(255)"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' not null default current_timestamp() updated comment('更新时间') DATETIME"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' not null default current_timestamp() created comment('创建时间') DATETIME"`
}

func (TbUser) TableName() string {
	return "tb_user"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3305)/go_project?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\n\r", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	doMigrate(db)
	//doInsert(db)
	//doDelete(db)
	//doUpdate(db)
	//doSelect(db)
}

func doInsert(db *gorm.DB) {
	newUser := TbUser{
		Name: "张三",
	}

	result := db.Create(&newUser)

	//fmt.Println(newUser.Id)
	fmt.Println(result.Error)
	//fmt.Println(result.RowsAffected)

	//批量插入
	users := []TbUser{{Name: "z1"}, {Name: "z2"}, {Name: "z3"}}
	//db.Create(&users)

	db.CreateInBatches(&users, 2) //两条两条的插入
}

func doDelete(db *gorm.DB) {
	//删除主键为1的数据
	db.Delete(&TbUser{}, 1)
	//删除name是张三的数据
	db.Delete(&TbUser{}, "name = ?", "张三")
}

func doUpdate(db *gorm.DB) {
	//只更新非零值
	db.Model(&TbUser{}).Where("name = ?", "张三").Updates(TbUser{Name: "李四"})
	//结构体也要改成NullString类型 解决updates零值问题，也可以吧Name指定称为*string
	//db.Model(&user).Updates(TbUser{Name:sql.NullString{"", true}})

	//update不会有更新零值的问题
	//db.Model(&TbUser{Id:2}).Update("Name","王五")
}

func doSelect(db *gorm.DB) {
	//不指定排序字段
	db.Take(&TbUser{})
	//根据主键升序排序,取第一条
	db.First(&TbUser{})
	db.First(&TbUser{}, 1)
	db.First(&TbUser{}, "name = ?", "张三")
	//根据主键降序，取第一条
	db.Last(&TbUser{})
	//查询所有数据
	db.Find(&[]TbUser{})
	//查询主键在1，2，3，4中的数据
	db.Find(&[]TbUser{}, []int{1, 2, 3, 4})

}

func doMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
