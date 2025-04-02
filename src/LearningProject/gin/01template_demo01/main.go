package main

//template demo

import (
	"fmt"
	"html/template"
	"net/http"
	//"github.com/gin-gonic/gin"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//2、解析模板
		t, err := template.ParseFiles("./Hello.tmpl")
		if err != nil {
			fmt.Printf("Parse file failed,err:%v\n", err)
			return
		}
		//3、渲染模板
		name := "小王子"
		err = t.Execute(w, name)
		if err != nil {
			fmt.Printf("Execute file failed,err:%v\n", err)
			return
		}
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed,err:%v\n", err)
		return
	}
}
