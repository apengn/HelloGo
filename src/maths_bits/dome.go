package main

import (
	bit "math/bits"
	 f "fmt"
)

func main() {
	var a uint=31

	f.Printf("bits.Len(%d)=%d\n",a,bit.Len(a))
	a++
	f.Printf("bits.Len(%d)=%d\n",a,bit.Len(a))

}
