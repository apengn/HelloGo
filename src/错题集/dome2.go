package main

import (
	"bytes"
	"io"
	"fmt"
)

const debug = false


//https://yq.aliyun.com/articles/619547
func main() {
	//var buf *bytes.Buffer  有大坑
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}
func f(out io.Writer) {
	fmt.Println(out)
	if out != nil {
		fmt.Println("surprise!")
	}
}
