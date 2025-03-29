package main

import "fmt"

func main() {
	//元素为map的切片
	var s1 = make([]map[string]int, 10, 10)
	s1[0] = make(map[string]int, 1)
	s1[0]["北京"] = 10
	//只有第一个是有内容的，其他都是nil（没有初始化）
	fmt.Println(s1) //[map[北京:10] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	//元素为切片的map
	var m1 = make(map[int][]int, 10)
	m1[0] = []int{1, 2, 10}
	fmt.Println(m1)	//map[0:[1 2 10]]
}
