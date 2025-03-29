package main

import (
	"fmt"
	"encoding/json"
)
type person struct{
	//用小写的字段名，打印时候内容不可见，用驼峰写法就可见
	Name string	`json:"name db:"name" ini:"name"` //这些表示在json、ini、yaml等格式中，字段的别名
	Age int
}
func main() {
	p1 := person{
		Name: "张三",
		Age: 25,
	}
	b,err := json.Marshal(p1)
	if err != nil {
		fmt.Println("marshal failed , err:",err)
	}
	fmt.Printf("%#v\n",string(b))

	//反序列化
	str:=`{"name":"李四","age":30}`
	var p2 person
	err = json.Unmarshal([]byte(str), &p2)	//注意,这里需要传字符切片以及指针
	if err!= nil {
		fmt.Println("unmarshal failed , err:",err)
	}else{
		fmt.Printf("%#v\n",p2)
	}

}