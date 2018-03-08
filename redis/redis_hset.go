package main

import "github.com/garyburd/redigo/redis"

func main() {
	c,err:=redis.Dial("tcp","localhost:6379")
}
