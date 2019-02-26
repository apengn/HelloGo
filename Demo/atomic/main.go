package main

import (
	"sync/atomic"
	"fmt"
	"sync"
)

func main() {
	f := fmt.Println

	//var ops int64 = 0
	//for i := 0; i < 50; i++ {
	//	atomic.AddInt64(&ops, 3)
	//	time.Sleep(time.Microsecond)
	//}
	//time.Sleep(time.Second)
	//f("opts: ", atomic.LoadInt64(&ops))
	//atomicAddOp(55)
	var pool = sync.Pool{New: func() interface{} {
		return "hello pool"
	}}

	val:="Hello World"
	pool.Put(val)

	f("------",pool.Get())
	f("===",pool.Get())   //再取就没有了,会自动调用NEW






}

var value int64

func atomicAddOp(tmp int64) {
	for {
		oldValue := value
		if atomic.CompareAndSwapInt64(&value, oldValue, oldValue+tmp) {
			return
		}
	}
}
