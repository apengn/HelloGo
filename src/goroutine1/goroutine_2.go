package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)


	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {

	ch <- "ONE"
	ch <- "TWO"
	ch <- "THREE"
	ch <- "FOUR"
	ch <- "FIVE"
}
func getData(ch chan string) {
	for {
		fmt.Printf("%s\n", <-ch)
	}
}
