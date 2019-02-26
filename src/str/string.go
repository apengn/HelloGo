package main

import (
    "fmt"
)

// func main() {
//     // v := "/a/b/c/"
//     // if strings.HasSuffix(v, "/") {
//     //     v = strings.TrimSuffix(v, "/")
//     // }
//     // fmt.Println(v)
//
//     s := fmt.Sprintf("%%%s%%", 433)
//     fmt.Println(string(s))
// }

type TestStruck struct {
    msg string
}

func (t *TestStruck) Read() {
    if len(t.msg) == 0 {
        fmt.Println("wwwwwwwwwwww")
        return
    }
    fmt.Println(t.msg)
}
func (t *TestStruck) Write(msg string) {
    t.msg = msg
}

func main() {
    // TestStruck{}.Read() // output error

    // struck := TestStruck{}
    // struck.Read()
    // var l *TestStruck

    var l *string
    if l == nil {
        fmt.Println("nil=======")
        return
    }

}
