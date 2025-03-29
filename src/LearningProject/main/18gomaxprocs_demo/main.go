package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
	wg.Done()
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
	wg.Done()
}

func main() {
	//runtime.GOMAXPROCS(1) //只占一个CPU核心 这样，如果有并发操作，会先把某一个执行完再执行另外一个
	runtime.GOMAXPROCS(2)	//占多个CPU核心
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
