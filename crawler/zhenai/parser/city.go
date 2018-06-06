package parser

import (
	"HelloGo/crawler/fetcher"
	"HelloGo/crawler/model"
	"fmt"
	"regexp"
)

var r = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[A-Za-z0-9]+)"[^>]+>([^<]+)</a>`)

func ParseCity(b []byte) (res model.ParseResult) {
	//必须匹配
	matchs := r.FindAllSubmatch(b, -1)
	var requests []model.Request
	for _, m := range matchs {
		r := model.Request{URL: string(m[1]), ParseFunc: ParseUserInfo}
		requests = append(requests, r)
		d, err := fetcher.Fetcher(r.URL)
		if err != nil {
			fmt.Printf("fetcher failed,url:%s,err：%s", r.URL, err)
			continue
		}
		go r.ParseFunc(d)

	}
	res.Request = requests
	return
}
