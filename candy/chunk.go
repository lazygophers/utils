package candy

// Chunk 将切片分割成指定大小的子切片
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
