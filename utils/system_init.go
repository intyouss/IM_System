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
	var ctx = context.Background()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password:     viper.GetString("redis.password"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
		DB:           viper.GetInt("redis.dbNumber"),
	})
	pong, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic("Failed the start Redis")
	}
	fmt.Println("Redis inited!!", pong)
}

func Init() {
	initConfig()
	initMysql()
	initRedis()
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) {
	n, err := RedisDB.Publish(ctx, channel, msg).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d clients received the message\n", n)
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	sub := RedisDB.Subscribe(ctx, channel)
	ch := sub.Channel()
	return ch
}
