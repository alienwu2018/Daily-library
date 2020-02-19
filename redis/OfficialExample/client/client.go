package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//client是redis的client，它代表一个单连接或者更加底层的连接，它能安全地同时给多个goroutine使用
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.100.93.236:6379", //连接地址
		Password: "qq88227260",         //连接密码
		DB:       0,                    //0号db
	})
	//redis set
	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	//redis get
	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	//#####################################
	val2, err := rdb.Get("missing_key").Result()
	if err == redis.Nil { //如果key不存在会返回nil
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
}
