package main

import (
	"fmt"
	"io"
	"os"
)

func f() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("file open failed : %v\n", err)
		return
	}
	//新建一个文件
	tmpFile, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("create file failed : %v\n", err)
		return
	}
	//拷贝文件
	var ret [1]byte
	if _, err = fileObj.Read(ret[:]); err != nil {
		fmt.Printf("read file failed : %v\n", err)
		return
	}
	tmpFile.Write(ret[:])
	//将要插入的内容写入到新文件里面
	tmpFile.WriteString("hello world")
	//将原文件的剩余内容写入新文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file failed : %v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./tmp.txt", "./sb.txt")
}
func main() {
	f()
}
