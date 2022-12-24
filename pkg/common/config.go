package common

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type ServerConfig struct {
	Port        int  `yaml:"port"`
	DebugMode   bool `yaml:"debugMode"`
	CallBackUrl string
}

type DBConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxConnLifeTime int    `yaml:"max_conn_life_time"`
}

type LogConfig struct {
	LogDir     string `yaml:"logDir"`
	LogName    string `yaml:"logName"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"maxsize"`
	MaxBackups int    `yaml:"maxbackups"`
	MaxAge     int    `yaml:"maxage"`
	Compress   bool   `yaml:"compress"`
	Encoder    string `yaml:"encoder"`
}

type YamlConfig struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Log    LogConfig    `yaml:"log"`
}

var AppConfig YamlConfig

func init() {
	// env > default value
	var config string
	if configEnv := os.Getenv("CONFIG_ENV"); configEnv == "" {
		config = "default.yaml"
	} else {
		config = fmt.Sprintf("%s.yaml", configEnv)
	}
	wd, _ := os.Getwd()
	config = fmt.Sprintf("%s/conf/%s", wd, config)

	fmt.Printf("use config file: %s\n", config)

	_, err := os.Stat(config)
	if err != nil {
		panic(fmt.Errorf("stat config fiel error: %v\n", err))
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err = v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file %v \n", err))
	}
	if err = v.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("Unmarshal yaml config file error %v \n", err))
	}
}
