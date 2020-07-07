package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// redis

// 声明一个全局redis变量
var redisdb *redis.Client

// 初始化链接
func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
	}
	//ret := redisdb.Get("haojie").Val()
	ret := redisdb.Get("haojie").Val()
	fmt.Println(ret)
}
