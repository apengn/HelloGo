package main

import "fmt"


//如果slice  在append的时候，  如果在append的之后slice的长度》slice的容量  则会重新分配一个slice  slice的默认容量是slice的长度+1
func main() {

	s:=make([]int,0,100)

	for i := 0; i < 7; i++ {
		s = append(s, i)
	}

	f(s)
	fmt.Println(s)

}

func f(s []int) {
	fmt.Println("cap", cap(s))
	s = s[:]
	s = append(s, 33,89)
	s[0] = 10
	fmt.Println(s)
}
