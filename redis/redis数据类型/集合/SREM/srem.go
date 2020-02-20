package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.SRem("dbs", "redis").Val()) //1
}
