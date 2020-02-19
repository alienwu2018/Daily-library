package main

import (
	"everyday/redis/globalRDB"
	"fmt"
	"time"
)

func main() {
	err := globalRDB.RDB.Set("test", "hello", 30*time.Second).Err() //生存30秒
	if err != nil {
		panic(err)
	}
	val1, err := globalRDB.RDB.Get("test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val1)
}
