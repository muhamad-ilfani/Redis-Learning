package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	var Option = redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&Option)
	file, _ := os.Open("data.json")
	data, _ := ioutil.ReadAll(file)

	err := client.Set("id1001", data, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("id1001").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
