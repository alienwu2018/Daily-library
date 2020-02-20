package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.RPush("languages", "python", "java")
	fmt.Println(globalRDB.RDB.LRange("languages", 0, -1).Val())
}
