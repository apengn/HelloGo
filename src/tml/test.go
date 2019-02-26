package main

import "fmt"

type Test1 struct {
	is int
}

func main() {

	i := 5
	Test(func(test1 Test1) bool {
		return i == test1.is
	})
}

func Test(f func(Test1) bool) {
	t := Test1{2}
	fmt.Println(f(t))
}
