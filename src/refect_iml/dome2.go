package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City string
	Area string
}
type Student struct {
	Address
	Name string
	Age  int
}

func (this Student) Say() {
	fmt.Println("say")
}
func (this Student) Hello(word string) {
	fmt.Println(word)
}
func (this Student) HelloName(word string,name string) {
	fmt.Println(word,name)
}
func (this Student) WelCome(word ...string) {
	fmt.Println(word)
}

func main() {
	//defer func() {
	//	if erro := recover(); erro != nil {
	//		fmt.Println(erro)
	//	}
	//}()
	stu := Student{Address{"成都", "高新"}, "golang", 90}
	//ReflectSet(&stu)
	//StructInfo(stu)
	//Annoy(stu)
	ReflectMethod(stu)

}

func StructInfo(o interface{}) {


	t := reflect.TypeOf(o)

	fmt.Println(t.Name(), "object type: ", t.Name())
	//判断类型
	if k := t.Kind(); k != reflect.Struct {
		panic(fmt.Sprintln("this object is not a struct,but it is", k))
		return
	}

	//获取对象的值

	v := reflect.ValueOf(o)
	fmt.Println(t.Name(), "object value: ", v)

	//获取对象的字段

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v \n", f.Name, f.Type, val)

		//通过递归调用获取子类型的信息

		t1 := reflect.TypeOf(val)
		if k := t1.Kind(); k == reflect.Struct {
			StructInfo(val)
		}

	}

}

//匿名字段的反射
func Annoy(o interface{}) {

	t := reflect.TypeOf(o)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%10s:%#v %s\n", f.Name,f)
	}

}


//通过反射设置字段

func  ReflectSet(o interface{})  {

	v:=reflect.ValueOf(o)
	if k:=v.Kind();k==reflect.Ptr&&!v.Elem().CanSet() {
		fmt.Println("修改失败")
		return
	}
	v=v.Elem()

	//获取字段


	f:=v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("修改失败")
		return

	}

	if f.Kind()==reflect.String {
		f.SetString("java")
	}
}
//通过反射调用函数
func ReflectMethod(o interface{})  {
	v:=reflect.ValueOf(o)


	//无参数
	m1:=v.MethodByName("Say")
	m1.Call([]reflect.Value{})
    //有一个参数
	m2:=v.MethodByName("Hello")
	m2.Call([]reflect.Value{reflect.ValueOf("Hello")})

   //可变参数
	m3:=v.MethodByName("WelCome")
	m3.CallSlice([]reflect.Value{reflect.ValueOf([]string{"one","two","three"})})

     //多个参数
	m4:=v.MethodByName("HelloName")
	m4.Call([]reflect.Value{reflect.ValueOf("Hello"),reflect.ValueOf("Name")})

}

