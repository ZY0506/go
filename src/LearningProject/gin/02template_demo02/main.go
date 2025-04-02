package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template demo

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse\n")
	}
	//渲染模板
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    18,
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server failed,err:%v\n", err)
		return
	}
}
