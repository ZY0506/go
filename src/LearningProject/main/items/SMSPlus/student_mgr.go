package main

import "fmt"

type student struct {
	id   int64
	name string
}
type stuMgr struct {
	allstudents map[int64]student
}

func newStu(id int64, name string) student {
	return student{id, name}
}

// 添加一个学生
func (s stuMgr) addStudent() {
	var (
		stuID   int64
		stuName string
	)
	//获取用户输入
	fmt.Print("Enter New StudentID:")
	fmt.Scanln(&stuID)
	fmt.Print("Enter New StudentName:")
	fmt.Scanln(&stuName)
	//判断学生是否存在
	if _, ok := s.allstudents[stuID]; ok {
		fmt.Println("This student already exists")
		return
	}
	//将学生添加到数组里面
	newstu := newStu(stuID, stuName)
	s.allstudents[stuID] = newstu
	//打印结果
	fmt.Println("Successfully Added")
}

// 打印所有学生
func (s stuMgr) showAllStudents() {
	fmt.Println("Student List:")
	for _, stu := range s.allstudents {
		fmt.Printf("stuID : %v,stuName : %v\n", stu.id, stu.name)
	}
}

// 修改一个学生
func (s stuMgr) modifyStudent() {
	//获取用户输入
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("Enter StudentID To Modify:")
	fmt.Scanln(&stuID)
	//判断学生是否存在
	stuObj, ok := s.allstudents[stuID]
	if !ok {
		fmt.Println("No Such Student")
		return
	}
	fmt.Printf("Student Infomation : StudentsID:%d,StudentName:%s\n", stuObj.id, stuObj.name)
	//获取新的学生名
	fmt.Print("Enter New StudentName:")
	fmt.Scanln(&stuName)
	//更新学生数据
	stuObj.name = stuName
	s.allstudents[stuID] = stuObj
	fmt.Println("Successully Modified")
}
func (s stuMgr) deleteStudent() {
	//获取用户输入
	var (
		stuID int64
	)
	fmt.Print("Enter StudentID To Delete:")
	fmt.Scanln(&stuID)
	//判断是否存在
	_, ok := s.allstudents[stuID]
	if !ok {
		fmt.Println("No Such Student")
		return
	}
	//更新学生数组
	delete(s.allstudents, stuID)
	fmt.Println("Successully Delete")
}