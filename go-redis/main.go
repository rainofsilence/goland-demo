package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("connect redis error: ", err)
		return
	}
	defer c.Close()
	fmt.Println("connect redis success ")

	_, err = c.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set err: ", err)
		return
	}
	fmt.Println("set success")

	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err: ", err)
		return
	}
	fmt.Println(r)
	fmt.Println("set success")
}
