package main

import (
	"time"
	"fmt"
	"runtime"
	"sync"
)

//sync.Cond{}
//在Go语言中 sync.Cond 代表条件变量，但它需要配置锁才能有用.
func main() {
	runtime.GOMAXPROCS(4)
	f := fmt.Println
	var mutex sync.Mutex
	c := sync.NewCond(&mutex)


	condition := false
	go func() {
		time.Sleep(time.Second * 1)
		c.L.Lock()
		f("[1] 变更condition状态,并发出变更通知.")
		condition = true
		c.Signal() //单发通知  让条件变量向至少一个正在    等待通知的线程发送通知，表示共享数据的状态已经改变


		//c.Broadcast()  //广播通知  ， 让条件变量给正在等待他的通知的所有线程都发送 通知
		f("[1],继续后续处理")
		c.L.Unlock()
	}()

	c.L.Lock()
	f("[2],conditon ....1 ")
	for !condition {
		f("[2] condition ........2")
		c.Wait() //等待通知|信号
		f("[2] condition ......3")
	}

	f("[2] condtion ......4")
	c.L.Unlock()
	f("main end")

}
