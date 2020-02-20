package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.SIsMember("bbs", "test").Val())       //false
	fmt.Println(globalRDB.RDB.SIsMember("bbs", "discuz.net").Val()) //true
}
