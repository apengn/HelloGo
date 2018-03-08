package main

import (
	"net/http"
	"fmt"

	"os"
	//"io/ioutil"
)

func Print(content interface{})  {
	fmt.Println(content)
}
func main() {
	response,err:=http.Get("http://baidu.com/")

	if err !=nil {
		Print(err)
		os.Exit(1)
	}else {
		defer response.Body.Close()
		//contents,err:=ioutil.ReadAll(response.Body)
		//if err !=nil {
		//	Print(err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s\n",string(contents))
       buf:=make([]byte,1024)

       f,err:=os.OpenFile("baidu.html",os.O_RDWR|os.O_CREATE|os.O_APPEND,os.ModePerm)

		if err !=nil {
			Print(err)
			return
		}
		defer  f.Close()

		for{
			n,_:=response.Body.Read(buf)
			if n==0 {
				break
			}
			f.WriteString(string(buf[:n]))
		}

	}

}
