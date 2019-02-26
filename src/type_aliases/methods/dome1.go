package main

import "fmt"

type T1 struct {
}

type T3 = T1

func (t1 T1) say() {

}

type S struct {
	T1
	T3
}

func (t3 T3) greeting() (i int){
	fmt.Println("xxx.txt")
	i=0
	return i
}

func main() {

	var s S

	s.say()


	//var t1 T1
	//var t3 T3
	//
	//t1.say()
	//t1.greeting()
	//
	//
	//t3.say()
	//t3.greeting()
}
