package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(1)
		fmt.Println("协程1完成")
		wg.Done()
	}()

	go func() {
		time.Sleep(1)
		fmt.Println("协程2完成")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("都完成")

}
