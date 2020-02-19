package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	err := globalRDB.RDB.MSet("date", "2012.3.30", "time", "11:00 a.m.", "weather", "sunny").Err()
	if err != nil {
		panic(err)
	}
	val, err := globalRDB.RDB.MGet("date", "time", "weather").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
