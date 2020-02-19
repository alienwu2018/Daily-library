package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	globalRDB.RDB.Set("greeting", "hello, my friend", 0)
	val1 := globalRDB.RDB.GetRange("greeting", 0, 4).Val()
	fmt.Println(val1) //hello
	val2 := globalRDB.RDB.GetRange("greeting", -1, -5).Val()
	fmt.Println(val2) //"" 不支持回绕操作
	val3 := globalRDB.RDB.GetRange("greeting", -3, -1).Val()
	fmt.Println(val3) //end
	val4 := globalRDB.RDB.GetRange("greeting", 0, -1).Val()
	fmt.Println(val4) //hello, my friend
	val5 := globalRDB.RDB.GetRange("greeting", 0, 1008611).Val()
	fmt.Println(val5) //hello, my friend 值域范围不超过实际字符串，超过部分自动被符略
}
