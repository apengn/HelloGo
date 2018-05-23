package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}


	//&stu 是一个临时变量地址，地址会被重用，所以会出现map中的是最后一个slice 的值
	//range 值拷贝
	for _, stu := range stus {
		s:=stu
		m[stu.Name] = &s
	}
	fmt.Printf("%+v\n",m)
}

func main() {
	pase_student()
}