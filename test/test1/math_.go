package main

import (
    "unsafe"
    "fmt"
    "math"
)

func main() {

    var v float64 = 00434343
    u := (*uint64)(unsafe.Pointer(&v))
    fmt.Println(*u)
    fmt.Println(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&u))+unsafe.Sizeof(u))))
    fmt.Println(math.Float64frombits(math.Float64bits(v) &^ (1 << 63)))



}
