package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.HExists("counter", "page_view").Val())
	globalRDB.RDB.HIncrBy("counter", "page_view", 200)
	fmt.Println(globalRDB.RDB.HGet("counter", "page_view").Val())
}
