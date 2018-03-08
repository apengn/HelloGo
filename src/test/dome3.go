package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

func main() {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//sum += 1                 //1
			atomic.AddUint32(&sum, 1)  //2
			fmt.Println(sum)
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
