package main

import (
	"fmt"
	"time"
)

//work_pool

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d,start job:%d\n", id, job)
		result<- job*2
		time.Sleep(time.Millisecond*500)
		fmt.Printf("worker:%d,end job:%d\n", id, job)

	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//发送五个任务
	for i:=0;i<5;i++{
		jobs<-i
	}
	
	//开启三个gotoutine
	for j:=0;j<3;j++{
		go worker(j,jobs,results)
	}

	close(jobs)	//关闭

	//输出结果
	// for ret:=range results{	//使用这种方式输出结果的话，会出现死锁-->当channel没有元素时，result没有关闭，range会一直读取channel
	// 	fmt.Println(ret)
	// }
	for i:=0;i<5;i++{
		ret := <-results
		fmt.Println(ret)
	}
}
