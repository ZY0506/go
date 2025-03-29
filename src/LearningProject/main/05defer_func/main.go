package main

import (
	"fmt"
)

//defer多用于函数结束之前释放资源，文件句柄、socket连接、数据库连接
func deferDemo(){
	fmt.Println("star")
	defer fmt.Println("哈哈哈")
	defer fmt.Println("嘿嘿嘿")
	fmt.Println("end")
}
func main() {
	deferDemo()
}
