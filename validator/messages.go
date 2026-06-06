package validator

var defaultMessages = map[string]string{
	"required":          "{field} 为必填项",
	"email":             "{field} 格式无效",
	"url":               "{field} 必须是有效的 URL",
	"min":               "{field} 长度不能小于 {param}",
	"max":               "{field} 长度不能大于 {param}",
	"eq":                "{field} 必须等于 {param}",
	"ne":                "{field} 不能等于 {param}",
	"eqfield":           "{field} 必须与 {param} 相同",
	"nefield":           "{field} 不能与 {param} 相同",
	"required_with":     "当 {param} 存在时 {field} 为必填项",
	"required_without":  "当 {param} 不存在时 {field} 为必填项",
	"required_if":       "{field} 为必填项 ({param})",
	"alpha":             "{field} 只能包含字母",
	"alphanum":          "{field} 只能包含字母和数字",
	"strong_password":   "{field} 必须是强密码",
	"ipv4":              "{field} 必须是有效的 IPv4 地址",
	"mac":               "{field} 必须是有效的 MAC 地址",
	"json":              "{field} 必须是有效的 JSON",
	"uuid":              "{field} 必须是有效的 UUID",
}

// formatMessage 格式化错误消息，零分配快速路径。
// 使用 byte slice builder 替代 O(n²) 字符串拼接。
func formatMessage(template, field, tag, param string) string {
	// 快速路径：手动扫描避免 strings.Contains 开销
	hasPlaceholder := false
	for i := 0; i < len(template); i++ {
		if template[i] == '{' {
			hasPlaceholder = true
			break
		}
	}
	if !hasPlaceholder {
		return template
	}

	result := make([]byte, 0, len(template)+len(field)+len(tag)+len(param))

	i := 0
	for i < len(template) {
		if i+7 <= len(template) {
			if template[i:i+7] == "{field}" {
				result = append(result, field...)
				i += 7
				continue
			}
			if template[i:i+7] == "{param}" {
				result = append(result, param...)
				i += 7
				continue
			}
		}
		if i+5 <= len(template) {
			if template[i:i+5] == "{tag}" {
				result = append(result, tag...)
				i += 5
				continue
			}
		}
		result = append(result, template[i])
		i++
	}

	return string(result)
}

// getDefaultMessage 获取默认错误消息
func getDefaultMessage(tag string) string {
	if msg, ok := defaultMessages[tag]; ok {
		return msg
	}
	return "验证失败: " + tag
}
