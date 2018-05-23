package main

import (
	"time"
	"fmt"
)

func main() {


	c:=make(chan int)

	w:=make(chan int)

	go func() {

		select {
		case  <-c:
		case <-time.After(time.Second*3):
			fmt.Println("Time out")
			w<-2
		}
	}()


	<-w
}