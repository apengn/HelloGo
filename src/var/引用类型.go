package main

import (
	"fmt"
)

//引用类型包括：channel slice map

//内置函数 new()  计算类型大小，为其分配零值内存 返回指针
//make 会被编译器翻译成具体的函数，为其分配内存和初始化成员结构 返回对象而非指针
//类型转换   显示转换   不支持隐式转换
//字符串是不可变值类型  内部用指针指向UTF-8字节数组
//默认值是""
//用索引号访问某字节 如s[i]
//不能用序列号获取字节元素指针  如：&s[i] 非法
//不可变类型，无法修改字节数组
//字节数组尾部不能包含NULL

func main() {
	//s := "abc"
	//b := []byte(s)
	//b[0] = 'u'
	//
	//println(string(b))

	b := []int64{-2,2,-3,4,-1,2,1,-5,3}

	fmt.Println(MaxSubseqSum1(b,int64(len(b))))
}
func MaxSubseqSum1(list []int64, N int64) int64 {
	var ThisSum, MaxSum int64= 0, 0
	var i, j, k int64
	for i = 0; i < N; i++ {
		for j = i; j < N; j++ {

			ThisSum = 0
			for k = i; k <= j; j++ {
				ThisSum += list[k]
			}
			if ThisSum > MaxSum {
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}
