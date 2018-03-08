package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {

	//设置可以执行的cpu的最大数量
	runtime.GOMAXPROCS(2)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
   fmt.Println(runtime.NumCPU(),"================")
	time.Sleep(1e9)

}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	if <-ch2==2000 {

	}

	for i := 0; ; i++ {
		select {
		case v := <-ch1:
			fmt.Printf("%d -Received on channel 1:%d\n", i, v)
		case v := <-ch2:
			fmt.Printf("%d -Received on channel 2:%d\n", i, v)
		}
	}

}
