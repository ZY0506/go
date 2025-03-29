package main

import (
	"fmt"
	"strconv"
)

func main() {
	//将基本数据类型转化为字符串
	// var num1 int = 78
	// var str1 string = fmt.Sprintf("%d", num1)
	// fmt.Printf("str1 的值 %q\n", str1)

	// var flag bool = false
	// var str2 string = fmt.Sprintf("%t", flag)
	// fmt.Printf("str2 的值 %q\n", str2)

	// var num2 float32 = 3.1415
	// var str3 string = fmt.Sprintf("%f", num2)
	// fmt.Printf("str3 的值 %q\n", str3)

	// fmt.Printf("str1 的数据类型 ： %T\n", str1)

	//将字符串转化为其他基本数据类型

	//func ParseBool(str string) (value bool, err error)
	//两个返回值：value指返回的值，err指可能出现的错误
	var str1 string = "false"
	var flag bool
	flag, _ = strconv.ParseBool(str1)
	fmt.Printf("flag 的数据类型 %T , flag = %v\n", flag, flag)

	//func ParseFloat(s string, bitSize int) (f float64, err error)
	var str2 string = "3.1415926"
	var num1 float64
	num1, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("num1 的数据类型 %T , num1 = %v\n", num1, num1)

	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//注意：转换成对应的不同大小的int类型，可能会溢出
	var str3 string = "12345678"
	var num2 int64
	num2, _ = strconv.ParseInt(str3, 10, 8)
	fmt.Printf("num2 的数据类型 %T , num2 = %v \n", num2, num2)

}
