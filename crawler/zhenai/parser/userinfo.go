package parser

import (
	"HelloGo/crawler/fetcher"
	"HelloGo/crawler/model"
	"fmt"
	"regexp"
)

var rInfo = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`)

func ParseUserInfo(b []byte) (res model.ParseResult) {

	sub := rInfo.FindAllSubmatch(b, -1)

	var req = make([]model.Request, 10000)

	for _, v := range sub {
		name := string(v[2])
		r := model.Request{URL: string(v[1]), ParseFunc: func(bytes []byte) model.ParseResult {
			return ParseProfile(bytes, name)
		}}
		req = append(req, r)
		b, err := fetcher.Fetcher(r.URL)
		if err != nil {
			fmt.Println("fetcher failed Url", r.URL)
			continue
		}
		go r.ParseFunc(b)
	}
	res.Request = req
	return

}
