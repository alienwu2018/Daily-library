package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.LPush("course", "algorithm001", "c++101")
	fmt.Println(globalRDB.RDB.LPop("course").Val()) //c++101
}
