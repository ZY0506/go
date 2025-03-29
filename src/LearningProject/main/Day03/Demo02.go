package main

import "fmt"

func main() {
	//切片(Slice)
	var str []string
	str = []string{"深圳", "揭阳", "东莞"}
	fmt.Println(str)

	var num = []int{1, 2, 4, 6, 7}
	fmt.Println(num)

	fmt.Println("str长度：",len(str),"容量：",cap(str))
	fmt.Println("num长度：",len(num),"容量：",cap(num))

	//利用数组定义切片(下标不能溢出)
	num1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	slice := num1[1:5] 		//左闭右开区间
	fmt.Println(slice)

	s1 :=num1[:4]		//从开始切到第4个
	fmt.Println(s1)
	s2 :=num1[1:]		//从第1个切到最后一个
	fmt.Println(s2)


	//切片的长度为切的底层数组的长度
	//切片的容量为，切的开头到数组末尾
	fmt.Println("slice长度：",len(slice),"容量：",cap(slice))
	fmt.Println("s1长度：",len(s1),"容量：",cap(s1))

	var s []string
	fmt.Println(s==nil)		//只是声明，切片为空
	var b = []bool{}
	fmt.Println(b==nil)		//创建了一个长度为0的切片，但不为空  
	var n = []int{}
	fmt.Println(n==nil)

	num2 := num1[1:5]	
	num3 := num2[2:4]
	num1[3] = 999
	fmt.Println(num3)

}
