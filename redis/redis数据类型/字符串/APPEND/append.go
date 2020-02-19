package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	//对不存在的key进行append
	l1, err := globalRDB.RDB.Append("fruits", "apple").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(l1) //l1=5

	//对已经存在的key进行append
	l2, err := globalRDB.RDB.Append("job", "s").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(l2) //l2=8  len(tester+s)=8
}
