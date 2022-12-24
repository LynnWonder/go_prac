package db

import (
	"fmt"
	"github.com/LynnWonder/gin_prac/pkg/common"
	"github.com/LynnWonder/gin_prac/pkg/logger"
	"github.com/spf13/cobra"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Models = []interface{}{}
var db *gorm.DB

type Result struct {
	ID           int
	Error        error
	RowsAffected int64
}

func Init(config *common.DBConfig, cmd *cobra.Command) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         255,
	}
	gormConfig := &gorm.Config{
		Logger:                                   logger.GetGormLogger(),
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var err error
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(config.MaxConnLifeTime) * time.Minute)

	// register tables
	autoDB := db.Set("gorm:table_options", "ENGINE = InnoDB DEFAULT CHARSET = utf8 ROW_FORMAT = Dynamic")
	// autoMigrate base Models
	if err = autoDB.AutoMigrate(Models...); err != nil {
		panic(err)
	}
}
