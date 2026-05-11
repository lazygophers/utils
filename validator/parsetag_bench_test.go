package validator

import (
	"strings"
	"testing"
)

// ========== 当前实现（Baseline） ==========

func parseTagBaseline(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案1：预分配切片 ==========

func parseTagOpt1(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案2：IndexByte ==========

func parseTagOpt2(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案3：预分配 + IndexByte ==========

func parseTagOpt3(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案4：批量处理 ==========

func parseTagOpt4(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案5：手动 trim ==========

func parseTagOpt5(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		start := 0
		for start < len(part) && part[start] == ' ' {
			start++
		}
		if start >= len(part) {
			continue
		}

		end := len(part)
		for end > start && part[end-1] == ' ' {
			end--
		}

		trimmed := part[start:end]

		if idx := strings.IndexByte(trimmed, '='); idx != -1 {
			ruleName := trimmed[:idx]
			param := trimmed[idx+1:]

			paramStart := 0
			for paramStart < len(param) && param[paramStart] == ' ' {
				paramStart++
			}
			paramEnd := len(param)
			for paramEnd > paramStart && param[paramEnd-1] == ' ' {
				paramEnd--
			}

			rules = append(rules, validationRule{
				tag:   ruleName,
				param: param[paramStart:paramEnd],
			})
		} else {
			rules = append(rules, validationRule{tag: trimmed, param: ""})
		}
	}

	return rules
}

// ========== 优化方案6：单次遍历 ==========

func parseTagOpt6(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	start := 0
	inTag := true
	var ruleStart, ruleEnd int
	var paramStart, paramEnd int

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		if inTag {
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == '=' || ch == ',' || i == len(tag) {
				ruleStart = start
				ruleEnd = i

				for ruleEnd > ruleStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				if ch == '=' {
					inTag = false
					start = i + 1
				} else {
					if ruleStart < ruleEnd {
						rules = append(rules, validationRule{
							tag:   tag[ruleStart:ruleEnd],
							param: "",
						})
					}
					start = i + 1
				}
			}
		} else {
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == ',' || i == len(tag) {
				paramStart = start
				paramEnd = i

				for paramEnd > paramStart && tag[paramEnd-1] == ' ' {
					paramEnd--
				}

				rules = append(rules, validationRule{
					tag:   tag[ruleStart:ruleEnd],
					param: tag[paramStart:paramEnd],
				})

				inTag = true
				start = i + 1
			}
		}
	}

	return rules
}

// ========== 优化方案7：优化 Index 追踪 ==========

func parseTagOpt7(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	start := 0
	eqIndex := -1
	trimStart := 0

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		if ch == '=' {
			eqIndex = i
			continue
		}

		if ch == ',' || i == len(tag) {
			for trimStart < i && tag[trimStart] == ' ' {
				trimStart++
			}
			if trimStart >= i {
				start = i + 1
				trimStart = start
				eqIndex = -1
				continue
			}

			end := i
			for end > trimStart && tag[end-1] == ' ' {
				end--
			}

			if eqIndex != -1 && eqIndex < end {
				ruleEnd := eqIndex
				for ruleEnd > trimStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				paramStart := eqIndex + 1
				for paramStart < end && tag[paramStart] == ' ' {
					paramStart++
				}

				rules = append(rules, validationRule{
					tag:   tag[trimStart:ruleEnd],
					param: tag[paramStart:end],
				})
			} else {
				rules = append(rules, validationRule{
					tag:   tag[trimStart:end],
					param: "",
				})
			}

			start = i + 1
			trimStart = start
			eqIndex = -1
		}
	}

	return rules
}

// ========== 优化方案8：混合方案 ==========

func parseTagOpt8(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		start := 0
		for start < len(part) && part[start] == ' ' {
			start++
		}
		if start >= len(part) {
			continue
		}

		end := len(part)
		for end > start && part[end-1] == ' ' {
			end--
		}

		trimmed := part[start:end]

		if idx := strings.IndexByte(trimmed, '='); idx != -1 {
			ruleName := trimmed[:idx]
			param := trimmed[idx+1:]

			paramStart := 0
			for paramStart < len(param) && param[paramStart] == ' ' {
				paramStart++
			}
			paramEnd := len(param)
			for paramEnd > paramStart && param[paramEnd-1] == ' ' {
				paramEnd--
			}

			rules = append(rules, validationRule{
				tag:   ruleName,
				param: param[paramStart:paramEnd],
			})
		} else {
			rules = append(rules, validationRule{tag: trimmed, param: ""})
		}
	}

	return rules
}

// ========== 优化方案9：最简优化 ==========

func parseTagOpt9(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		idx := strings.IndexByte(part, '=')
		if idx == -1 {
			rules = append(rules, validationRule{tag: part, param: ""})
			continue
		}

		ruleName := part[:idx]
		param := part[idx+1:]

		trimRuleEnd := len(ruleName)
		for trimRuleEnd > 0 && ruleName[trimRuleEnd-1] == ' ' {
			trimRuleEnd--
		}

		trimParamStart := 0
		for trimParamStart < len(param) && param[trimParamStart] == ' ' {
			trimParamStart++
		}

		trimParamEnd := len(param)
		for trimParamEnd > trimParamStart && param[trimParamEnd-1] == ' ' {
			trimParamEnd--
		}

		rules = append(rules, validationRule{
			tag:   ruleName[:trimRuleEnd],
			param: param[trimParamStart:trimParamEnd],
		})
	}

	return rules
}

// ========== 优化方案10：完全手动解析 ==========

func parseTagOpt10(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	i := 0
	n := len(tag)

	for i < n {
		// 跳过前导空格和逗号
		for i < n && (tag[i] == ' ' || tag[i] == ',') {
			i++
		}
		if i >= n {
			break
		}

		start := i
		eqPos := -1

		// 查找规则结束位置（逗号）和等号位置
		for i < n && tag[i] != ',' {
			if tag[i] == '=' && eqPos == -1 {
				eqPos = i
			}
			i++
		}

		end := i
		// trim尾部空格
		for end > start && tag[end-1] == ' ' {
			end--
		}

		if eqPos != -1 && eqPos < end {
			// 有参数
			ruleEnd := eqPos
			for ruleEnd > start && tag[ruleEnd-1] == ' ' {
				ruleEnd--
			}

			paramStart := eqPos + 1
			for paramStart < end && tag[paramStart] == ' ' {
				paramStart++
			}

			rules = append(rules, validationRule{
				tag:   tag[start:ruleEnd],
				param: tag[paramStart:end],
			})
		} else {
			// 无参数
			rules = append(rules, validationRule{
				tag:   tag[start:end],
				param: "",
			})
		}

		i++
	}

	return rules
}

// ========== 基准测试 ==========

func BenchmarkParseTag_Baseline_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagBaseline(tag)
	}
}

func BenchmarkParseTag_Opt1_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt1(tag)
	}
}

func BenchmarkParseTag_Opt2_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt2(tag)
	}
}

func BenchmarkParseTag_Opt3_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt3(tag)
	}
}

func BenchmarkParseTag_Opt4_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt4(tag)
	}
}

func BenchmarkParseTag_Opt5_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt5(tag)
	}
}

func BenchmarkParseTag_Opt6_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt6(tag)
	}
}

func BenchmarkParseTag_Opt7_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt7(tag)
	}
}

func BenchmarkParseTag_Opt8_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt8(tag)
	}
}

func BenchmarkParseTag_Opt9_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt9(tag)
	}
}

func BenchmarkParseTag_Opt10_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt10(tag)
	}
}

// ========== 复杂场景 ==========

func BenchmarkParseTag_Baseline_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagBaseline(tag)
	}
}

func BenchmarkParseTag_Opt3_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt3(tag)
	}
}

func BenchmarkParseTag_Opt4_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt4(tag)
	}
}

func BenchmarkParseTag_Opt10_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt10(tag)
	}
}

// ========== 内存分配测试 ==========

func BenchmarkParseTag_Baseline_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagBaseline(tag)
	}
}

func BenchmarkParseTag_Opt3_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt3(tag)
	}
}

func BenchmarkParseTag_Opt4_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt4(tag)
	}
}

func BenchmarkParseTag_Opt10_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt10(tag)
	}
}
