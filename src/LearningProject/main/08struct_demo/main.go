package main

import (
	"fmt"
)

type person struct{
	name string
	age int
	gender string
}


func main() {
	p1 := person{"张三", 25, "男"}
	fmt.Println(p1)

	pointer := &p1
	fmt.Println(pointer.name, pointer.age, pointer.gender)

	//匿名结构体
	var s struct{
		name string
		age int
	}
	s.age = 25
	s.name = "李四"
	fmt.Printf("type : %T, value : %v\n",s,s)
}
