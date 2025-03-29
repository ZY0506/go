package main

import (
	"fmt"
	"io"
	"net/http"
)

//net/http client

func main() {
	resp, err := http.Get("http://10.61.38.244:9090/xxx/?name=sb&age=18")
	if err != nil {
		fmt.Printf("get url failed,err:%v\n", err)
		return
	}
	//从resp中把服务器返回的数据读出来
	b,err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Printf("read resp.Body failed,err:%v\n",err)
		return
	}
	fmt.Println(string(b))
}
