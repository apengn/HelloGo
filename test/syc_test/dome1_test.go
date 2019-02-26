package syc_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 用来保存和复用临时对象，以减少内存分配，降低CG压力。
// 保存在Pool中的对象会在没有任何通知的情况下被自动移除掉。
// 实际上，这个清理过程是在每次垃圾回收之前做的。垃圾回收是固定两分钟触发一次。
// 而且每次清理会将Pool中的所有对象都清理掉！
func TestDome1(t *testing.T) {
	var pool = &sync.Pool{}
	var i = 999
	pool.Put(i)
	fmt.Println(pool.Get())
}

var i = 99

func TestDome1_1(t *testing.T) {
	var mutex = sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go func() {
		for {
			cond.L.Lock()
			time.Sleep(time.Second * 2)
			if i == 99 {
				fmt.Println("wait")
				cond.Wait()
			}
			i = 88
			cond.L.Unlock()
			cond.Signal()
			fmt.Println("wait end")
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			time.Sleep(time.Second * 2)
			if i == 88 {
				fmt.Println("wait1")
				cond.Wait()
			}

			i = 99
			cond.L.Unlock()
			cond.Signal()
			fmt.Println("wait end1")

		}

	}()

	time.Sleep(1000 * time.Minute)
}

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

var capacity = 10
var consumerNum = 3
var producerNum = 5

func producer(out chan<- int) {
	for i := 0; i < producerNum; i++ {
		go func(nu int) {
			for {
				cond.L.Lock()
				for len(out) == capacity {
					fmt.Println("Capacity Full, stop Produce")
					cond.Wait()
				}
				num := rand.Intn(100)
				out <- num
				fmt.Printf("Produce %d produce: num %d\n", nu, num)
				cond.L.Unlock()
				cond.Signal()

				time.Sleep(time.Second)
			}
		}(i)
	}
}

func consumer(in <-chan int) {
	for i := 0; i < consumerNum; i++ {
		go func(nu int) {

			for {
				cond.L.Lock()
				for len(in) == 0 {
					fmt.Println("Capacity Empty, stop Consume")
					cond.Wait()
				}
				num := <-in
				fmt.Printf("Goroutine %d: consume num %d\n", nu, num)
				cond.L.Unlock()
				time.Sleep(time.Millisecond * 500)
				cond.Signal()
			}
		}(i)
	}
}

func TestDome1_2(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	product := make(chan int, capacity)

	producer(product)
	consumer(product)

	<-quit
}
