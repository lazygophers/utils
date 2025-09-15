package candy

import (
	"cmp"
)

// Package candy 提供常用的工具函数和语法糖
// 本文件包含数组/切片差异比较相关的函数

// Diff 计算两个有序切片之间的差异
//
// 参数:
//   - ss: 第一个切片
//   - against: 第二个切片
//
// 返回值:
//   - added: 在 against 中存在但不在 ss 中的元素
//   - removed: 在 ss 中存在但不在 against 中的元素
//
// 示例:
//
//	ss := []int{1, 2, 3}
//	against := []int{2, 3, 4}
//	added, removed := Diff(ss, against)
//	// added = [4]
//	// removed = [1]
func Diff[T cmp.Ordered](ss []T, against []T) (added, removed []T) {
	removed = Remove(ss, against)
	added = Remove(against, ss)

	return
}
