package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {

	m := pase_map()

	for k, v := range m {
		fmt.Printf("key = %s,value =%v\n", k, v)
	}

}

func pase_map() map[string]*student {

	m := make(map[string]*student)

	stu := []student{{"joy", 12}, {"lei", 14}}

	for _, v := range stu {
		fmt.Println(&v)
		m[v.Name] = &v
	}
	return m
}
