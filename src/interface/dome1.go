package main

import (
	"fmt"
)

func init() {

}
func init() {

}

type I interface {
	set(int)
	get() int
}

type S struct {
	Age int
}

func (s *S) set(i int) {
	s.Age = i
}
func (s S) get() (age int) {
	return s.Age
}

func main() {

	s := S{}

	var i I
	i = &s
	i.set(10)
	fmt.Println(i.get())

	if t,ok:=i.(*S);ok {
		fmt.Println("xxxxxxxxxx",t)
	}


	switch t:=i.(type){
	case *S:
		fmt.Println("",t)


	}
}
