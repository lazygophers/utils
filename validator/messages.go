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

// formatMessage 格式化错误消息，支持模板替换
func formatMessage(template, field, tag, param string) string {
	msg := template
	msg = stringsReplaceAll(msg, "{field}", field)
	msg = stringsReplaceAll(msg, "{tag}", tag)
	msg = stringsReplaceAll(msg, "{param}", param)
	return msg
}

// stringsReplaceAll 替换字符串（避免使用 strings 包导入问题）
func stringsReplaceAll(s, old, new string) string {
	result := ""
	start := 0
	for {
		idx := indexOf(s, old, start)
		if idx == -1 {
			result += s[start:]
			break
		}
		result += s[start:idx] + new
		start = idx + len(old)
	}
	return result
}

// indexOf 查找子字符串位置
func indexOf(s, substr string, start int) int {
	if len(substr) == 0 {
		return start
	}
	for i := start; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

// getDefaultMessage 获取默认错误消息
func getDefaultMessage(tag string) string {
	if msg, ok := defaultMessages[tag]; ok {
		return msg
	}
	return "验证失败: " + tag
}
