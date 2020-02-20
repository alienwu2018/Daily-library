package main

import "everyday/redis/globalRDB"

func main() {
	globalRDB.RDB.HMSet("website", map[string]interface{}{"google": "www.google", "yahoo": "www.yahoo.com"})
}
