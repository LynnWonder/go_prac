module github.com/LynnWonder/gin_prac

go 1.16

require (
	github.com/gin-contrib/zap v0.0.2
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/universal-translator v0.18.0 // indirect i18n
	github.com/go-playground/validator/v10 v10.11.0 // gin validator
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/spf13/viper v1.12.0 // indirect 解析配置文件 写滚动日志
	github.com/ugorji/go v1.2.6 // indirect
	go.uber.org/zap v1.21.0 // indirect 提供结构化日志
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.6
	moul.io/zapgorm2 v1.1.3 // indirect  😄 zapgorm2 is a zap logging driver for gorm v2
)
