package candy

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/lazygophers/utils/json"
)

// ToString 将任意类型转换为字符串
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
	switch x := val.(type) {
	case bool:
		if x {
			return "1"
		}
		return "0"
	case int:
		return fmt.Sprintf("%d", x)
	case int8:
		return fmt.Sprintf("%d", x)
	case int16:
		return fmt.Sprintf("%d", x)
	case int32:
		return fmt.Sprintf("%d", x)
	case int64:
		return fmt.Sprintf("%d", x)
	case uint:
		return fmt.Sprintf("%d", x)
	case uint8:
		return fmt.Sprintf("%d", x)
	case uint16:
		return fmt.Sprintf("%d", x)
	case uint32:
		return fmt.Sprintf("%d", x)
	case uint64:
		return fmt.Sprintf("%d", x)
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return strconv.FormatFloat(float64(x), 'f', 0, 32)
		}

		return strconv.FormatFloat(float64(x), 'f', 15, 32)
	case float64:
		if math.Floor(x) == x {
			return fmt.Sprintf("%.0f", x)
		}

		return strconv.FormatFloat(x, 'f', 6, 64)
	case time.Duration:
		return x.String()
	case string:
		return x
	case []byte:
		return string(x)
	case nil:
		return ""
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