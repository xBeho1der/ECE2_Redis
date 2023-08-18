package cmd

import (
	"bufio"
	"fmt"
	"os"
	"redisdb/redis"
	"strings"
)

func Execute() {
	// 初始化 Redis 实例
	r := redis.NewRedis()

	// 创建一个命令行输入的循环
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command (or 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		// 解析和执行命令
		// 这里可以根据你的需求添加逻辑来解析和执行不同的命令

		switch input {
		case "get":
			value, exists := r.Get("key1")
			if exists {
				fmt.Println("Value of key1:", value)
			}
		case "keys":
			keys := r.Keys()
			fmt.Println("All keys:", keys)
		default:
			fmt.Println("Invalid command")
		}
	}

	// 关闭 Redis 实例
	r.Close()
	os.Exit(0)
}
