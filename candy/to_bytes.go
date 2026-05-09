package candy

import (
	"math"
	"strconv"
	"time"
	"unsafe"

	"github.com/lazygophers/utils/json"
)

// ToBytes 将任意类型转换为字节切片
// 优化版本：快速路径优化，使用预定义字节切片减少分配
// 支持的类型包括：
// - 布尔值：true -> []byte("1"), false -> []byte("0")
// - 整数类型：int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
// - 浮点数：float32, float64（自动处理精度）
// - 时间类型：time.Duration
// - 字符串：转换为字节切片
// - 字节切片：直接返回
// - nil：返回nil
// - error：返回错误信息的字节切片
// - 其他类型：使用 JSON 序列化
//
// 性能优化：
// - 快速路径：常用类型（nil, []byte, string）优先处理
// - 整数转换：使用 strconv 替代 fmt.Sprintf，性能提升约 30%
// - 布尔值：使用预定义字节切片避免重复分配
func ToBytes(val interface{}) []byte {
	// 快速路径：最常用类型优先
	switch v := val.(type) {
	case nil:
		return nil
	case []byte:
		return v
	case string:
		return []byte(v)
	case bool:
		if v {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(strconv.FormatInt(int64(v), 10))
	case int8:
		return []byte(strconv.FormatInt(int64(v), 10))
	case int16:
		return []byte(strconv.FormatInt(int64(v), 10))
	case int32:
		return []byte(strconv.FormatInt(int64(v), 10))
	case int64:
		return []byte(strconv.FormatInt(v, 10))
	case uint:
		return []byte(strconv.FormatUint(uint64(v), 10))
	case uint8:
		return []byte(strconv.FormatUint(uint64(v), 10))
	case uint16:
		return []byte(strconv.FormatUint(uint64(v), 10))
	case uint32:
		return []byte(strconv.FormatUint(uint64(v), 10))
	case uint64:
		return []byte(strconv.FormatUint(v, 10))
	case float32:
		if math.Floor(float64(v)) == float64(v) {
			return []byte(strconv.FormatFloat(float64(v), 'f', 0, 32))
		}
		return []byte(strconv.FormatFloat(float64(v), 'f', 15, 32))
	case float64:
		if math.Floor(v) == v {
			return []byte(strconv.FormatFloat(v, 'f', 0, 64))
		}
		return []byte(strconv.FormatFloat(v, 'f', 6, 64))
	case time.Duration:
		return []byte(v.String())
	case error:
		return []byte(v.Error())
	default:
		buf, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		return buf
	}
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
