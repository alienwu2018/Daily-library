package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	fmt.Println(globalRDB.RDB.HExists("phone", "myphone").Val()) //false

	globalRDB.RDB.HSet("phone", "myphone", "nokia-1110")
	fmt.Println(globalRDB.RDB.HExists("phone", "myphone").Val()) //true
}
