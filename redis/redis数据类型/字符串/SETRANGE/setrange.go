package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	err := globalRDB.RDB.Set("greeting", "hello world", 0).Err()
	if err != nil {
		panic(err)
	}
	l1, err := globalRDB.RDB.SetRange("greeting", 6, "Redis").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(l1) //11
	val1, err := globalRDB.RDB.Get("greeting").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val1) //"hello Redis"

	//对空白字符串/不存在的键执行setrange
	globalRDB.RDB.SetRange("empty_string", 5, "Redis!") //不存在的key
	l2 := globalRDB.RDB.StrLen("empty_string").Val()
	fmt.Println(l2) //l2=11
	val2 := globalRDB.RDB.Get("empty_string").Val()
	fmt.Println(val2) //val2=\x00\x00\x00\x00\x00Redis!
}
