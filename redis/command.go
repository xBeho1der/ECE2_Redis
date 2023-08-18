// commands.go

package redis

import (
	"fmt"
)

func Set(key, value string, expiration int) {
	// 实现 Redis SET 命令的逻辑
	fmt.Printf("SET %s %s (expiration: %d)\n", key, value, expiration)
}

func Get(key string) (string, bool) {
	// 实现 Redis GET 命令的逻辑
	fmt.Printf("GET %s\n", key)
	// 假设这里是从数据库中获取数据的逻辑
	return "some_value", true
}

func Keys() []string {
	// 实现 Redis KEYS 命令的逻辑
	fmt.Println("KEYS")
	// 假设这里是从数据库中获取所有键的逻辑
	return []string{"key1", "key2", "key3"}
}
