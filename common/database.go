package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin-gorm/model"
)

// DB /* common/database.go */
var DB *gorm.DB

// InitDB 数据库初始化
func InitDB() *gorm.DB {
	driverName := "mysql"
	user := "root"
	password := "kljym"
	host := "localhost"
	port := "3306"
	database := "blog"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset)
	// 连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	// 迁移数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

// GetDB 数据库信息获取
func GetDB() *gorm.DB {
	return DB
}
