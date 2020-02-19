package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	l1, err := globalRDB.RDB.StrLen("job").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(l1) //l1=7     teacher
	//当key不存在
	l2, err := globalRDB.RDB.StrLen("job233").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(l2) //0
}
