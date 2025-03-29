package main

import "fmt"

//make函数构造切片

func main() {
	//make([]T, size, cap)
	s1 := make([]int, 5, 10)
	fmt.Printf("s1: %v 长度：%d,容量%d\n", s1, len(s1), cap(s1))

	//这两个切片都指向同一个数组
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := s2
	s3[0] = 1000
	fmt.Println(s2, s3)

	//切片支持
	for _, v := range s3 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//append 必须用原来的切片接收返回值
	s4 := []string{"广州", "深圳", "东莞"}
	fmt.Printf("s4:%v,长度：%d,容量：%d\n", s4, len(s4), cap(s4))
	s4 = append(s4, "揭阳")
	fmt.Println(s4)
	fmt.Printf("s4:%v,长度：%d,容量：%d\n", s4, len(s4), cap(s4))
	s4 = append(s4, "汕头")
	fmt.Printf("s4:%v,长度：%d,容量：%d\n", s4, len(s4), cap(s4))

	var slice []string
	slice = append(slice, "北京")
	slice = append(slice, "上海")
	slice = append(slice, "广州")

	ss := []string{"西安","揭阳"}
	//ss...表示将ss切片中的元素作为参数传递给append函数，...表示把切片拆开，变成一个个的元素
	slice = append(slice, ss...)
	
	
}
