package main

import "fmt"

func main() {
	//1、‘&’符号的用法
	//2、’*’符号的用法

	// n := 10
	// p := &n                  //获取n的地址
	// fmt.Println("n的地址是：", p) //输出n的地址
	// fmt.Println("n的值是：", *p) //输出n的值

	// fmt.Printf("%T\n",p)

	// *p = 20                 //修改n的值
	// fmt.Println("n的值是：", n) //输出n的值

	//3、new()函数的用法
	var a = new(int)
	*a = 10
	fmt.Println(*a)

}
