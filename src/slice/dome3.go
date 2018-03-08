package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	fmt.Println(cap(s))
	reverse5(s)
	fmt.Println(s)
}

func reverse2(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
func reverse5(s []int) {
	newElem := 999
	for len(s) < cap(s) {
		fmt.Println("Adding an element:", newElem, "cap:", cap(s), "len:", len(s))
		s = append(s, newElem)
		newElem++
	}
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}