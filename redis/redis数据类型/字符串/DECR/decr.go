package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.Set("failure_times", 10, 0)
	globalRDB.RDB.Decr("failure_times")
	val := globalRDB.RDB.Get("failure_times").Val()
	fmt.Println(val) //"9"
}
