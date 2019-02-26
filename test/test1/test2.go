package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	a := AddFunc(func(x, y int) (result int) {

		defer func(t time.Time) {
			log.Printf("took=%v, x=%v, y=%v, result=%v", time.Since(t), x, y, result)
		}(time.Now())
		return x + y
	})
	fmt.Println(Do(a))
}

type AddFunc func(x, y int) int

func (a AddFunc) Add(x, y int) int {
	return a(x, y)
}
func Do(adder Adder) int {
	return adder.Add(2, 3)
}

type Adder interface {
	Add(x, y int) int
}
