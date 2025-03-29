package main

import "fmt"

func main01() {
	// for i:=1;i<10;i++{
	// 	for j:=1;j<=i;j++{
	// 		fmt.Printf("%d*%d=%d\t",j,i,i*j)
	// 	}
	// 	fmt.Println()
	// }

	// var a [3]bool
	// fmt.Println(a)
	// a = [3]bool{true, true, true}
	// fmt.Println(a)

	// a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} //根据初始化数量算出长度
	// fmt.Println(len(a1))

	// a2 := [10]int{0: 1, 9: 999}
	// fmt.Println(a2)

	// citys:=[...]string{"北京","上海","天津","广州","深圳"}
	// //i:下标，v:值
	// for i,v:=range citys{
	// 	fmt.Println(i,v)
	// }

	// //多维数组
	// var a11 [3][2]int
	// a11 = [3][2]int{
	// 	[2]int{1,2},
	// 	[2]int{3,4},
	// 	[2]int{5,6},
	// }
	// fmt.Println(a11)

	// //遍历
	// for _,v1:=range a11{
	// 	//fmt.Println(v1)
	// 	for _,v2:=range v1{
	// 		fmt.Println(v2)
	// 	}
	// }

	array :=[...]int{1,3,5,67,7}
	var sum int
	for _,v :=range array{
		sum+=v
	}
	fmt.Println(sum)
}
