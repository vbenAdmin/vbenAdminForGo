package Dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "vbenAdmin:WD82ixHxRdwjXBpf@tcp(159.75.188.250:3306)/vbenadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库链接失败：", err)
	}
	log.Println("数据库已连接！")
	return db
}
