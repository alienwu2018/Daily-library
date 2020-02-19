package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.HSet("db", "redis", "redis.com")
	globalRDB.RDB.HSet("db", "mysql", "mysql.com")
	l := globalRDB.RDB.HLen("db").Val()
	fmt.Println(l) //2
}
