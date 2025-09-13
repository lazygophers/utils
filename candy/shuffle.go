// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

import "github.com/lazygophers/utils/randx"

// Shuffle 随机打乱切片中的元素顺序
//
// 类型参数：
//   - T: 任意类型
//
// 参数：
//   - ss: 待打乱的切片，可以是任意类型
//
// 返回值：
//   - []T: 打乱后的切片（原地修改，返回原切片的引用）
//
// 特点：
//   - 使用 Fisher-Yates 洗牌算法，确保均匀随机分布
//   - 原地修改，不创建新切片，内存效率高
//   - 支持任意类型的切片
//   - 对于空切片或单元素切片，直接返回原切片
//   - 高性能优化：使用 randx 包的高性能随机数生成器
//
// 示例：
//
//	// 打乱整数切片
//	data := []int{1, 2, 3, 4, 5}
//	result := Shuffle(data)
//	// result 是打乱后的切片，与 data 是同一个切片
//
//	// 打乱字符串切片
//	names := []string{"Alice", "Bob", "Charlie", "David"}
//	shuffled := Shuffle(names)
//	// shuffled 包含随机顺序的名字
func Shuffle[T any](ss []T) []T {
	if len(ss) <= 1 {
		return ss
	}

	for i := len(ss) - 1; i > 0; i-- {
		j := randx.FastIntn(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}
