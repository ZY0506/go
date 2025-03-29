package main

import "fmt"
func main() {
	// a1 := []int{1, 2, 3, 4, 5}
	// a2 := a1
	// a3 := make([]int, 5, 5)
	// copy(a3, a1)
	// fmt.Println(a1, a2, a3)

	// a1[0] = 100
	// fmt.Println(a1, a2, a3) //[100 2 3 4 5] [100 2 3 4 5] [1 2 3 4 5]

	// //将索引为 1 的元素删除
	// a1 = append(a1[:1], a1[2:]...)
	// fmt.Println(a1) //[100 3 4 5]
	// fmt.Println("len(a1):", len(a1),"cap(a1):", cap(a1))

	x1 := []int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))	//[1 3 5] 3 3
	//切片不保存具体的值
	//切片对应一个底层数组
	//底层数组是一块连续的内存空间
	s1 = append(s1[:1],s1[2:]...)	
	fmt.Println(s1, len(s1), cap(s1))	//[1 5] 2 2
	fmt.Println(x1, len(x1), cap(x1))	//[1 5 5] 3 3
}
