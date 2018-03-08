package main

import "fmt"
//在常量组中，如不提供类型和初始化值，那么视作与上一常量相同
const (
	a="abc"
	b
)

func main() {


	fmt.Println(b)

}
