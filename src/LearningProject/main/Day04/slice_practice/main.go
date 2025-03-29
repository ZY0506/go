package main

import (
	"fmt"
	"sort"
)

func main(){
	var a = make([]int,5,10)
	fmt.Println(a)	//[0 0 0 0 0]
	for i:=0;i<10;i++{
		a=append(a,i)
	}
	fmt.Println(a)	//[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println("len(a):",len(a),"cap(a):",cap(a))	//len(a): 15 cap(a): 20


	var a1 = [...]int{3,40,2,23,48,11,9}
	sort.Ints(a1[:])
	fmt.Println(a1)	//[2 3 9 11 23 40 48]
}