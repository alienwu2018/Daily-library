package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.Set("decimal", 3.0, 0)
	globalRDB.RDB.IncrByFloat("decimal", 2.56)
	val := globalRDB.RDB.Get("decimal").Val()
	fmt.Println(val) //"5.56"
}
