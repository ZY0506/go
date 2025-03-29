package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("Year:%d,Month:%d,Day:%d\n", now.Year(), now.Month(), now.Day())
	//时间戳
	fmt.Println(now.Unix())     //秒
	fmt.Println(now.UnixNano()) //纳秒
	//time.Unix()
	ret := time.Unix(1742287857, 0) //第一个参数为秒，第二个为纳秒
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Day())
	//now+1小时
	fmt.Println(now.Add(time.Hour))

	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer{
	// 	fmt.Println(t)
	// }

	//格式化时间
	fmt.Println(now.Format("2006-01-02"))

	fmt.Println(now.Format("2006/01/02"))

	fmt.Println("====================================================")

	//将时间解析为字符串
	timeObj, err := time.Parse("2006-01-02", "2005-06-08")
	if err != nil {
		fmt.Println("time parse failed,err:", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	//sub
	fmt.Println(now.Sub(timeObj))

	//sleep
	n := 100                     //int 类型
	time.Sleep(time.Duration(n)) //time.Duration类型 （int64）
	fmt.Println("begin")
	time.Sleep(5 * time.Second)
	fmt.Println("end")
}
func f2() {
	now := time.Now()
	fmt.Println(now)
	//明天的这个时间
	//按照指定格式解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2025-03-19 23:25:23")
	//按照东八区的时区和格式解析一个字符串格式的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err!=nil{
		fmt.Printf("load loction failed , err:%v\n",err)
		return
	}
	//按照指定时区和格式解析一个字符串格式的时间
	timeObj ,err :=time.ParseInLocation("2006-01-02 15:04:05", "2025-03-19 23:25:23", loc)
	if err!=nil{
		fmt.Printf("parse failed , err:%v\n",err)
		return
	}
	fmt.Println(timeObj)
	td := timeObj.Sub(now)
	fmt.Println(td)
}
func main() {
	f2()

}
