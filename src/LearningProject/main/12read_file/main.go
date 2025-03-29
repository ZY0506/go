package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfilebyopen(){
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	//关闭文件
	defer file.Close()
	for {
		var tmp [128]byte
		n, err := file.Read(tmp[:])
		if(err==io.EOF){
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read file failed:%v\n", err)
			return
		}
		fmt.Printf("读了%v个字节\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

//bufio 读文件
func readfilebybufio(){
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	//关闭文件
	defer file.Close()

	reader:=bufio.NewReader(file)
	for {
		line,err:=reader.ReadString('\n')
		if(err==io.EOF){
			fmt.Println("读完了")
			return
		}
		if(err!=nil){
			fmt.Printf("read file failed:%v\n",err)
			return
		}
		fmt.Print(line)
	}
}

func readfilebyiotil(){
	ret,err := ioutil.ReadFile("./main.go")
	if err!=nil{
		fmt.Printf("read file failed : %v\n",err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	//readfilebyopen()
	//readfilebybufio()
	readfilebyiotil()
}
