package main

import (
	"bytes"
	"time"
	"fmt"
	"strings"
)

//字符串连接的3中方式

func main() {

	//1  最快

	var buffer bytes.Buffer
	s := time.Now()
	for i := 0; i < 10000; i++ {

		buffer.WriteString("test is here\n")

	}

	buffer.String()
	e := time.Now()

	fmt.Println("1 time is ", e.Sub(s).Seconds())

	//2

	s = time.Now()
	var s1 []string
	for i := 0; i < 100000; i++ {
		s1 = append(s1, "test is here\n")
	}
	strings.Join(s1, "")

	e = time.Now()

	fmt.Println("2 time is ", e.Sub(s).Seconds())

}
