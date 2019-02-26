package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID int `json:"id"`
}
type Address struct {
	Person
	AddressId int `json:"address_id"`
}

func main() {

	a := Address{Person{3}, 999}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))

}
