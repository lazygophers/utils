package candy

// ToInt64Slice 将一个切片接口尽力转换为 []int64。
//
// 支持的输入切片类型包括：
//   - []bool, []int, []int8, ..., []uint64, []float32, []float64, []string, [][]byte, []interface{}
//
// 切片中的每一个元素都会通过 ToInt64 函数进行转换。
// 如果输入为 nil，将直接返回 nil。
// 如果输入为不支持的类型，将返回一个空的 []int64{}。
func ToInt64Slice(val interface{}) []int64 {
	switch x := val.(type) {
	case []bool:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	case []int:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int8:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int16:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int32:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int64:
		var v []int64
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint8:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint16:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint32:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint64:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []float32:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []float64:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []string:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	case [][]byte:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	case []interface{}:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	default:
		return []int64{}
	}
}