package main

import (
	"fmt"
	"math/rand"
	"sort"

	//"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	score := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		score[key] = value
	}
	fmt.Println(score)	//通常这里打印是无序的

	// 对map进行排序
	keys := make([]string, 0, 200)
	for key := range score {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, score[key])
	}

}
