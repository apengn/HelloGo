package bishi

import (
	"math"
)

func JudgeStr(str string) (count int, values string) {
	length := len(str)
	if length == 0 {
		return 0, ""
	}
	runne := []rune(str)
	maps := make(map[rune]int)
	var strRune = make([]rune, 0)
	var strRunes = make([][]rune, 0)
	var max float64
	for i, j := 0, 0; i < length; i++ {
		item := runne[i]
		k, ok := maps[item]

		if ok {
			j = int(math.Max(float64(j), float64(k+1)))
		}

		maps[item] = i

		strRune = append(strRune, item)

		max = math.Max(float64(max), float64(i-j+1))
		//fmt.Printf("j:%d,max:%f  size:%d\n", j, max, i-j+1)
		//
		//fmt.Println(string(strRune))
		//fmt.Printf("===%s\n", string(strRune[j:int(i)+1]))
		strRunes = append(strRunes, strRune[j:int(i)+1])
	}

	for _, v := range strRunes {
		if int(max) == len(v) {
			values = string(v)
		}
	}
	count=int(max)
	return
}
