package main

import (
	"fmt"
	"sync"
)

//sync.Map 并发安全的map

var (
	wg sync.WaitGroup
)
var m = make(map[int]int)
var m2 = sync.Map{}	//开箱即用，不需要初始化

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

// func main() {
// 	for i := 0; i < 50; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			set(i, i+100)
// 			fmt.Printf("key:%v,value:%v", i, get(i))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key:%v,value:%v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
