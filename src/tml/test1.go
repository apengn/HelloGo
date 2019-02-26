package main

import (
    "github.com/robfig/cron"
    "fmt"
    "time"
)

type W struct {
    E string
}

func main() {
   go  test5()
   fmt.Println("xxxxxxxxx")
   time.Sleep(90000*time.Second)
    // tem := template.New("")
    // m := map[string]string{"E": "www"}
    // template.Must(tem.Parse("{{ $e := .E}}{{if $e}}{{$e}}{{else}}ppp{{end}}")).Execute(os.Stdout, m)
}

func test5() {

    cron := cron.New()
    cron.AddFunc("0/15 * * * * ? ", func() {
        fmt.Println(time.Now().Unix())
    })
    cron.Start()
}
