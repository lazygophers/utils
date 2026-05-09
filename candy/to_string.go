package candy

import (
	"math"
	"strconv"
	"time"

	"github.com/lazygophers/utils/json"
)

// ToString 将任意类型转换为字符串
// 优化版本：快速路径优化，常见类型前置，减少 nil 检查开销
//
// 支持的类型包括：
// - 布尔值：true -> "1", false -> "0"
// - 整数类型：int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
// - 浮点数：float32, float64（自动处理精度）
// - 时间类型：time.Duration
// - 字符串：直接返回
// - 字节切片：转换为字符串
// - nil：返回空字符串
// - error：返回错误信息
// - 其他类型：使用 JSON 序列化
func ToString(val interface{}) string {
	switch v := val.(type) {
	case nil:
		return ""
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		if math.Floor(v) == v {
			return strconv.FormatFloat(v, 'f', -1, 64)
		}
		return strconv.FormatFloat(v, 'f', 6, 64)
	case bool:
		if v {
			return "1"
		}
		return "0"
	case []byte:
		return string(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		if math.Floor(float64(v)) == float64(v) {
			return strconv.FormatFloat(float64(v), 'f', -1, 32)
		}
		return strconv.FormatFloat(float64(v), 'f', 15, 32)
	case time.Duration:
		return v.String()
	case error:
		return v.Error()

	default:
		buf, err := json.Marshal(v)
		if err != nil {
			return ""
		}

		return string(buf)
	}
}
