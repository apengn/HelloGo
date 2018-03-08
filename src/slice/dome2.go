package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse1(s)
	fmt.Println(s)
}

func reverse1(s []int) {
	s = append(s, 999)

	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}

}
