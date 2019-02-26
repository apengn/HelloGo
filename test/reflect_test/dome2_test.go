package main

import (
	"reflect"

	"fmt"

	"testing"

	"github.com/kr/pretty"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc(name string) {
	u.Name = name
	pretty.Println(name)
}

func Test1(t *testing.T) {
	user := User{Id: 1, Name: "test_name", Age: 5}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	pretty.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all value is", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		pretty.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 利用反射 获取方法并调用方法
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v \n", m.Name, m.Type)
		methodByName := getValue.MethodByName(m.Name)
		args := []reflect.Value{reflect.ValueOf("test methods")}
		methodByName.Call(args)
	}

}

// reflect.Value设置实际变量的值

// reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，即：要修改反射类型的对象就一定要保证其值是“addressable”的。
func Test2(t *testing.T) {
	var num float64 = 1.234
	fmt.Println("old value of pointer:", num)

	// pointer := reflect.ValueOf(num) //panic   call of reflect.Value.Elem on float64 Value
	pointer := reflect.ValueOf(&num)
	newValus := pointer.Elem()

	fmt.Println("type of pointer:", newValus.Type())

	fmt.Println("settability of pointer", newValus.CanSet())
	newValus.SetFloat(88)
	fmt.Println("new value of pointer:", newValus)

}
