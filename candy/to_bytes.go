package candy

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"

	"github.com/lazygophers/utils/json"
)

// ToBytes 将任意类型转换为字节切片
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
func ToBytes(val interface{}) []byte {
	switch x := val.(type) {
	case bool:
		if x {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(fmt.Sprintf("%d", x))
	case int8:
		return []byte(fmt.Sprintf("%d", x))
	case int16:
		return []byte(fmt.Sprintf("%d", x))
	case int32:
		return []byte(fmt.Sprintf("%d", x))
	case int64:
		return []byte(fmt.Sprintf("%d", x))
	case uint:
		return []byte(fmt.Sprintf("%d", x))
	case uint8:
		return []byte(fmt.Sprintf("%d", x))
	case uint16:
		return []byte(fmt.Sprintf("%d", x))
	case uint32:
		return []byte(fmt.Sprintf("%d", x))
	case uint64:
		return []byte(fmt.Sprintf("%d", x))
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return []byte(strconv.FormatFloat(float64(x), 'f', 0, 32))
		}

		return []byte(strconv.FormatFloat(float64(x), 'f', 15, 32))
	case float64:
		if math.Floor(x) == x {
			return []byte(fmt.Sprintf("%.0f", x))
		}

		return []byte(strconv.FormatFloat(x, 'f', 6, 64))
	case time.Duration:
		return []byte(x.String())
	case string:
		return []byte(x)
	case []byte:
		return x
	case nil:
		return nil
	case error:
		return []byte(x.Error())

	default:
		buf, err := json.Marshal(x)
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
