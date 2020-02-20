package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.LLen("greet"))
	fmt.Println(globalRDB.RDB.LPushX("greet", "hello").Val()) //尝试 LPUSHX，失败，因为列表为空

	globalRDB.RDB.LPush("greet", "hello")
	globalRDB.RDB.LPushX("greet", "morning")
	fmt.Println(globalRDB.RDB.LRange("greet", 0, -1).Val()) //morning hello
}
