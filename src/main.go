package main

import "fmt"

func main() {
	mySlice := make([]int, 100)

	copy(mySlice[3:5] ,[]int{3, 4})
	fmt.Println(mySlice[3:5])
}
