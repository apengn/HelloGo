package rex

import (
	"fmt"
	"strconv"
	"testing"
)

func Test2(t *testing.T) {
	i, err := strconv.ParseInt("012345", 10, 64)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(i)
	println("xxxxxxxxxxxx")
	println("xxxxxxxxxxx1111111x")
	fmt.Println(999999)
}

func Test3(t *testing.T) {
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
