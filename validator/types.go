package validator

import (
	"fmt"
	"strings"
)

// FieldError 表示单个字段的验证错误
type FieldError struct {
	Field       string      // 字段名称（优先 JSON tag）
	Tag         string      // 验证标签
	Value       interface{} // 字段值
	Param       string      // 验证参数
	ActualTag   string      // 实际验证标签
	Namespace   string      // 完整命名空间
	StructField string      // 结构体字段路径
	Message     string      // 错误消息
}

// Error 实现 error 接口
func (e *FieldError) Error() string {
	return e.Message
}

// ValidationErrors 验证错误集合
type ValidationErrors []*FieldError

// Error 实现 error 接口
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	if len(e) == 1 {
		return e[0].Error()
	}

	var messages []string
	for _, err := range e {
		messages = append(messages, err.Error())
	}

	return strings.Join(messages, "; ")
}

// First 返回第一个错误
func (e ValidationErrors) First() *FieldError {
	if len(e) == 0 {
		return nil
	}
	return e[0]
}

// FirstError 返回第一个错误的消息
func (e ValidationErrors) FirstError() string {
	if first := e.First(); first != nil {
		return first.Error()
	}
	return ""
}

// ByField 根据字段名获取错误
func (e ValidationErrors) ByField(field string) *FieldError {
	for _, err := range e {
		if err.Field == field {
			return err
		}
	}
	return nil
}

// HasField 检查是否包含指定字段的错误
func (e ValidationErrors) HasField(field string) bool {
	return e.ByField(field) != nil
}

// Fields 返回所有出错的字段名
func (e ValidationErrors) Fields() []string {
	var fields []string
	for _, err := range e {
		fields = append(fields, err.Field)
	}
	return fields
}

// Messages 返回所有错误消息
func (e ValidationErrors) Messages() []string {
	var messages []string
	for _, err := range e {
		messages = append(messages, err.Message)
	}
	return messages
}

// ToMap 转换为字段名到错误消息的映射
func (e ValidationErrors) ToMap() map[string]string {
	result := make(map[string]string)
	for _, err := range e {
		result[err.Field] = err.Message
	}
	return result
}

// ToDetailMap 转换为字段名到详细错误信息的映射
func (e ValidationErrors) ToDetailMap() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})
	for _, err := range e {
		result[err.Field] = map[string]interface{}{
			"tag":     err.Tag,
			"value":   err.Value,
			"param":   err.Param,
			"message": err.Message,
		}
	}
	return result
}

// String 返回格式化的错误字符串
func (e ValidationErrors) String() string {
	return e.Error()
}

// JSON 返回 JSON 格式的错误信息
func (e ValidationErrors) JSON() map[string]interface{} {
	errors := make([]map[string]interface{}, 0, len(e))
	for _, err := range e {
		errors = append(errors, map[string]interface{}{
			"field":   err.Field,
			"tag":     err.Tag,
			"value":   err.Value,
			"param":   err.Param,
			"message": err.Message,
		})
	}

	return map[string]interface{}{
		"errors": errors,
		"count":  len(e),
	}
}

// Len 返回错误数量
func (e ValidationErrors) Len() int {
	return len(e)
}

// IsEmpty 检查是否为空
func (e ValidationErrors) IsEmpty() bool {
	return len(e) == 0
}

// Add 添加错误
func (e *ValidationErrors) Add(err *FieldError) {
	*e = append(*e, err)
}

// Merge 合并其他验证错误
func (e *ValidationErrors) Merge(other ValidationErrors) {
	*e = append(*e, other...)
}

// Filter 根据条件过滤错误
func (e ValidationErrors) Filter(fn func(*FieldError) bool) ValidationErrors {
	var result ValidationErrors
	for _, err := range e {
		if fn(err) {
			result = append(result, err)
		}
	}
	return result
}

// ForField 获取指定字段的所有错误
func (e ValidationErrors) ForField(field string) ValidationErrors {
	return e.Filter(func(err *FieldError) bool {
		return err.Field == field
	})
}

// ForTag 获取指定标签的所有错误
func (e ValidationErrors) ForTag(tag string) ValidationErrors {
	return e.Filter(func(err *FieldError) bool {
		return err.Tag == tag
	})
}

// Format 格式化错误信息
func (e ValidationErrors) Format(format string) string {
	var formatted []string
	for _, err := range e {
		msg := format
		msg = strings.ReplaceAll(msg, "{field}", err.Field)
		msg = strings.ReplaceAll(msg, "{tag}", err.Tag)
		msg = strings.ReplaceAll(msg, "{message}", err.Message)
		msg = strings.ReplaceAll(msg, "{value}", fmt.Sprintf("%v", err.Value))
		msg = strings.ReplaceAll(msg, "{param}", err.Param)
		formatted = append(formatted, msg)
	}
	return strings.Join(formatted, "; ")
}
