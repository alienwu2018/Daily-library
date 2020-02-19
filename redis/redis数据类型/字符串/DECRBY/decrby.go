package main

import (
	"everyday/redis/globalRDB"
	"fmt"
	"time"
)

func main() {
	globalRDB.RDB.Set("count", 100, 60*time.Second)
	globalRDB.RDB.DecrBy("count", 20)
	val := globalRDB.RDB.Get("count").Val()
	fmt.Println(val) //"80"
}
