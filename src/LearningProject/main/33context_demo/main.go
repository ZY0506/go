package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// "fmt"
// "sync"
// "time"

//为什么需要context

var (
	wg sync.WaitGroup
	// exitchan = make(chan struct{},1)
)

// func f() {
// 	defer wg.Done()
// 	for {
// 		fmt.Println("周琳")
// 		time.Sleep(500 * time.Millisecond)
// 		select {
// 		case <-exitchan:
// 			fmt.Println("exit")
// 			return
// 		default:
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go f()
// 	time.Sleep(5 * time.Second)
// 	exitchan<-struct{}{}
// 	wg.Wait()
// }

//测试函数

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit2")
			break LOOP
		default:
			fmt.Println("北京")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit1")
			break LOOP
		default:
			fmt.Println("周琳")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(5 * time.Second)
	//通知子协程退出
	cancel()
	wg.Wait()
}