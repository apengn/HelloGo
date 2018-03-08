package main

import (
	"unsafe"
	"strings"
)

func main() {
	var f float64 = 43
	Float64toBits(f)
}
func Float64toBits(f float64) uint64{
	floatPtr:=&f

	unsafePtr:=unsafe.Pointer(floatPtr)

	uintptr:=(*uint64)(unsafePtr)
	strings.Fields()

	uintVal:=*(uintptr)
	return  uintVal
}
