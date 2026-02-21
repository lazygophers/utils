package candy

// TypedSliceCopy 类型安全的切片复制
func TypedSliceCopy[T any](src []T) []T {
	if src == nil {
		return nil
	}

	dst := make([]T, len(src))

	// 检查是否是基本类型
	if len(src) > 0 {
		switch any(src[0]).(type) {
		case bool, string,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128:
			// 对于基本类型，使用快速复制
			copy(dst, src)
			return dst
		}
	}

	// 对于复杂类型，逐个深度复制
	for i := range src {
		DeepCopy(src[i], &dst[i])
	}
	return dst
}

// TypedMapCopy 类型安全的 map 复制
func TypedMapCopy[K comparable, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}

	dst := make(map[K]V, len(src))

	// 检查值类型是否是基本类型
	var sampleValue V
	switch any(sampleValue).(type) {
	case bool, string,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128:
		// 对于基本类型，直接复制
		for k, v := range src {
			dst[k] = v
		}
		return dst
	}

	// 对于复杂类型，深度复制值
	for k, v := range src {
		var dstV V
		DeepCopy(v, &dstV)
		dst[k] = dstV
	}
	return dst
}
