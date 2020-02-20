package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.ZIncrBy("page_rank", 100.0, "bing.com").Val())
}
