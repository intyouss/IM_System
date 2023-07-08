package utils

import (
	"IM_System/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var RedisDB *redis.Client

func initMysql() {
	mysqlConfig := config.GetMysql()
	dns := mysqlConfig["username"] + ":" + mysqlConfig["password"] + "@tcp(" + mysqlConfig["host"] + ":" +
		mysqlConfig["port"] + ")/" + mysqlConfig["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: MysqlLogger,
	})
	if err != nil {
		panic("Failed the start Mysql")
	}
	fmt.Println("Mysql inited!!")
}

func initRedis() {
	var ctx = context.Background()
	redisConfig := config.GetRedis()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:         redisConfig["host"].(string) + ":" + redisConfig["port"].(string),
		Password:     redisConfig["password"].(string),
		PoolSize:     redisConfig["poolsize"].(int),
		MinIdleConns: redisConfig["minidleconn"].(int),
		DB:           redisConfig["dbnumber"].(int),
	})
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic("Failed the start Redis")
	}
	fmt.Println("Redis inited!!")
}

func InitSystem() {
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
		log.Println("Publish 发布消息失败:", err)
	}
	fmt.Printf("%d clients received the message\n", n)
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	sub := RedisDB.Subscribe(ctx, channel)
	ch := sub.Channel()
	return ch
}
