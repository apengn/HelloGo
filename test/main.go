package main

import (
    "time"
    "fmt"
    "github.com/pkg/errors"
)

func main() {

    ch := make(chan int, 2)

    go func() {
        defer func() {
            if err := recover(); err != nil {

                fmt.Println(err.(error).Error())
                ch <- 2
                fmt.Println("wwwwwwwwwwwwwwwwww")
            }
        }()

        time.Sleep(19 * time.Second)
        ch <- 2
        panic(errors.New("xxxxxxxxxxxxxxxx"))
    }()

    go func() {
        time.Sleep(10 * time.Second)
        ch <- 1
    }()
    go func() {
        time.Sleep(17*time.Second)
        fmt.Println("2")
        ch<-1
        ch<-1

    }()

    <-ch
    fmt.Println("end")
    time.Sleep(200*time.Second)
}
