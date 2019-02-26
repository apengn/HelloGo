package bishi

import (
	"sort"
)

func HeBing(s ...[]int) (p []int) {
	p = make([]int, 0)
	for _, v := range s {
		p = append(p, v...)
	}
	sort.Ints(p)
	//fmt.Printf("%d\n",p)
	return
}
