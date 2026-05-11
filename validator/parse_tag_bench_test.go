package validator

import (
	"strings"
	"testing"
)

// ========== 测试数据生成 ==========

// 生成固定种子的测试标签（保证可重复性）
func genTags(tagType string, count int) string {
	switch tagType {
	case "simple":
		// 简单标签：required,email,max=100
		return "required,email,max=100"
	case "medium":
		// 中等复杂度：required,email,min=18,max=100,len=6-20
		return "required,email,min=18,max=100,len=6-20"
	case "complex":
		// 复杂标签：多个规则
		return "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$,url,mobile,alpha,alphanum,alpha_dash"
	case "whitespace":
		// 带空格：required , email , max = 100
		return "required , email , max = 100"
	case "realistic":
		// 真实场景：用户验证规则
		return "required,email,min=3,max=50,len=6-20,regex=^[a-zA-Z0-9]+$"
	case "many":
		// 大量规则：20个规则
		rules := make([]string, count)
		for i := 0; i < count; i++ {
			if i%3 == 0 {
				rules[i] = "rule" + string(rune('A'+i))
			} else if i%3 == 1 {
				rules[i] = "rule" + string(rune('A'+i)) + "=param"
			} else {
				rules[i] = "rule" + string(rune('A'+i)) + " = value "
			}
		}
		return strings.Join(rules, ",")
	default:
		return "required"
	}
}

// ========== 当前实现（Baseline） ==========

func parseTag_Current(tag string) []validationRule {
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

// ========== 优化方案1：预分配切片容量 ==========

func parseTag_Prealloc(tag string) []validationRule {
	// 估算规则数：按逗号分隔，假设平均每个规则1个字符
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

// ========== 优化方案2：使用 IndexByte 代替 Index ==========

func parseTag_IndexByte(tag string) []validationRule {
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

func parseTag_Prealloc_IndexByte(tag string) []validationRule {
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

// ========== 优化方案4：减少 TrimSpace 调用（只在需要时trim） ==========

func parseTag_ReduceTrim(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		// 只在开始时trim一次
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			// 提取子串时不再trim，假设用户格式正确
			ruleName := part[:idx]
			param := part[idx+1:]
			// 手动trim前导空格
			for len(ruleName) > 0 && ruleName[0] == ' ' {
				ruleName = ruleName[1:]
			}
			for len(ruleName) > 0 && ruleName[len(ruleName)-1] == ' ' {
				ruleName = ruleName[:len(ruleName)-1]
			}
			for len(param) > 0 && param[0] == ' ' {
				param = param[1:]
			}
			for len(param) > 0 && param[len(param)-1] == ' ' {
				param = param[:len(param)-1]
			}
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案5：手动字节级解析（零分配） ==========

func parseTag_Bytes(tag string) []validationRule {
	// 预估规则数
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
			// 跳过前导空格
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == '=' || ch == ',' || i == len(tag) {
				ruleStart = start
				ruleEnd = i

				// trim尾部空格
				for ruleEnd > ruleStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				if ch == '=' {
					inTag = false
					start = i + 1
				} else {
					// 完成规则解析
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
			// 跳过前导空格
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == ',' || i == len(tag) {
				paramStart = start
				paramEnd = i

				// trim尾部空格
				for paramEnd > paramStart && tag[paramEnd-1] == ' ' {
					paramEnd--
				}

				// 完成规则解析
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

// ========== 优化方案6：strings.Builder 重用 ==========

func parseTag_Builder(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	var builder strings.Builder
	builder.Grow(50)

	start := 0
	inTag := true

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
				end := i
				for end > start && tag[end-1] == ' ' {
					end--
				}

				if ch == '=' {
					builder.Reset()
					builder.WriteString(tag[start:end])
					inTag = false
					start = i + 1
				} else {
					if start < end {
						rules = append(rules, validationRule{
							tag:   tag[start:end],
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
				end := i
				for end > start && tag[end-1] == ' ' {
					end--
				}

				ruleName := builder.String()
				rules = append(rules, validationRule{
					tag:   ruleName,
					param: tag[start:end],
				})

				inTag = true
				start = i + 1
			}
		}
	}

	return rules
}

// ========== 优化方案7：单次遍历 + 字符串切片重用 ==========

func parseTag_SinglePass(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	var currentRule strings.Builder
	var currentParam strings.Builder
	currentRule.Grow(20)
	currentParam.Grow(20)

	inParam := false
	trimmed := false

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		// 跳过前导空格
		if !trimmed && ch == ' ' {
			continue
		}
		trimmed = true

		if ch == '=' {
			inParam = true
			trimmed = false
			continue
		}

		if ch == ',' || i == len(tag) {
			if inParam {
				// trim尾部空格
				paramStr := currentParam.String()
				end := len(paramStr)
				for end > 0 && paramStr[end-1] == ' ' {
					end--
				}
				rules = append(rules, validationRule{
					tag:   currentRule.String(),
					param: paramStr[:end],
				})
				currentParam.Reset()
			} else {
				ruleStr := currentRule.String()
				rules = append(rules, validationRule{
					tag:   ruleStr,
					param: "",
				})
			}
			currentRule.Reset()
			inParam = false
			trimmed = false
			continue
		}

		if inParam {
			currentParam.WriteByte(ch)
		} else {
			currentRule.WriteByte(ch)
		}
	}

	return rules
}

// ========== 优化方案8：使用 strings.Split 后批量处理 ==========

func parseTag_Batch(tag string) []validationRule {
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

// ========== 优化方案9：手动解析 + 索引优化 ==========

func parseTag_ManualIndex(tag string) []validationRule {
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

		// 跟踪等号位置
		if ch == '=' {
			eqIndex = i
			continue
		}

		if ch == ',' || i == len(tag) {
			// 处理前导空格
			for trimStart < i && tag[trimStart] == ' ' {
				trimStart++
			}
			if trimStart >= i {
				start = i + 1
				trimStart = start
				eqIndex = -1
				continue
			}

			// 处理尾部空格
			end := i
			for end > trimStart && tag[end-1] == ' ' {
				end--
			}

			if eqIndex != -1 && eqIndex < end {
				// 有参数
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
				// 无参数
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

// ========== 优化方案10：混合方案（Split + IndexByte + 预分配） ==========

func parseTag_Hybrid(tag string) []validationRule {
	// 使用 Split 简化逻辑，但优化其他部分
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		// 快速跳过空字符串
		if part == "" {
			continue
		}

		// 跳过前导空格
		start := 0
		for start < len(part) && part[start] == ' ' {
			start++
		}
		if start >= len(part) {
			continue
		}

		// 跳过尾部空格
		end := len(part)
		for end > start && part[end-1] == ' ' {
			end--
		}

		trimmed := part[start:end]

		// 使用 IndexByte
		if idx := strings.IndexByte(trimmed, '='); idx != -1 {
			ruleName := trimmed[:idx]
			param := trimmed[idx+1:]

			// 只在必要时trim参数
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

// ========== 基准测试 ==========

func BenchmarkParseTag_Current_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Prealloc_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc(tag)
	}
}

func BenchmarkParseTag_IndexByte_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_IndexByte(tag)
	}
}

func BenchmarkParseTag_Prealloc_IndexByte_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc_IndexByte(tag)
	}
}

func BenchmarkParseTag_ReduceTrim_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ReduceTrim(tag)
	}
}

func BenchmarkParseTag_Bytes_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Bytes(tag)
	}
}

func BenchmarkParseTag_Builder_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Builder(tag)
	}
}

func BenchmarkParseTag_SinglePass_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_SinglePass(tag)
	}
}

func BenchmarkParseTag_Batch_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Batch(tag)
	}
}

func BenchmarkParseTag_ManualIndex_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ManualIndex(tag)
	}
}

func BenchmarkParseTag_Hybrid_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Hybrid(tag)
	}
}

// ========== 复杂场景基准 ==========

func BenchmarkParseTag_Current_Medium(b *testing.B) {
	tag := genTags("medium", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Complex(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Whitespace(b *testing.B) {
	tag := genTags("whitespace", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Realistic(b *testing.B) {
	tag := genTags("realistic", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Many(b *testing.B) {
	tag := genTags("many", 20)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

// ========== 内存分配测试 ==========

func BenchmarkParseTag_Current_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Prealloc_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc(tag)
	}
}

func BenchmarkParseTag_IndexByte_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_IndexByte(tag)
	}
}

func BenchmarkParseTag_Prealloc_IndexByte_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc_IndexByte(tag)
	}
}

func BenchmarkParseTag_ReduceTrim_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ReduceTrim(tag)
	}
}

func BenchmarkParseTag_Bytes_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Bytes(tag)
	}
}

func BenchmarkParseTag_Builder_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Builder(tag)
	}
}

func BenchmarkParseTag_SinglePass_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_SinglePass(tag)
	}
}

func BenchmarkParseTag_Batch_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Batch(tag)
	}
}

func BenchmarkParseTag_ManualIndex_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ManualIndex(tag)
	}
}

func BenchmarkParseTag_Hybrid_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Hybrid(tag)
	}
}
