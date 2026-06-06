package validator

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

// formatMessage 原始实现（用于对比）
func formatMessageOriginal(template string, err *FieldError) string {
	result := template
	result = strings.ReplaceAll(result, "{field}", err.Field)
	result = strings.ReplaceAll(result, "{tag}", err.Tag)
	result = strings.ReplaceAll(result, "{param}", err.Param)
	if err.Value != nil {
		result = strings.ReplaceAll(result, "{value}", fmt.Sprintf("%v", err.Value))
	} else {
		result = strings.ReplaceAll(result, "{value}", "")
	}
	return result
}

// formatMessageBuilder 方案1: strings.Builder (减少内存分配)
func formatMessageBuilder(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}
	var builder strings.Builder
	builder.Grow(len(template) + 50)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) && template[i:i+7] == "{field}" {
			builder.WriteString(err.Field)
			i += 7
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{param}" {
			builder.WriteString(err.Param)
			i += 7
			continue
		}
		if i+5 <= len(template) && template[i:i+5] == "{tag}" {
			builder.WriteString(err.Tag)
			i += 5
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{value}" {
			if err.Value != nil {
				builder.WriteString(fmt.Sprintf("%v", err.Value))
			}
			i += 7
			continue
		}
		builder.WriteByte(template[i])
		i++
	}
	return builder.String()
}

// 编译模板相关
type compiledTemplate struct {
	parts []templatePart
}

type templatePart struct {
	isPlaceholder bool
	value         string
}

var templateCache sync.Map

// formatMessageCompiled 方案2: 预编译模板（缓存占位符位置）
func formatMessageCompiled(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	compiled, _ := templateCache.LoadOrStore(template, compileTemplate(template))
	ct := compiled.(*compiledTemplate)

	var builder strings.Builder
	for _, part := range ct.parts {
		if part.isPlaceholder {
			switch part.value {
			case "field":
				builder.WriteString(err.Field)
			case "tag":
				builder.WriteString(err.Tag)
			case "param":
				builder.WriteString(err.Param)
			case "value":
				if err.Value != nil {
					builder.WriteString(fmt.Sprintf("%v", err.Value))
				}
			}
		} else {
			builder.WriteString(part.value)
		}
	}
	return builder.String()
}

func compileTemplate(template string) *compiledTemplate {
	var parts []templatePart
	i := 0
	for i < len(template) {
		if i+7 <= len(template) && template[i:i+7] == "{field}" {
			parts = append(parts, templatePart{isPlaceholder: true, value: "field"})
			i += 7
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{param}" {
			parts = append(parts, templatePart{isPlaceholder: true, value: "param"})
			i += 7
			continue
		}
		if i+5 <= len(template) && template[i:i+5] == "{tag}" {
			parts = append(parts, templatePart{isPlaceholder: true, value: "tag"})
			i += 5
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{value}" {
			parts = append(parts, templatePart{isPlaceholder: true, value: "value"})
			i += 7
			continue
		}
		start := i
		for i < len(template) && template[i] != '{' {
			i++
		}
		parts = append(parts, templatePart{isPlaceholder: false, value: template[start:i]})
	}
	return &compiledTemplate{parts: parts}
}

// formatMessageByteSlice 方案3: 字节数组预分配（当前实现的改进版）
func formatMessageByteSlice(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	estimatedSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	result := make([]byte, 0, estimatedSize)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) {
			switch template[i:i+7] {
			case "{field}":
				result = append(result, err.Field...)
				i += 7
				continue
			case "{param}":
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}
		if i+5 <= len(template) && template[i:i+5] == "{tag}" {
			result = append(result, err.Tag...)
			i += 5
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{value}" {
			if err.Value != nil {
				result = append(result, fmt.Sprintf("%v", err.Value)...)
			}
			i += 7
			continue
		}
		result = append(result, template[i])
		i++
	}
	return string(result)
}

// 快速值格式化
func formatValueFast(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return fmt.Sprintf("%d", val)
	case int64:
		return fmt.Sprintf("%d", val)
	case float64:
		return fmt.Sprintf("%f", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatMessageNoFmt 方案4: 消除 fmt.Sprintf 调用（value 快速格式化）
func formatMessageNoFmt(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	estimatedSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	result := make([]byte, 0, estimatedSize)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) {
			switch template[i:i+7] {
			case "{field}":
				result = append(result, err.Field...)
				i += 7
				continue
			case "{param}":
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}
		if i+5 <= len(template) && template[i:i+5] == "{tag}" {
			result = append(result, err.Tag...)
			i += 5
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{value}" {
			if err.Value != nil {
				valStr := formatValueFast(err.Value)
				result = append(result, valStr...)
			}
			i += 7
			continue
		}
		result = append(result, template[i])
		i++
	}
	return string(result)
}

// formatMessageSingleReplace 方案5: 单次 strings.ReplaceAll 链式调用
func formatMessageSingleReplace(template string, err *FieldError) string {
	result := template
	if strings.Contains(result, "{value}") {
		valueStr := ""
		if err.Value != nil {
			valueStr = fmt.Sprintf("%v", err.Value)
		}
		result = strings.ReplaceAll(result, "{value}", valueStr)
	}
	if strings.Contains(result, "{field}") {
		result = strings.ReplaceAll(result, "{field}", err.Field)
	}
	if strings.Contains(result, "{tag}") {
		result = strings.ReplaceAll(result, "{tag}", err.Tag)
	}
	if strings.Contains(result, "{param}") {
		result = strings.ReplaceAll(result, "{param}", err.Param)
	}
	return result
}

// formatMessageHashtable 方案7: 模板哈希表查找（O(1)占位符匹配）
func formatMessageHashtable(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	placeholders := map[string]string{
		"{field}": err.Field,
		"{tag}":   err.Tag,
		"{param}": err.Param,
	}

	if err.Value != nil {
		placeholders["{value}"] = fmt.Sprintf("%v", err.Value)
	}

	result := template
	for placeholder, value := range placeholders {
		if strings.Contains(result, placeholder) {
			result = strings.ReplaceAll(result, placeholder, value)
		}
	}
	return result
}

// formatMessageInlineCheck 方案8: 手动内联 strings.Contains 检查
func formatMessageInlineCheck(template string, err *FieldError) string {
	hasOpenBrace := false
	for _, c := range template {
		if c == '{' {
			hasOpenBrace = true
			break
		}
	}
	if !hasOpenBrace {
		return template
	}

	estimatedSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	result := make([]byte, 0, estimatedSize)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) {
			switch template[i:i+7] {
			case "{field}":
				result = append(result, err.Field...)
				i += 7
				continue
			case "{param}":
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}
		if i+5 <= len(template) && template[i:i+5] == "{tag}" {
			result = append(result, err.Tag...)
			i += 5
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{value}" {
			if err.Value != nil {
				result = append(result, fmt.Sprintf("%v", err.Value)...)
			}
			i += 7
			continue
		}
		result = append(result, template[i])
		i++
	}
	return string(result)
}

// formatMessagePrecompute 方案9: 预计算所有替换值
func formatMessagePrecompute(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	valueStr := ""
	if err.Value != nil {
		valueStr = fmt.Sprintf("%v", err.Value)
	}

	replacements := [][]string{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
		{"{value}", valueStr},
	}

	result := template
	for _, repl := range replacements {
		if strings.Contains(result, repl[0]) {
			result = strings.ReplaceAll(result, repl[0], repl[1])
		}
	}
	return result
}

// formatMessageBytesBuffer 方案10: 使用 bytes.Buffer (替代 strings.Builder)
func formatMessageBytesBuffer(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	var buf bytes.Buffer
	buf.Grow(len(template) + 50)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) && template[i:i+7] == "{field}" {
			buf.WriteString(err.Field)
			i += 7
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{param}" {
			buf.WriteString(err.Param)
			i += 7
			continue
		}
		if i+5 <= len(template) && template[i:i+5] == "{tag}" {
			buf.WriteString(err.Tag)
			i += 5
			continue
		}
		if i+7 <= len(template) && template[i:i+7] == "{value}" {
			if err.Value != nil {
				buf.WriteString(fmt.Sprintf("%v", err.Value))
			}
			i += 7
			continue
		}
		buf.WriteByte(template[i])
		i++
	}
	return buf.String()
}

// formatMessageOptimizedCurrent 方案11: 当前实现的改进版（减少边界检查）
func formatMessageOptimizedCurrent(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	estimatedSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	result := make([]byte, 0, estimatedSize)

	templateLen := len(template)
	i := 0

	for i < templateLen {
		// 检查 7 字节占位符
		if i+7 <= templateLen {
			switch template[i:i+7] {
			case "{field}":
				result = append(result, err.Field...)
				i += 7
				continue
			case "{param}":
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}

		// 检查 6 字节占位符
		if i+7 <= templateLen && template[i:i+7] == "{value}" {
			if err.Value != nil {
				result = append(result, fmt.Sprintf("%v", err.Value)...)
			}
			i += 7
			continue
		}

		// 检查 5 字节占位符
		if i+5 <= templateLen && template[i:i+5] == "{tag}" {
			result = append(result, err.Tag...)
			i += 5
			continue
		}

		result = append(result, template[i])
		i++
	}

	return string(result)
}

// formatMessageFastPath 方案12: 消除 strings.Contains 调用（内联快速路径）
func formatMessageFastPath(template string, err *FieldError) string {
	// 内联检查是否有 '{'
	hasPlaceholder := false
	for j := 0; j < len(template); j++ {
		if template[j] == '{' {
			hasPlaceholder = true
			break
		}
	}
	if !hasPlaceholder {
		return template
	}

	estimatedSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	result := make([]byte, 0, estimatedSize)

	i := 0
	templateLen := len(template)

	for i < templateLen {
		if i+7 <= templateLen {
			chunk := template[i:i+7]
			if chunk == "{field}" {
				result = append(result, err.Field...)
				i += 7
				continue
			}
			if chunk == "{param}" {
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}

		if i+7 <= templateLen && template[i:i+7] == "{value}" {
			if err.Value != nil {
				result = append(result, fmt.Sprintf("%v", err.Value)...)
			}
			i += 7
			continue
		}

		if i+5 <= templateLen && template[i:i+5] == "{tag}" {
			result = append(result, err.Tag...)
			i += 5
			continue
		}

		result = append(result, template[i])
		i++
	}

	return string(result)
}
