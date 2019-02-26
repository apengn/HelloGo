package main

import (
	"sync"
	"time"
	"fmt"
)

func TestTicker(wg sync.WaitGroup){

	calDuration:= func(duration time.Duration)time.Duration {
		now:=time.Now()

		return now.Truncate(duration).Add(duration).Sub(now)
	}

	fmt.Println(calDuration(time.Second*5))



}

func main()  {
	 TestTicker(sync.WaitGroup{})
}
