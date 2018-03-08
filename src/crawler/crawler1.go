package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
    //通过http.GET（url）来获取html源文件
	response, err := http.Get("http://baidu.com")
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
