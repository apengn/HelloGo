package main

import (
	"fmt"
	"testing"
)

type Writer interface {
	Writer()
}
type Reader interface {
	Reader()
}

type File struct {
}

func (File) Writer() {
	fmt.Println("writer")
}

func (File) Reader() {
	fmt.Println("reader")
}

// 接口类型转换
func TestType(t *testing.T) {
	var w Writer
	f := File{}
	//interface  w 中将包含（value,type）
	w = f
	// w 中将包含（f,File）
	switch t := w.(type) {
	case Reader:
		t.Reader()
	}

}
