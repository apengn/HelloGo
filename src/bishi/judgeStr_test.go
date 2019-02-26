package bishi

import (
	"testing"
	"fmt"
)

func TestJudegStr(t *testing.T) {
	length, values := JudgeStr("aabcceddabc")
	fmt.Printf(`"%s" %d`, values, int(length))
}

func TestHeBing(t *testing.T) {
	v := HeBing([]int{1, 3, 4}, []int{1, 1, 2, 5})
	fmt.Printf("%d\n", v)
}
