package candy

// Chunk 将切片分割成指定大小的子切片
//
// 参数:
//   - ss: 要分割的切片
//   - size: 每个子切片的大小
//
// 返回:
//   - [][]T: 分割后的子切片集合
//
// 注意:
//   - 如果输入切片为空或 size <= 0，返回空切片
//   - 最后一个子切片可能小于指定大小
func Chunk[T any](ss []T, size int) (ret [][]T) {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}

	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}

		ret = append(ret, ss[i:end])
	}

	return
}
