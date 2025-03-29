package main

import (
	"fmt"
)
type address struct{
	province string
	city string
}
type person struct{
	name string
	age int
	address	//匿名嵌套，但是只能有一个字段
	company
}
type company struct{
	address
	name string
}
func main() {
	p1 := person{
		name : "Tom",
		age : 25,
		address : address{
			province : "Beijing",
			city : "Beijing",
		},
		company : company{
			address : address{
				province : "Shanghai",
				city : "Shanghai",
			},
			name : "Alibaba",
		},
	}
	fmt.Println(p1)
	// fmt.Println(p1.address.province)
	// fmt.Println(p1.city) //error: 如果有两个匿名嵌套，就只能使用以下方式
	fmt.Println(p1.address.city)
	fmt.Println(p1.company.name)
	fmt.Println(p1.company.address.province)
}
