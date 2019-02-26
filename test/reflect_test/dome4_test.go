package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDome4(t *testing.T) {

	var f float64 = 1.2234

	value := reflect.ValueOf(f)

	fmt.Println(value.Kind())

	fmt.Println(value.Float())

}

// v 的 Kind 仍然是 reflect.Int，尽管 x 的静态类型是 MyInt，而不是 int。换句话说，Kind 无法从 MyInt 中区分 int，而 Type 可以。
func TestDome4_1(t *testing.T) {
	type Myint int
	var x Myint = 9
	v := reflect.ValueOf(x)
	fmt.Println(v.CanSet())
	fmt.Println(reflect.TypeOf(x)) // Myint
	fmt.Println(v.Kind())          // int

}

type Test struct {
	Id   int64
	Name string
}

func TestDome4_2(t *testing.T) {

	ts := []Test{Test{Id: 4, Name: "xxxxxxxxxx"}}

	for _, v := range ts {
		i := reflect.ValueOf(v)
		fmt.Println(i.Type())
	}

}
