package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	str, err := os.ReadFile("xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(str)
}

func f2(w http.ResponseWriter, r *http.Request) {
	//对于get请求，参数都放在url上（query param），请求体时没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(io.ReadAll(r.Body)) //我在服务器打印客户端发来的请求body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("10.61.38.244:9090", nil)

}
