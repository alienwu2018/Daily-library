package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.ZScore("page_rank", "bing.com").Val())
}
