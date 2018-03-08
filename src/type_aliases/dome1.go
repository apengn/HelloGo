package type_aliases

import (
	"fmt"
	"time"
)

type Type struct {
	Name string
}

//类型定义
type definituiiion int

type type_aliases = int

type T1=type_aliases

//类型别名定义语法
type identifer = Type

type F = func()

type G = interface {
}


func(t type_aliases)Dome(){

}
//类型定义  互相赋值  类型定义要和原始类型赋值的时候需要类型转换

func TypeDefine() {
	var v definituiiion = 90

	var v1 int

	v1 = int(v)
	fmt.Println(v1)

}

func typeAliases() {

	v := 100
	var d type_aliases = v
	var i definituiiion = definituiiion(v)

}

type T = time.Time

func main1() {

	var ty identifer = Type{"type aliases"}
	var g G = "HELLO TYPE ALIASE"
	var foo F = func() {
		fmt.Println(ty.Name)
		fmt.Println(g)
		var t T = time.Now()
		fmt.Println(t)
	}
	foo()



}

func swtichType(){

	var v interface{}

	var d type_aliases=100

	v=d
//既然类型别名和原类型相同，那么在switch-type中不能将 原类型和类型别名作为2个分支，因为这两个是重复的
	switch i:=v.(type) {
	case int:
	case type_aliases:

	}

}
