package candy


// Abs 计算数值的绝对值
func Abs[T interface{ ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 }](s T) T {
	if s < 0 {
		return -s
	}

	return s
}
