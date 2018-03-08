package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	cha2()
}
func cha1() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	for v := range ch {
		fmt.Println(v)
		if len(ch) == 0 {
			break
		}
	}
}
func cha2() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)  //close 了之后ch 变成只读的
	for v := range ch {
		fmt.Println(v)
	}
	runtime.GOMAXPROCS()
}
