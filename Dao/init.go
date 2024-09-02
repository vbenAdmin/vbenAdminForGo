package Dao

import (
	"goVben/Utils/Config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := Config.Conf.Mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库链接失败：", err)
	}
	log.Println("数据库已连接！")
	return db
}
