package main

import "fmt"

func main() {
	wait := make(chan struct{})
	n := 0
	go func() {
		// 译注：注意下面这一行
		n++ // 一次访问: 读, 递增, 写
		close(wait)
	}()
	// 译注：注意下面这一行
	n++ // 另一次冲突的访问
	<-wait
	fmt.Println(n) // 输出：未指定
}
