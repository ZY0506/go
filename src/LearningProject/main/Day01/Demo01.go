package main
import "fmt"

//全局变量
var num1 = 11
var num2 = 90

var(
	name = "Jack"
	height = 180.9
)

func main(){
	var num int
	num = 18
	fmt.Println(num)

	var age = 18
	fmt.Println(age)

	temp := "abcdefg"
	fmt.Println(temp)

	fmt.Println("-------------------------------------------------")
	
	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Println(name)
	fmt.Println(height)

}