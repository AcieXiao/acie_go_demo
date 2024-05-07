package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Process(ch chan int) {
	_time := rand.Intn(10) + 1
	fmt.Println(_time)
	time.Sleep(time.Second * time.Duration(_time))
	ch <- 1
}

func main() {
	channels := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}

	for i, channel := range channels {
		<-channel
		fmt.Println("routine ", i, " quit!")
	}

}
