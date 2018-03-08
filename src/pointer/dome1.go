package main

import (
	"fmt"
)

func modify(arr [5]int) {
	(arr)[0] = 100
	return
}

func main() {
	var a [5]int
   b:=[8]int{1,1}

   c:=[...]int{2,2,22,9,9:32}
   fmt.Println(len(c))
   fmt.Println(b)
	modify(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
