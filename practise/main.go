package main

import (
	"bytes"
	"fmt"
)
func main() {
	//定义四个chan
	chs := make([]chan int, 4)

	for i := 0; i < len(chs); i++ {
		//初始化chan
		chs[i] = make(chan int)

		go func(i int) {
			for {
				//写入数据
				chs[i] <- i + 1
			}
		}(i)
	}


	f := make([]bytes.Buffer, len(chs))//slice

	for i := 0; i < 10; i++ {


		for j := 0; j < len(f); j++ {


			k:=(i+j)%len(chs)
			fmt.Fprintf(&f[j], "%d ", <-chs[k])
		}
	}
	for i := 0; i < len(f); i++ {
		fmt.Printf("%d: %s\n", i, f[i].String())
	}
}
