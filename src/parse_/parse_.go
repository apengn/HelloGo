package main

import (
	"go/token"
	"go/parser"
	"log"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fs := token.NewFileSet()

	f, err := parser.ParseFile(fs, "", "package main;var a=3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	spew.Dump(f)
}
