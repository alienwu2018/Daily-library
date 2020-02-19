package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	//判断key是否存在
	flag, err := globalRDB.RDB.Exists("job").Result()
	if err != nil {
		panic(err)
	}
	if flag == 0 { //若不存在插入
		err = globalRDB.RDB.SetNX("job", "programmer", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	val, err := globalRDB.RDB.SetNX("job", "code-farmer", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val) //val=false  尝试覆盖job,失败
	job, err := globalRDB.RDB.Get("job").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(job) //job=programmer
}
