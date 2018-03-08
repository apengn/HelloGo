package main

import (
	"time"
	"fmt"
)

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		// 译注：注意这里将close函数调用注释掉了
		//close(ch)

	}()
	return ch
}
func main() {
	//wait :=
		Publish("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for the news...")
	// 译注：注意下面这一句
	//<-wait
	fmt.Println("The news is out, time to leave.")
}
