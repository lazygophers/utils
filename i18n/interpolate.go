package i18n

import (
	"fmt"
	"strconv"
	"strings"
)

// interpolate 把 template 中 {name}/{idx} 占位符替换为 args 对应值，零分配快路径无 `{` 直接返回。
func interpolate(template string, args []any) string {
	// 快路径：无 `{` 直接返回原串，避免任何分配。
	if !hasBrace(template) {
		return template
	}
	if len(args) == 0 {
		return template
	}

	// 自动判别：args 非空、长度为偶数且首元素为 string → 命名模式；否则位置模式。
	if isNamed(args) {
		return interpolateNamed(template, args)
	}
	return interpolatePositional(template, args)
}

func hasBrace(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			return true
		}
	}
	return false
}

func isNamed(args []any) bool {
	if len(args)%2 != 0 {
		return false
	}
	_, ok := args[0].(string)
	return ok
}

// interpolateNamed 处理 `{name}` 占位符；args 形如 "k1", v1, "k2", v2。
func interpolateNamed(template string, args []any) string {
	var b strings.Builder
	b.Grow(len(template) + 16)

	n := len(template)
	i := 0
	for i < n {
		c := template[i]
		if c != '{' {
			b.WriteByte(c)
			i++
			continue
		}
		// 找匹配 '}'
		end := indexCloseBrace(template, i+1)
		if end < 0 {
			// 无闭合：剩余原样写出
			b.WriteString(template[i:])
			break
		}
		name := template[i+1 : end]
		val, ok := lookupNamed(args, name)
		if ok {
			writeValue(&b, val)
			i = end + 1
			continue
		}
		// 未命中：保留原样
		b.WriteString(template[i : end+1])
		i = end + 1
	}
	return b.String()
}

// interpolatePositional 处理 `{0}` 风格占位符。
func interpolatePositional(template string, args []any) string {
	var b strings.Builder
	b.Grow(len(template) + 16)

	n := len(template)
	i := 0
	for i < n {
		c := template[i]
		if c != '{' {
			b.WriteByte(c)
			i++
			continue
		}
		end := indexCloseBrace(template, i+1)
		if end < 0 {
			b.WriteString(template[i:])
			break
		}
		token := template[i+1 : end]
		idx, err := strconv.Atoi(token)
		if err == nil && idx >= 0 && idx < len(args) {
			writeValue(&b, args[idx])
			i = end + 1
			continue
		}
		// 非数字或越界：保留原样
		b.WriteString(template[i : end+1])
		i = end + 1
	}
	return b.String()
}

// indexCloseBrace 自 start 起寻找 `}` 的下标，未找到返回 -1。
func indexCloseBrace(s string, start int) int {
	for j := start; j < len(s); j++ {
		if s[j] == '}' {
			return j
		}
	}
	return -1
}

// lookupNamed 从 "k1",v1,"k2",v2 形式 args 中按 name 查值。
func lookupNamed(args []any, name string) (any, bool) {
	for i := 0; i+1 < len(args); i += 2 {
		k, ok := args[i].(string)
		if !ok {
			continue
		}
		if k == name {
			return args[i+1], true
		}
	}
	return nil, false
}

// writeValue 把任意值写入 builder，常见标量走快路径避免 fmt 反射开销。
func writeValue(b *strings.Builder, v any) {
	switch x := v.(type) {
	case string:
		b.WriteString(x)
	case int:
		b.WriteString(strconv.Itoa(x))
	case int64:
		b.WriteString(strconv.FormatInt(x, 10))
	case int32:
		b.WriteString(strconv.FormatInt(int64(x), 10))
	case uint:
		b.WriteString(strconv.FormatUint(uint64(x), 10))
	case uint64:
		b.WriteString(strconv.FormatUint(x, 10))
	case uint32:
		b.WriteString(strconv.FormatUint(uint64(x), 10))
	case bool:
		b.WriteString(strconv.FormatBool(x))
	case float64:
		b.WriteString(strconv.FormatFloat(x, 'g', -1, 64))
	case float32:
		b.WriteString(strconv.FormatFloat(float64(x), 'g', -1, 32))
	case []byte:
		b.Write(x)
	case fmt.Stringer:
		b.WriteString(x.String())
	default:
		fmt.Fprint(b, v)
	}
}
