package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitGorm(mysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlDataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		},
	})
	if err != nil {
		panic("链接数据库失败 error: " + err.Error())
	}

	fmt.Println("mysql链接成功")
	return db

}
