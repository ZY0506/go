package main

import (
	"context"
	"fmt"
	"time"
)

//context.WithDeadline

func main() {
	d := time.Now().Add(5 * time.Second)
	// WithDeadline函数
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	//因为ctx 50毫秒后就会过期，所以ctx.Done()会先接收到context到期通知，并且会打印ctx.Err()的内容。

}
