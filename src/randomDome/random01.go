package main

import "fmt"

func main() {
	ch := make(chan int)
	go printCh(ch)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
	}


}
func printCh(ch chan int) {
	for   {
		fmt.Println(<-ch)
	}
}
