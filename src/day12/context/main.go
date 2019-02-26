package main

import "context"

func main() {



	context.WithValue(context.Background(),"xx","xxx.txt")
	context.Background().Value("xx")
}
