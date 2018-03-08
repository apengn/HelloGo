package main

import (
	"time"
	"fmt"
)

func main() {
	ch:=make(chan int,2)

	go func() {

		time.Sleep(99*1e9)

		x:=<-ch
		fmt.Println("receviced",x)
	}()

	fmt.Println("send 10")
	ch<-10
	ch<-10
	ch<-10
	fmt.Println("sent 10")
}
