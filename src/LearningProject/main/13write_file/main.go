package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("file open failed : %v\n", err)
		return
	}
	defer fileObj.Close()
	//wirte
	fileObj.Write([]byte("heizi shige shabi\n"))
	//write string
	fileObj.WriteString("黑子是个傻逼")
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("file open failed : %v\n", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("heizi shige shabi")
	wr.WriteString("黑子是个傻逼")	//将内容写到缓冲区
	wr.Flush()
}

func writeDemo3(){
	str:="hello world!"
	err:=ioutil.WriteFile("./xx.txt",[]byte(str),0666)
	if err!=nil {
		fmt.Printf("write file failed : %v",err)
		return
	}
}

func main() {
	// writeDemo1()
	//writeDemo2()
	writeDemo3()
}
