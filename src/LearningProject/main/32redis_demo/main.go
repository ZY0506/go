package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis demo

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// ping
	_, err = redisdb.Ping().Result()
	return
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed,err:%v\n", err)
		return
	}
	fmt.Println("connect redis success")
}

//hello
