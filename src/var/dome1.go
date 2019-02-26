package main

import (
	"fmt"
	"strings"
)

//在常量组中，如不提供类型和初始化值，那么视作与上一常量相同
const (
	a = "abc"
	b
)

type P struct {
	NAME string
}

type s struct {
	P
	S string
}

func main() {

	//fmt.Println(b)
	//test333()
	//map1()

	m:=map[string]string{}

	delete(m,"xx")
	fmt.Println("delete success",m)
}

func test333() {
	var slice []int
	//slice = append(slice, 3, 5)
	fmt.Println(slice)
}

func map1() {
    //从left to right   删除头部连续的包含在cutset中的字符
	x := strings.TrimLeft("agb3ba", "ag3")
	fmt.Println(x)
	fmt.Print(strings.TrimSuffix("agb3ba","ba"))

}
