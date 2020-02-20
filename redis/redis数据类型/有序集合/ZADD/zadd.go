package main

import (
	"everyday/redis/globalRDB"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println(globalRDB.RDB.ZAdd("page_rank", redis.Z{10.0, "google.com"}).Val())
	fmt.Println(globalRDB.RDB.ZAdd("page_rank", redis.Z{9.0, "baidu.com"}, redis.Z{8.0, "bing.com"}).Val())
	fmt.Println(globalRDB.RDB.ZRange("page_rank", 0, -1).Val())
}
