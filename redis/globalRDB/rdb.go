package globalRDB

import (
	"github.com/go-redis/redis"
	"log"
)

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "47.100.93.236:6379",
		Password: "qq88227260",
		DB:       0,
	})
	if err := RDB.Ping().Err(); err != nil {
		log.Fatal(err)
	}
}
