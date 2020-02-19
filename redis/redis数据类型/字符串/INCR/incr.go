package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.Set("page_view", 20, 0)
	globalRDB.RDB.Incr("page_view")
	val := globalRDB.RDB.Get("page_view").Val()
	fmt.Println(val) //"21"
}
