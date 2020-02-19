package main

import (
	"everyday/redis/globalRDB"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	preval, err := globalRDB.RDB.GetSet("job", "teacher").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(preval) //preval=programmer
	//
	preval2, err := globalRDB.RDB.GetSet("job1", "py").Result()
	if err == redis.Nil {
		fmt.Println("nil") //nil
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(preval2)
	}

}
