// Package utils 提供通用的工具函数和类型转换功能
package utils

import (
	"github.com/lazygophers/utils/candy"
)

// ToString 将任意类型转换为字符串
// 支持的类型包括：
// - 布尔值：true -> "1", false -> "0"
// - 整数类型：int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
// - 浮点数：float32, float64（自动处理精度）
// - 时间类型：time.Duration
// - 字符串：直接返回
// - 字节切片：转换为字符串
// - nil：返回空字符串
// - error：返回错误信息
// - 其他类型：使用 JSON 序列化
func ToString(val interface{}) string {
	return candy.ToString(val)
}
