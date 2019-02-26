package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	a := 1
	go func() {
		a=2
	}()
	a=3
	fmt.Println(a)
	time.Sleep(2*time.Second)
	log.Fatal()
}
