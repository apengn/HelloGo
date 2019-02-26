package main

import (
	"reflect"

	"github.com/kr/pretty"
)

// 已知原有类型【进行“强制转换”】
func main() {
	var num float64 = 1.234

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPoint := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	pretty.Println(convertPoint)
	pretty.Println(convertValue)
}
