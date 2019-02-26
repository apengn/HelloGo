package main

import "fmt"

//func Foo(x interface{}) {
//	if x == nil {
//		fmt.Println("empty interface")
//		return
//	}
//	fmt.Println("non-empty interface")
//}
//func main() {
//	var x *int = nil
//	Foo(x)
//	panic(func() {})
//}

type Persion struct {
	Name string
	Address
}

type Address struct {
	Addr string
}

func main() {

	p1 := Persion{}
	p1.Addr = "p1"
	p1.Name = "n1"

	p2 := Persion{}
	p2.Addr = "p2"
	p2.Name = "n2"

	changeAddr(&p1.Address, "change1")
	changeAddr(&p2.Address, "change1")

	fmt.Println(p1)
	fmt.Println(p1)

}

func changeAddr(a *Address, address string) {
	a.Addr = address
}
