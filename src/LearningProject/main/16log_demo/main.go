package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("file open failed : %v\n", err)
		return
	}
	defer fileObj.Close()
	log.SetOutput(fileObj)
	for i := 0; i < 10; i++ {
		log.Printf("第%d条日志", i)
	}
}
