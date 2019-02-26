package bishi

import "fmt"

//1、给定的字符串，找出其中最长的连续没有重复的字符串，并给出它的长度
func ExampleJudgeStr() {

	length, values := JudgeStr("aabcceddabc")

	fmt.Printf(`"%s" %d`, values, int(length))

	//Output:"dabc" 4

}

//合并两个升序排序的int的slice
//例如：
//输入：[1,3,4] [1,1,2,5] 输出：[1,1,1,2,3,4,5]
func ExampleHeBing() {
	v := HeBing([]int{1, 3, 4}, []int{1, 1, 2, 5})
	fmt.Printf("%d\n", v)

	//Output:[1 1 1 2 3 4 5]
}

//3、非负整数的slice，每一位上的数字代表它所能向前前进的最大步数，从第一位开始计算，判断是否能够到达最后一个节点

func ExampleHeBing3() {
	//没懂起意思，有点尴尬

}
