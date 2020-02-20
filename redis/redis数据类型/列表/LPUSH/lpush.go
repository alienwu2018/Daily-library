package main

import (
	"everyday/redis/globalRDB"
	"fmt"
)

func main() {
	//列表允许重复元素
	fmt.Println(globalRDB.RDB.LPush("languages", "python", "java", "go").Val())
}
