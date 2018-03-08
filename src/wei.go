package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	a := make(map[string]int)

	a["a"] = 1
	a["b"] = 2
	a["c"] = 3
	a["d"] = 4
	test(a)

}
func test(a map[string]int) {
	wg := sync.WaitGroup{}
	for k, v := range a {
		wg.Add(1)
		//go func(k string, v int) {
		//	wg.Done()
		//	check(k,v)
		//}(k, v)
		go func() {
			wg.Done()
			check(k,v)
		}()
		time.Sleep(1000)
	}
	wg.Wait()

}
func check(k string, v int) {
	fmt.Printf("K:%s,v:%d\n", k, v)
}
