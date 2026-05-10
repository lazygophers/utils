package anyx

import (
	"strings"
	"testing"
)

// ============ 不同实现方案 ============

// 方案1: 当前实现（基于 strings.Builder）
func splitKeyCurrent(key string, sep string) []string {
	var parts []string
	current := new(strings.Builder)
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0
	endsWithSep := strings.HasSuffix(key, sep)

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			current.WriteByte(c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current.WriteByte(c)
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			if current.Len() > 0 || !afterBrackets {
				parts = append(parts, current.String())
			}
			current.Reset()
			i += sepLen - 1
			afterBrackets = false
		default:
			current.WriteByte(c)
			afterBrackets = false
		}
		i++
	}

	if current.Len() > 0 || endsWithSep {
		parts = append(parts, current.String())
	}

	return parts
}

// 方案2: 预分配切片优化（减少扩容）
func splitKeyPreAlloc(key string, sep string) []string {
	// 估算：假设平均每个部分 10 个字符
	estimatedParts := (len(key) + 9) / 10
	parts := make([]string, 0, estimatedParts)
	current := new(strings.Builder)
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0
	endsWithSep := strings.HasSuffix(key, sep)

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			current.WriteByte(c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current.WriteByte(c)
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			if current.Len() > 0 || !afterBrackets {
				parts = append(parts, current.String())
			}
			current.Reset()
			i += sepLen - 1
			afterBrackets = false
		default:
			current.WriteByte(c)
			afterBrackets = false
		}
		i++
	}

	if current.Len() > 0 || endsWithSep {
		parts = append(parts, current.String())
	}

	return parts
}

// 方案3: 使用字符串拼接（对比 baseline）
func splitKeyStringConcat(key string, sep string) []string {
	var parts []string
	var current string
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0
	endsWithSep := strings.HasSuffix(key, sep)

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
			current += string(c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current += string(c)
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			if current != "" || !afterBrackets {
				parts = append(parts, current)
			}
			current = ""
			i += sepLen - 1
			afterBrackets = false
		default:
			current += string(c)
			afterBrackets = false
		}
		i++
	}

	if current != "" || endsWithSep {
		parts = append(parts, current)
	}

	return parts
}

// 方案4: 手动字节切片优化（避免 Builder 开销）
func splitKeyByteSlice(key string, sep string) []string {
	var parts []string
	var current []byte
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0
	endsWithSep := strings.HasSuffix(key, sep)

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if len(current) > 0 {
				parts = append(parts, string(current))
				current = current[:0]
			}
			current = append(current, c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current = append(current, c)
			if len(current) > 0 {
				parts = append(parts, string(current))
				current = current[:0]
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			if len(current) > 0 || !afterBrackets {
				parts = append(parts, string(current))
			}
			current = current[:0]
			i += sepLen - 1
			afterBrackets = false
		default:
			current = append(current, c)
			afterBrackets = false
		}
		i++
	}

	if len(current) > 0 || endsWithSep {
		parts = append(parts, string(current))
	}

	return parts
}

// 方案5: 标准库 strings.Split（对比 baseline）
func splitKeyStringsSplit(key string, sep string) []string {
	// 注意：这个简化实现不完全等价，仅作性能对比
	return strings.Split(key, sep)
}

// 方案6: 预分配字节切片优化
func splitKeyByteSlicePreAlloc(key string, sep string) []string {
	estimatedParts := (len(key) + 9) / 10
	parts := make([]string, 0, estimatedParts)
	current := make([]byte, 0, 32) // 预分配 32 字节缓冲
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0
	endsWithSep := strings.HasSuffix(key, sep)

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if len(current) > 0 {
				parts = append(parts, string(current))
				current = current[:0]
			}
			current = append(current, c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current = append(current, c)
			if len(current) > 0 {
				parts = append(parts, string(current))
				current = current[:0]
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			if len(current) > 0 || !afterBrackets {
				parts = append(parts, string(current))
			}
			current = current[:0]
			i += sepLen - 1
			afterBrackets = false
		default:
			current = append(current, c)
			afterBrackets = false
		}
		i++
	}

	if len(current) > 0 || endsWithSep {
		parts = append(parts, string(current))
	}

	return parts
}

// 方案7: 去除 HasSuffix 预检查（内联到循环末尾）
func splitKeyInlineSuffix(key string, sep string) []string {
	var parts []string
	current := new(strings.Builder)
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			current.WriteByte(c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current.WriteByte(c)
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			if current.Len() > 0 || !afterBrackets {
				parts = append(parts, current.String())
			}
			current.Reset()
			i += sepLen - 1
			afterBrackets = false
		default:
			current.WriteByte(c)
			afterBrackets = false
		}
		i++
	}

	// 内联判断：最后添加当前部分或空字符串（如果以 sep 结尾）
	if current.Len() > 0 || (len(key) >= sepLen && key[len(key)-sepLen:] == sep) {
		parts = append(parts, current.String())
	}

	return parts
}

// 方案8: 简化状态机（减少分支）
func splitKeySimplified(key string, sep string) []string {
	var parts []string
	current := new(strings.Builder)
	sepLen := len(sep)
	i := 0

	for i < len(key) {
		// 检查分隔符（不在括号内）
		if i+sepLen <= len(key) && key[i:i+sepLen] == sep {
			// 检查前后是否都在括号外
			inBracketsBefore := false
			for j := 0; j < i; j++ {
				if key[j] == ']' && (j+1 >= i || key[j+1] != '[') {
					inBracketsBefore = false
				}
				if key[j] == '[' {
					inBracketsBefore = true
				}
			}
			if !inBracketsBefore {
				parts = append(parts, current.String())
				current.Reset()
				i += sepLen
				continue
			}
		}

		c := key[i]
		current.WriteByte(c)

		// 遇到 [ 或 ] 时完成当前部分
		if c == '[' || c == ']' {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		}

		i++
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}

// ============ Benchmark 场景 ============

// 场景1: 简单 key（点分隔）
func BenchmarkSplitKeySimpleDot(b *testing.B) {
	key := "user.profile.name"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("StringConcat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyStringConcat(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("StringsSplit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyStringsSplit(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
	b.Run("InlineSuffix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyInlineSuffix(key, sep)
		}
	})
	b.Run("Simplified", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeySimplified(key, sep)
		}
	})
}

// 场景2: 带数组索引
func BenchmarkSplitKeyWithArray(b *testing.B) {
	key := "data.items[0].user.profile[2].name"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("StringConcat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyStringConcat(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
	b.Run("InlineSuffix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyInlineSuffix(key, sep)
		}
	})
}

// 场景3: 不同分隔符（斜杠）
func BenchmarkSplitKeySlash(b *testing.B) {
	key := "api/v1/users/profile"
	sep := "/"
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景4: 不同分隔符（双冒号）
func BenchmarkSplitKeyDoubleColon(b *testing.B) {
	key := "namespace::class::method::field"
	sep := "::"
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景5: 深层嵌套
func BenchmarkSplitKeyDeepNesting(b *testing.B) {
	key := "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景6: 纯数组索引（无分隔符）
func BenchmarkSplitKeyPureArray(b *testing.B) {
	key := "matrix[0][1][2][3]"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景7: 长字符串
func BenchmarkSplitKeyLongString(b *testing.B) {
	key := "very_long_key_name_with_many_underscores.and.another.very_long_key_name_with_many_underscores.and.one.more"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景8: 单字符 key
func BenchmarkSplitKeySingleChar(b *testing.B) {
	key := "a"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景9: 边界 - 空字符串
func BenchmarkSplitKeyEmpty(b *testing.B) {
	key := ""
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景10: 真实复杂场景（API 响应路径）
func BenchmarkSplitKeyRealWorld(b *testing.B) {
	key := "data.results[0].user.profile.settings.preferences.theme"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景11: 连续分隔符
func BenchmarkSplitKeyConsecutiveSep(b *testing.B) {
	key := "a..b...c"
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}

// 场景12: 以分隔符开头/结尾
func BenchmarkSplitKeyStartEndWithSep(b *testing.B) {
	key := ".a.b.c."
	sep := "."
	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyCurrent(key, sep)
		}
	})
	b.Run("PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyPreAlloc(key, sep)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlice(key, sep)
		}
	})
	b.Run("ByteSlicePreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = splitKeyByteSlicePreAlloc(key, sep)
		}
	})
}
