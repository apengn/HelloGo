package main

import (
    "fmt"
)

type P struct {
    Name string
}

func main(){
    p:=p1()

    p=4

    fmt.Println(p1())
    fmt.Println(p)

}


func p1()int {


    return 2
}
