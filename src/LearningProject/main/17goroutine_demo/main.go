package main

import (
	"fmt"
	"sync"
)

//goroutine demo

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello 娜扎",i)
	wg.Done()	//执行完 通知wg把计数器-1
}

// func main() { //开启一个主goroutine去执行main函数

// 	wg.Add(1000)
// 	for i:=0;i<1000;i++{
// 		//wg.Add(1)
// 		go hello(i)
// 	}

// 	// wg.Add(1)  //计数牌-1
// 	// go hello() //开启了一个goroutine去执行hello函数	有可能打印不出来（main函数执行的比较快，main结束了hello函数还没执行完）
// 	fmt.Println("hello main")
// 	//解决主goroutine执行完，其他goroutine没执行完的问题
// 	//time.Sleep(time.Second)	//不建议
// 	wg.Wait() //等所有小弟干完活才结束
// }

func main(){
	wg.Add(10000)
	for i:=1;i<=10000;i++{
		go func(i int){
			fmt.Println("hello",i)	//可能会出现 同一个i打印多次(并发的问题)
			wg.Done()
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait()
}

