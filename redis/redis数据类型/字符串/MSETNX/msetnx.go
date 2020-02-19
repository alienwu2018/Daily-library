package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.MSetNX("rmdbs", "MySQL", "nosql", "MongoDB", "key-value-store", "redis")
	val := globalRDB.RDB.MGet("rmdbs", "nosql", "key-value-store").Val()
	fmt.Println(val)

	b := globalRDB.RDB.MSetNX("rmdbs", "Sqlite", "language", "python").Val()
	fmt.Println(b) //false
	val1 := globalRDB.RDB.Exists("language").Val()
	fmt.Println(val1) //0
	val2 := globalRDB.RDB.Get("rmdbs").Val()
	fmt.Println(val2) //MySQL
}
