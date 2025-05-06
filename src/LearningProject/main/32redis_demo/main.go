package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//redis demo

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,   //默认数据库
		PoolSize: 100, //连接池大小
	})
	// ping
	_, err = rdb.Ping().Result()
	return
}

// doCommand go-redis基本使用示例
func doCommand() {

	// 执行命令获取结果
	val, err := rdb.Get("key").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := rdb.Get("key")
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = rdb.Set("key", 10, time.Hour).Err()
	fmt.Println(err)

	// 直接执行命令获取值
	value := rdb.Get("key").Val()
	fmt.Println(value)
}

// zset
// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
	// 注意：v8版本使用[]*redis.Z；此处为v9版本使用[]redis.Z
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// ZADD
	err := rdb.ZAdd(zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed,err:%v\n", err)
		return
	}
	fmt.Println("connect redis success")
	defer rdb.Close()
	// doCommand()
	zsetDemo()
}

//hello
