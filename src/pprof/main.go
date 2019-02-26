package main

import (
	_ "net/http/pprof"
	"net/http"
	"fmt"
)

//curl http://localhost:8899/debug/pprof/
func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Println(i, "xxxxxxxxxx")
		}
	}()

	func() {
		http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "xxxxxxxxxxxxxx")
		})
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()


}
