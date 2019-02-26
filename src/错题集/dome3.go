package main

import (
	"encoding/json"
	"fmt"
)
//Json反序列化数字到interface{}类型的值中，默认解析为float64类型

//使用 Golang 解析 JSON 格式数据时，若以 interface{} 接收数据，则会按照下列规则进行解析：
// bool, for JSON booleans
//float64, for JSON numbers
//string, for JSON strings
//[]interface{}, for JSON arrays
//map[string]interface{}, for JSON objects
//nil for JSON null
func main() {
	jsonStr:=`{"id":900,"name":"gdfdsfds"}`
	var maps map[string]interface{}
	json.Unmarshal([]byte(jsonStr),&maps)

	switch maps["id"].(type) {
	case float64:
		fmt.Println(maps["id"])
	}
}