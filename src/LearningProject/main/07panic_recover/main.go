package main

import (
	"fmt"
)

func funcA(){
	fmt.Println("A")
}
func funcB(){
	//刚刚打开的数据库连接
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
		fmt.Println("释放数据库连接。。。")
	}()
	panic("出现了严重的错误！！！！")//程序崩溃退出
	fmt.Println("B")
}
func funcC(){
	fmt.Println("C")
}


func main() {
	funcA()
	funcB()
	funcC()
}
