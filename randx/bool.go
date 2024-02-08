package randx

func Bool() bool {
	return Intn(2) == 0
}

func Booln(n float64) bool {
	if n >= 100 {
		return true
	} else if n <= 0 {
		return false
	}

	return Float64()*100 < n
}
