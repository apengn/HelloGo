package model

var Citys = make([]City, 100)

type Request struct {
	URL       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}

type City struct {
	Url      string `json:"url"`
	CityName string `json:"city_name"`
	Users    []User `json:"users"`
}

type Profile struct {
	RName string `json:"r_name"`
	RYear      string `json:"r_year"`
	RHight     string `json:"r_hight"`
	RIncome    string `json:"r_income"`
	RHun       string `json:"r_hun"`
	RXueLi     string `json:"r_xue_li"`
	RWorkPlace string `json:"r_work_place"`
	Rzhiye     string `json:"rzhiye"`
	RStart     string `json:"r_start"`
	RJiGuan    string `json:"r_ji_guan"`
}

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Profile Profile `json:"profile"`
}
