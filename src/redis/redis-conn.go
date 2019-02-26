package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {

	c, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println("redis connect fail")
		return
	}
	fmt.Println("redis connect sucess")

	defer c.Close()

	_, err = c.Do("set", "abc", 100)

	if err != nil {
		fmt.Println(err.Error())
	}
	r, err := c.Do("get", "abc")

	if err != nil {
		fmt.Println(err.Error())
	}
	rInt, err := redis.Int(r, err)
	fmt.Println("get redis abc value", rInt)

}
