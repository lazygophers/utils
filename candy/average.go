package candy


// Average 计算数值切片的平均值
func Average[T interface{ ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 }](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	var sum float64
	for _, s := range ss {
		sum += float64(s)
	}
	return T(sum / float64(len(ss)))
}
