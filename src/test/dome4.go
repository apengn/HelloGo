package main

import "fmt"

func main() {
	run()
}

func run() {
	b := false
	c := false
	var ch = make(chan bool, 1)
	go func() {
		if c {
			b= true
			ch <- true
			fmt.Println("xxxxxxx")
		}
	}()
	go func() {
		if b {
			ch <- true
			fmt.Println("xxxxxxx")
		}
	}()
	fmt.Println("1111")
	<-ch
	fmt.Println("end")
	return
}
