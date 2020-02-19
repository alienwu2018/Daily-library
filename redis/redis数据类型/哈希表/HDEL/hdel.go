package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	//删除一个域
	val1 := globalRDB.RDB.HDel("phone", "myphone").Val()
	fmt.Println(val1) //1
	//删除多个域
	val2 := globalRDB.RDB.HDel("phone", "myphone", "yourphone").Val()
	fmt.Println(val2)
}
