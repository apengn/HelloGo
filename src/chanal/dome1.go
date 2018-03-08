package main

import (
	"fmt"
	"go/ast"
	"sync"
)

//检查chan是否 关闭

type T int

type i=int



func IsClosed(ch <-chan T) bool {

	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)
	return true
}

func SafeSend(ch chan T, value T) (closed bool) {

	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value

	return false
}

type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func main() {
	c := make(chan T)

	fmt.Println(IsClosed(c))
	close(c)
	fmt.Println(IsClosed(c))
}
