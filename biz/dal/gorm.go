package dal

import (
	"fmt"
	"github.com/LynnWonder/gin_prac/biz/config"
	"github.com/LynnWonder/gin_prac/biz/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var DBConnection *gorm.DB

// init 函数在程序初始化阶段执行，获得了全局变量 DBConnection
func init() {
	dbConfig := config.AppConfig.DB
	// 返回格式化后的数据库地址包含主机+用户名+密码
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 zapgorm2.New(zap.L()),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(fmt.Errorf("connection to database error: %v", err))
	}
	DBConnection = db

	// 只去更改发生改动的字段
	if err = db.AutoMigrate(&model.Person{}); err != nil {
		panic(fmt.Errorf("gorm auto migrate error: %v", err))
	}
}
