package candy

import "math/rand"

// Shuffle 随机打乱切片中的元素顺序
//
// 该函数使用 Fisher-Yates 洗牌算法对切片进行原地随机打乱。
// 对于每个元素，随机选择一个位置进行交换，确保每个元素都有均等的机会出现在任何位置。
//
// 参数:
//   - ss: 待打乱的切片，可以是任意类型
//
// 返回:
//   - []T: 打乱后的切片（原地修改，返回原切片的引用）
//
// 注意:
//   - 该函数会直接修改原切片，而不是创建新切片
//   - 对于空切片或单元素切片，函数会直接返回原切片
//   - 使用 math/rand 生成随机数，确保随机性
//
// 示例:
//   data := []int{1, 2, 3, 4, 5}
//   result := Shuffle(data)
//   // result 是打乱后的切片，与 data 是同一个切片
func Shuffle[T any](ss []T) []T {
	for i := range ss {
		j := rand.Intn(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}