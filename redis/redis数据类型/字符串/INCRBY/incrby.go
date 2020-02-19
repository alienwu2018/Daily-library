package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.Set("rank", 50, 0)
	globalRDB.RDB.IncrBy("rank", 20)
	val := globalRDB.RDB.Get("rank").Val()
	fmt.Println(val) //70
}
