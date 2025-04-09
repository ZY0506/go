package main

import (
	"LearningProject/main/items/Logger/mylogger"
	"fmt"
	"time"
)

func main() {
	strat := time.Now()
	log, err := mylogger.NewLog("debug", "zhangyu")
	if err != nil {
		panic(err)
	}
	defer log.Close()
	for i := 0; i < 100; i++ {
		log.Debug("这是一条debug日志")
		log.Trace("这是一条trace日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "小王子"
		log.Error("这是一条error日志,id : %d,name : %s", id, name)
		log.Fatal("这是一条fatal日志")
		time.Sleep(time.Millisecond * 200)
	}
	end := time.Now()
	fmt.Println(end.Sub(strat))
}
