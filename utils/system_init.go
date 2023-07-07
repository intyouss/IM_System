package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var RedisDB *redis.Client
var Ctx = context.Background()

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
	newLogger := logger.New(
		log.New(os.Stdout, "[Mysql]", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	mysqlConfig := viper.GetStringMapString("mysql")
	dns := mysqlConfig["username"] + ":" + mysqlConfig["password"] + "@tcp(" + mysqlConfig["host"] + ":" +
		mysqlConfig["port"] + ")/" + mysqlConfig["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("Failed the start Mysql")
	}
	fmt.Println("Mysql inited!!")
}

func initRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password:     viper.GetString("redis.password"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
		DB:           viper.GetInt("redis.dbNumber"),
	})
	_, err := RedisDB.Ping(Ctx).Result()
	if err != nil {
		panic("Failed the start Redis")
	}
	fmt.Println("Redis inited!!")
}

func Init() {
	initConfig()
	initMysql()
	initRedis()
}
