package main

import (
	"context"
	"fmt"
	"time"
)

func HandleRequest(ctx context.Context) {
	go WriteDatabase(ctx)
	go WriteRedis(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest 完成")
			return
		default:
			fmt.Println("HandleRequest 继续")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase 完成")
			return
		default:
			fmt.Println("WriteDatabase 继续")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis 完成")
			return
		default:
			fmt.Println("WriteRedis 继续")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()

	//Just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

//docker run --restart=always -p 6379:6379 --name myredis -d redis:latest  --requirepass 123456
