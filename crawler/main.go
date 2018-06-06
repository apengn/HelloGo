package main

import (
	"HelloGo/crawler/fetcher"
	"HelloGo/crawler/zhenai/parser"
	"fmt"
)




func main() {

	b, err := fetcher.Fetcher("http://www.zhenai.com/zhenghun")

	if err != nil {
		fmt.Println(err)
		return
	}
	parser.ParseCity(b)

}
