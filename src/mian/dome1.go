package main

import (
	//"sync"
	"fmt"
)

func main() {
	ch := make(chan int)
	//wg := sync.WaitGroup{}
	//wg.Add(20)
	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
			//wg.Done()
		}
	}()
	//go func() {
	//	for {
	//		select {
	//		case c := <-ch:
	//			fmt.Println(c)
	//		}
	//	}
	//}()

	for {
		select {
		case c := <-ch:
			fmt.Println(c)
			if c == 20 {
				break
			}
		}
	}

	//wg.Wait()
}
