package main

import (
	"reflect"

	"fmt"
)

type Ye struct {
	Size float64
}
type Tree struct {
	Ye
	Name string
	hight float64

}

func ( y Ye)Method1()  {
	fmt.Println("method1")
}
func (y Ye)Method2()  {
	fmt.Println("method2")
}

func ( y Tree)Method3()  {
	fmt.Println("method3")
}
func (y Tree)Method4()  {
	fmt.Println("method4")
}
func main() {

	defer func() {
		if err:=recover() ;err!=nil{
			fmt.Println(err)
		}

	}()
	t:=Tree{Ye{90},"golang",89}
  t.Method4()
	types:= reflect.TypeOf(t)
	kind:=types.Kind()
  fmt.Println(kind.String())

  fmt.Println(types.Name())
  fmt.Println(types.PkgPath())

  fmt.Println(types.NumField())

  fmt.Println(types.Field(1))
  //fmt.Println(types.FieldByIndex([]int{0,1}))

  fmt.Println(types.FieldByName(types.Field(2).Name))

	//switch types.(type) {
	//case string:
	//	fmt.Println("string")
	//default:
	//	fmt.Println("default")
	//
	//}
	fmt.Println("=================")
	fmt.Println(types.NumMethod())
    //fmt.Println(types.Method(0))
    //fmt.Println(types.MethodByName("method3"))

	v:=reflect.ValueOf(t)

	fmt.Println(v)
	fmt.Println(v.Kind())


	//fmt.Println(v.IsNil())


	//fmt.Println(v.Elem())

	fmt.Println(v.NumField())

	fmt.Println(v.Field(0))
	fmt.Println(v.FieldByName("Name"))

	fmt.Println(v.FieldByIndex([]int{0,0}))

	fmt.Println(v.NumMethod())
}
