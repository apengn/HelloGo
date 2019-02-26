package main

import (
	"context"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
)

type Result struct {
	r   *http.Response
	err error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	c := make(chan Result, 1)

	req, err := http.NewRequest("GET", "https://www.baidu.com/", nil)

	if err != nil {
		fmt.Println("http request failed err:", err)
		return
	}

	go func() {
		res, err := client.Do(req)
		pack := Result{res, err,
		}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		pack := <-c
		fmt.Println("TimeOut", pack.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)

		fmt.Printf("Server Response:%s", out)
	}

}
