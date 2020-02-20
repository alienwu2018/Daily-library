package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.SAdd("bbs", "discuz.net", "discuz.net")
	fmt.Println(globalRDB.RDB.SMembers("bbs").Val()) //{"discuz.net"}
}
