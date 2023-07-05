package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
	fmt.Println("System config inited!!")
}

func initMysql() {
	mysqlConfig := viper.GetStringMapString("mysql")
	dns := mysqlConfig["username"] + ":" + mysqlConfig["password"] + "@tcp(" + mysqlConfig["host"] + ":" +
		mysqlConfig["port"] + ")/" + mysqlConfig["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed the start database")
	}
	fmt.Println("Database inited!!")
}

func Init() {
	initConfig()
	initMysql()
}
