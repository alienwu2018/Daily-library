package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	b := globalRDB.RDB.HSet("website", "google", "www.gooole.cn").Val()
	fmt.Println(b) //true
	val := globalRDB.RDB.HGet("website", "google").Val()
	fmt.Println(val) //"www.google.cn"

	b1 := globalRDB.RDB.HSet("website", "google", "test").Val()
	fmt.Println(b1)
	val1 := globalRDB.RDB.HGet("website", "google").Val()
	fmt.Println(val1)
}
