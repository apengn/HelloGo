package parser

import (
	"HelloGo/crawler/model"
	"fmt"
	"regexp"
	"encoding/json"
)

var RYear = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+?)</td>`)
var RHight = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+?)</td>`)
var RIncome = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+?)</td>`)
var RHun = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+?)</td>`)
var RXueLi = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+?)</td>`)
var RWorkPlace = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+?)</td>`)
var Rzhiye = regexp.MustCompile(`<td><span class="label">职业：</span>([^<]+?)</td>`)
var RStart = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+?)</td>`)
var RJiGuan = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+?)</td>`)

func ParseProfile(content []byte,name string) (res model.ParseResult) {
	var profile = model.Profile{}
	profile.RYear = GetSubData(RYear, content)
	profile.RName=name;
	profile.RYear = GetSubData(RHight, content)
	profile.RIncome = GetSubData(RIncome, content)
	profile.RHun = GetSubData(RHun, content)
	profile.RXueLi = GetSubData(RXueLi, content)
	profile.RWorkPlace = GetSubData(RWorkPlace, content)
	profile.Rzhiye = GetSubData(Rzhiye, content)
	profile.RStart = GetSubData(RStart, content)
	profile.RJiGuan = GetSubData(RJiGuan, content)
	json,err:=json.Marshal(profile)
	if err!=nil {
		fmt.Printf("json marshal faile err:%s\n",err)
	}
	fmt.Println(string(json))
	return
}

func GetSubData(r *regexp.Regexp, content []byte) (result string) {
	year := r.FindSubmatch(content)
	str := ""
	for _, y := range year {
		str = string(y)
	}
	result = str
	return
}
