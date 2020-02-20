package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.SAdd("dbs", "mysql", "mongodb", "redis").Err())
	fmt.Println(globalRDB.RDB.SPop("dbs").Val()) //随机一个val
}
