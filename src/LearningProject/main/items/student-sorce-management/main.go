package main

import (
	"fmt"
	"os"
)
var(
	allStudents map[int64]*student
)
//定义学生结构体
type student struct {
	id     int64
	name   string
	gender string
}

func TotalMenu(){
	fmt.Println("Welcome to Student Sorce Management System")
	fmt.Println("1. Add Student")
	fmt.Println("2. View Student")
	fmt.Println("3. Delete Student")
	fmt.Println("0. Exiting Program")
}

func AddStudent(){
	//获取学生信息
	var(
		id     int64
		name   string
		gender string
	)
	fmt.Println("Enter Student Information")
	fmt.Print("Enter ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Gender: ")
	fmt.Scanln(&gender)
	//添加学生信息到map
	allStudents[id] = &student{id,name,gender}
}

func ViewStudent(){
	//打印所有学生信息
	for k,v := range allStudents{
		fmt.Printf("ID: %d, Name: %s, Gender: %s\n",k,v.name,v.gender)
	}
}

func DeleteStudent(){
	//获取学生ID
	var(
		deleteID int64
	)
	fmt.Print("Enter ID of Student to Delete: ")
	fmt.Scanln(&deleteID)
	//删除学生信息
	delete(allStudents,deleteID)
}
func main() {
	//初始化学生信息
	allStudents = make(map[int64]*student)
	for true {
		//打印菜单
		TotalMenu()
		//选择功能
		var choice int
		fmt.Scanln(&choice)
		//匹配选择
		switch choice {
		case 1:
			AddStudent()
		case 2:
			ViewStudent()
		case 3:
			DeleteStudent()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
