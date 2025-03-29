package main

import (
	"fmt"
)
type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func printValue (val interface{}){
	fmt.Printf("Value: %v,Type: %T\n",val,val)
}

type Reader interface{
	Read() string
}
type Writer interface{
	Write(string) 
}

type ReadWriter interface{
	Reader
	Writer
}
type File struct{}

func (f File) Read() string {
        return "Reading data"
}

func (f File) Write(data string) {
        fmt.Println("Writing data:", data)
}

func main() {
	c := Circle{5}
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())

	printValue(42)
	printValue("Hello, World!")

	//类型断言
	var iface interface{} = "hello"
	str := iface.(string)
	fmt.Println(str)

	//类型判断
	if val, ok := iface.(string); ok {
		fmt.Printf("Value: %s,Type: %T\n",val,val)
	} else {
		fmt.Println("not string")
	}

	//接口组合
	rw := File{}
	var r ReadWriter = rw
	fmt.Println(r.Read())
	r.Write("Hello, World!")


}
