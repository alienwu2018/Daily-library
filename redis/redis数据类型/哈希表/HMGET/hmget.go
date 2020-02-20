package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.HMGet("website", "google", "yahoo").Val())
}
