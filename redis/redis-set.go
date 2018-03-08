package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {

	c, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		panic(fmt.Sprintln("链接失败", err.Error()))
	}
	defer c.Close()

	//在redis中设置值
	_, err = c.Do("set", "key1", "redis_set_value哈哈")

	if err != nil {
		panic(err.Error())
	}

	//重reids中获取值

	r,err:=redis.String(c.Do("get","key1"))

	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println("get key1 value ",r)

}
