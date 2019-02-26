package main

import (
	"runtime/debug"
	"sync/atomic"
	"sync"
	"fmt"
	"runtime"
)

func main() {

	f := fmt.Println
	// 禁用GC，并保证在main函数执行结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}

	pool := sync.Pool{New: newFunc}

	v1 := pool.Get()
	f("v1: ", v1)

	pool.Put(newFunc())
	pool.Put(newFunc())
	pool.Put(newFunc())
	v2 := pool.Get()
	f("v2:", v2)

	// 垃圾回收对临时对象池的影响

	debug.SetGCPercent(100)

	runtime.GC()
	v3 := pool.Get()
	f("v3:", v3)

	pool.New = nil
	v4 := pool.Get()
	f("v4:", v4)
}
