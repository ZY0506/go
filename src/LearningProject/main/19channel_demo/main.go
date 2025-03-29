package main

import "fmt"

func main() {
	var ch1 chan int        //引用类型，必须初始化才能使用
	ch1 = make(chan int, 1) //有缓冲区的通道

	//ch1 = make(chan int)	//无缓冲区的通道，又称为同步通道 会发生死锁
	ch1 <- 10 //发送值
	x := <-ch1
	fmt.Println(x)
	fmt.Printf("len : %v;cap : %v", len(ch1), cap(ch1))
	close(ch1) //最好手动关掉
}
