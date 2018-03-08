package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck1(stream)
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck1(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
