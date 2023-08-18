package redis

import (
	"fmt"
)

// 一些辅助函数和工具函数

func ValidateKey(key string) bool {
	// 校验键的有效性的逻辑
	fmt.Printf("Validating key: %s\n", key)
	return true
}

func ValidateValue(value string) bool {
	// 校验值的有效性的逻辑
	fmt.Printf("Validating value: %s\n", value)
	return true
}
