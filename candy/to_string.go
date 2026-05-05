package candy

import (
	"math"
	"strconv"
	"time"

	"github.com/lazygophers/utils/json"
)

// ToString 将任意类型转换为字符串
// 优化版本：将最常用的类型放在前面，减少类型分支开销
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
	// 快速路径：nil 检查
	if val == nil {
		return ""
	}

	switch x := val.(type) {
	// 常见类型优先
	case string: // 最常见 - 直接返回，零拷贝
		return x
	case int: // 常见
		return strconv.FormatInt(int64(x), 10)
	case int64: // 常见
		return strconv.FormatInt(x, 10)
	case float64: // 常见
		if math.Floor(x) == x {
			return strconv.FormatFloat(x, 'f', -1, 64)
		}
		return strconv.FormatFloat(x, 'f', 6, 64)
	case bool: // 常见
		if x {
			return "1"
		}
		return "0"
	case []byte: // 常见
		return string(x)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case uint:
		return strconv.FormatUint(uint64(x), 10)
	case uint8:
		return strconv.FormatUint(uint64(x), 10)
	case uint16:
		return strconv.FormatUint(uint64(x), 10)
	case uint32:
		return strconv.FormatUint(uint64(x), 10)
	case uint64:
		return strconv.FormatUint(x, 10)
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return strconv.FormatFloat(float64(x), 'f', -1, 32)
		}
		return strconv.FormatFloat(float64(x), 'f', 15, 32)
	case time.Duration:
		return x.String()
	case error:
		return x.Error()

	default:
		buf, err := json.Marshal(x)
		if err != nil {
			return ""
		}

		return string(buf)
	}
}
