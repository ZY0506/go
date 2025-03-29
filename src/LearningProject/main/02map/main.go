package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 声明map后，必须初始化（分配空间）
	m1 = make(map[string]int, 10) // 声明并初始化map，尽量预估map的容量大小
	m1["娜扎"] = 20
	m1["小美"] = 18
	m1["小刚"] = 19
	fmt.Println(m1)

	value, ok := m1["小丑"] //判断某个key是否存在
	if !ok {
		fmt.Println("找不到对应的key")
	} else {
		fmt.Println(value)
	}

	// 遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range m1{
		fmt.Println(k)
	}
	//只遍历value
	for _ ,v := range m1{
		fmt.Println(v)
	}

	// 删除元素
	delete(m1, "娜扎")
	fmt.Println(m1)
	delete(m1,"小丑")	//If m is nil or there is no such element, delete is a no-op.
	fmt.Println(m1)
}
