package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.HSet("mykey", "field", "10.05")
	fmt.Println(globalRDB.RDB.HIncrByFloat("mykey", "field", 0.1).Val())

	globalRDB.RDB.HSet("mykey", "field", 5.0e3)
	fmt.Println(globalRDB.RDB.HIncrByFloat("mykey", "field", 2.0e2).Val())
}
