package main

import (
	"fmt"
	"os"
)

var (
	stumgr stuMgr
)

// 总菜单
func totalMenu() {
	fmt.Println("Welcome to my Student Management System")
	fmt.Println(`
	1. Add Student
	2. Show All Students
	3. Modify Student
	4. Delete Student
	5. Exit
	`)
}

func main() {
	stumgr = stuMgr{
		allstudents: make(map[int64]student, 100),
	}
	for {
		totalMenu()
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			stumgr.addStudent()
		case 2:
			stumgr.showAllStudents()
		case 3:
			stumgr.modifyStudent()
		case 4:
			stumgr.deleteStudent()
		case 5:
			fmt.Println("Thank you for using our system")
			os.Exit(1)
		default:
			fmt.Println("滚蛋~")
		}
	}

}
