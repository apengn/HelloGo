package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := make([]byte, 0, 90)
	b = append(b, []byte("xxxxx")...)
	buffer := bytes.Buffer{}

	buffer.WriteString("wwwwwwwwwwwwwww")
	buffer.Write(b)
	fmt.Println(buffer.String())


	newBuffer:=bytes.NewBuffer(b)
	newBuffer.Write(b)
	// newBuffer.ReadBytes()
fmt.Println(newBuffer.String())
}
