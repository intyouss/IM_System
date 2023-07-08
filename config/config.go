package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
	fmt.Println("System config inited!!")
}

func GetMysql() map[string]string {
	return viper.GetStringMapString("mysql")
}

func GetRedis() map[string]interface{} {
	return viper.GetStringMap("redis")
}

func GetMysqlLog() map[string]interface{} {
	return viper.GetStringMap("mysql.log")
}
