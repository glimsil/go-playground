package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	client.Set("hello", "world", 0)

	val, err := client.Get("hello").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world getting value from key 'hello': ", string(val))
}
