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

func init() {
	dbConfig := config.AppConfig.DB
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

	if err = db.AutoMigrate(&model.Person{}); err != nil {
		panic(fmt.Errorf("gorm auto migrate error: %v", err))
	}
}
