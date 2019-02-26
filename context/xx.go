package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*println(&x, x)
	  x := "abc"
	  println(&x, x)*/
	x, _ := strconv.Atoi("12")
	i, _ := strconv.ParseInt("12345", 8, 64)

	println("x:", x)
	println("i:", i)
	println(strconv.FormatBool(true))
	fmt.Println(strconv.ParseBool("true"))

	fmt.Println("-------------------")
}
