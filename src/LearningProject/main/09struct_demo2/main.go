package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}
func newPerson(name string, age int) person {
	return person{name: name, age: age}
}

type Circle struct{
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	p1 := newPerson("Alice", 25)
	fmt.Println(p1)
	c1 := Circle{radius: 5.0}
	fmt.Println(c1.area())
}
