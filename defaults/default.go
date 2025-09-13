package defaults

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// DefaultFunc 自定义默认值函数类型
type DefaultFunc func() interface{}

// Options 配置选项
type Options struct {
	// ErrorMode 错误处理模式
	ErrorMode ErrorMode
	// CustomDefaults 自定义默认值函数映射
	CustomDefaults map[string]DefaultFunc
	// ValidateDefaults 是否验证默认值
	ValidateDefaults bool
	// AllowOverwrite 是否允许覆盖非零值
	AllowOverwrite bool
}

// ErrorMode 错误处理模式
type ErrorMode int

const (
	// ErrorModePanic 遇到错误时 panic（默认）
	ErrorModePanic ErrorMode = iota
	// ErrorModeIgnore 忽略错误，继续处理
	ErrorModeIgnore
	// ErrorModeReturn 返回错误
	ErrorModeReturn
)

var (
	// 默认配置
	defaultOptions = &Options{
		ErrorMode:        ErrorModePanic,
		CustomDefaults:   make(map[string]DefaultFunc),
		ValidateDefaults: false,
		AllowOverwrite:   false,
	}
)

// SetDefaultsWithOptions 使用自定义选项设置默认值
func SetDefaultsWithOptions(value interface{}, opts *Options) error {
	if opts == nil {
		opts = defaultOptions
	}

	return setDefaultWithOptions(reflect.ValueOf(value), "", opts)
}

// SetDefaults 设置默认值（使用默认配置）
func SetDefaults(value interface{}) {
	err := SetDefaultsWithOptions(value, defaultOptions)
	if err != nil {
		panic(err)
	}
}

// setDefaultWithOptions 内部实现函数
func setDefaultWithOptions(vv reflect.Value, defaultStr string, opts *Options) error {
	if !vv.IsValid() {
		return handleError("invalid reflect value", opts.ErrorMode)
	}

	switch vv.Kind() {
	case reflect.String:
		return setStringDefault(vv, defaultStr, opts)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUintDefault(vv, defaultStr, opts)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setIntDefault(vv, defaultStr, opts)

	case reflect.Float32, reflect.Float64:
		return setFloatDefault(vv, defaultStr, opts)

	case reflect.Bool:
		return setBoolDefault(vv, defaultStr, opts)

	case reflect.Ptr:
		return setPtrDefault(vv, defaultStr, opts)

	case reflect.Struct:
		return setStructDefault(vv, defaultStr, opts)

	case reflect.Interface:
		return setInterfaceDefault(vv, defaultStr, opts)

	case reflect.Slice:
		return setSliceDefault(vv, defaultStr, opts)

	case reflect.Array:
		return setArrayDefault(vv, defaultStr, opts)

	case reflect.Map:
		return setMapDefault(vv, defaultStr, opts)

	case reflect.Chan:
		return setChanDefault(vv, defaultStr, opts)

	case reflect.Func:
		return setFuncDefault(vv, defaultStr, opts)

	default:
		return handleError(fmt.Sprintf("unsupported kind: %s", vv.Kind().String()), opts.ErrorMode)
	}
}

// setStringDefault 设置字符串默认值
func setStringDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.String() == "", defaultStr, opts.AllowOverwrite) {
		if customFunc, ok := opts.CustomDefaults["string"]; ok {
			if val := customFunc(); val != nil {
				if strVal, ok := val.(string); ok {
					vv.SetString(strVal)
				}
			}
		} else if defaultStr != "" {
			vv.SetString(defaultStr)
		}
	}
	return nil
}

// setUintDefault 设置无符号整数默认值
func setUintDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.Uint() == 0, defaultStr, opts.AllowOverwrite) {
		if customFunc, ok := opts.CustomDefaults["uint"]; ok {
			if val := customFunc(); val != nil {
				if uintVal, ok := val.(uint64); ok {
					vv.SetUint(uintVal)
				}
			}
		} else if defaultStr != "" {
			val, err := strconv.ParseUint(defaultStr, 10, 64)
			if err != nil {
				return handleError(fmt.Sprintf("invalid default value for uint field: %s", defaultStr), opts.ErrorMode)
			}
			vv.SetUint(val)
		}
	}
	return nil
}

// setIntDefault 设置整数默认值
func setIntDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.Int() == 0, defaultStr, opts.AllowOverwrite) {
		if customFunc, ok := opts.CustomDefaults["int"]; ok {
			if val := customFunc(); val != nil {
				if intVal, ok := val.(int64); ok {
					vv.SetInt(intVal)
				}
			}
		} else if defaultStr != "" {
			val, err := strconv.ParseInt(defaultStr, 10, 64)
			if err != nil {
				return handleError(fmt.Sprintf("invalid default value for int field: %s", defaultStr), opts.ErrorMode)
			}
			vv.SetInt(val)
		}
	}
	return nil
}

// setFloatDefault 设置浮点数默认值
func setFloatDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.Float() == 0, defaultStr, opts.AllowOverwrite) {
		if customFunc, ok := opts.CustomDefaults["float"]; ok {
			if val := customFunc(); val != nil {
				if floatVal, ok := val.(float64); ok {
					vv.SetFloat(floatVal)
				}
			}
		} else if defaultStr != "" {
			val, err := strconv.ParseFloat(defaultStr, 64)
			if err != nil {
				return handleError(fmt.Sprintf("invalid default value for float field: %s", defaultStr), opts.ErrorMode)
			}
			vv.SetFloat(val)
		}
	}
	return nil
}

// setBoolDefault 设置布尔值默认值
func setBoolDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.Bool() == false, defaultStr, opts.AllowOverwrite) {
		if customFunc, ok := opts.CustomDefaults["bool"]; ok {
			if val := customFunc(); val != nil {
				if boolVal, ok := val.(bool); ok {
					vv.SetBool(boolVal)
				}
			}
		} else if defaultStr != "" {
			val, err := strconv.ParseBool(defaultStr)
			if err != nil {
				return handleError(fmt.Sprintf("invalid default value for bool field: %s", defaultStr), opts.ErrorMode)
			}
			vv.SetBool(val)
		}
	}
	return nil
}

// setPtrDefault 设置指针默认值
func setPtrDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() {
		vv.Set(reflect.New(vv.Type().Elem()))
	}

	// 处理多层指针
	for vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
		if vv.Kind() == reflect.Ptr && vv.IsNil() {
			vv.Set(reflect.New(vv.Type().Elem()))
		}
	}

	return setDefaultWithOptions(vv, defaultStr, opts)
}

// setStructDefault 设置结构体默认值
func setStructDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	// 处理特殊类型
	if vv.Type() == reflect.TypeOf(time.Time{}) {
		return setTimeDefault(vv, defaultStr, opts)
	}

	for i := 0; i < vv.NumField(); i++ {
		field := vv.Field(i)
		fieldType := vv.Type().Field(i)

		if !field.CanSet() {
			continue
		}

		defaultTag := fieldType.Tag.Get("default")
		if err := setDefaultWithOptions(field, defaultTag, opts); err != nil {
			if opts.ErrorMode == ErrorModeReturn {
				return fmt.Errorf("failed to set default for field %s: %w", fieldType.Name, err)
			}
		}
	}
	return nil
}

// setTimeDefault 设置时间默认值
func setTimeDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.Interface().(time.Time).IsZero(), defaultStr, opts.AllowOverwrite) {
		if defaultStr == "now" {
			vv.Set(reflect.ValueOf(time.Now()))
		} else if defaultStr != "" {
			// 尝试解析时间字符串
			layouts := []string{
				time.RFC3339,
				time.RFC3339Nano,
				"2006-01-02 15:04:05",
				"2006-01-02",
				"15:04:05",
			}

			var t time.Time
			var err error
			for _, layout := range layouts {
				t, err = time.Parse(layout, defaultStr)
				if err == nil {
					break
				}
			}

			if err != nil {
				return handleError(fmt.Sprintf("invalid time format: %s", defaultStr), opts.ErrorMode)
			}
			vv.Set(reflect.ValueOf(t))
		}
	}
	return nil
}

// setInterfaceDefault 设置接口默认值
func setInterfaceDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	// 接口类型通常不设置默认值，除非有特殊指定
	if defaultStr != "" && vv.IsNil() {
		// 尝试设置简单类型的默认值
		if strings.Contains(defaultStr, "{") || strings.Contains(defaultStr, "[") {
			// JSON 格式
			var result interface{}
			if err := json.Unmarshal([]byte(defaultStr), &result); err == nil {
				vv.Set(reflect.ValueOf(result))
			}
		} else {
			// 简单字符串
			vv.Set(reflect.ValueOf(defaultStr))
		}
	}
	return nil
}

// setSliceDefault 设置切片默认值
func setSliceDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() {
		if defaultStr == "" {
			// 初始化为空切片
			vv.Set(reflect.MakeSlice(vv.Type(), 0, 0))
		} else {
			// 解析默认值
			if err := parseSliceDefault(vv, defaultStr, opts); err != nil {
				return err
			}
		}
	}

	// 为已存在的切片元素设置默认值
	for i := 0; i < vv.Len(); i++ {
		elem := vv.Index(i)
		if elem.CanSet() {
			if err := setDefaultWithOptions(elem, "", opts); err != nil && opts.ErrorMode == ErrorModeReturn {
				return err
			}
		}
	}

	return nil
}

// setArrayDefault 设置数组默认值
func setArrayDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if defaultStr != "" {
		if err := parseArrayDefault(vv, defaultStr, opts); err != nil {
			return err
		}
	}

	// 为数组元素设置默认值
	for i := 0; i < vv.Len(); i++ {
		elem := vv.Index(i)
		if elem.CanSet() {
			if err := setDefaultWithOptions(elem, "", opts); err != nil && opts.ErrorMode == ErrorModeReturn {
				return err
			}
		}
	}

	return nil
}

// setMapDefault 设置映射默认值
func setMapDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() {
		if defaultStr == "" {
			// 初始化为空映射
			vv.Set(reflect.MakeMap(vv.Type()))
		} else {
			// 解析默认值
			if err := parseMapDefault(vv, defaultStr, opts); err != nil {
				return err
			}
		}
	}
	return nil
}

// setChanDefault 设置通道默认值
func setChanDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() && defaultStr != "" {
		// 解析缓冲区大小
		bufSize := 0
		if defaultStr != "0" {
			var err error
			bufSize, err = strconv.Atoi(defaultStr)
			if err != nil {
				return handleError(fmt.Sprintf("invalid channel buffer size: %s", defaultStr), opts.ErrorMode)
			}
		}
		vv.Set(reflect.MakeChan(vv.Type(), bufSize))
	}
	return nil
}

// setFuncDefault 设置函数默认值
func setFuncDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	// 函数类型通常不设置默认值
	// 如果需要，可以通过自定义函数实现
	if customFunc, ok := opts.CustomDefaults["func"]; ok {
		if vv.IsNil() {
			if val := customFunc(); val != nil {
				if reflect.TypeOf(val).AssignableTo(vv.Type()) {
					vv.Set(reflect.ValueOf(val))
				}
			}
		}
	}
	return nil
}

// parseSliceDefault 解析切片默认值
func parseSliceDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	// 尝试 JSON 解析
	if strings.HasPrefix(defaultStr, "[") && strings.HasSuffix(defaultStr, "]") {
		// 创建切片类型的实例用于解析
		slicePtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), slicePtr.Interface()); err == nil {
			vv.Set(slicePtr.Elem())
			return nil
		}
	}

	// 简单值解析，用逗号分隔
	if strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))

		for i, part := range parts {
			elem := slice.Index(i)
			if err := setDefaultWithOptions(elem, strings.TrimSpace(part), opts); err != nil {
				return err
			}
		}
		vv.Set(slice)
		return nil
	}

	return handleError(fmt.Sprintf("unable to parse slice default: %s", defaultStr), opts.ErrorMode)
}

// parseArrayDefault 解析数组默认值
func parseArrayDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	// 类似于切片的解析逻辑
	if strings.HasPrefix(defaultStr, "[") && strings.HasSuffix(defaultStr, "]") {
		arrayPtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), arrayPtr.Interface()); err == nil {
			vv.Set(arrayPtr.Elem())
			return nil
		}
	}

	// 简单值解析
	if strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		maxParts := vv.Len()
		// 只有在数组有容量时才截断，这样零长度数组会保持parts不截断
		// 从而在循环中触发边界检查
		if len(parts) > maxParts && maxParts > 0 {
			parts = parts[:maxParts]
		}

		for i, part := range parts {
			if i >= vv.Len() {
				break // 对于零长度数组，这个分支现在可以被触发
			}
			elem := vv.Index(i)
			if err := setDefaultWithOptions(elem, strings.TrimSpace(part), opts); err != nil {
				return err
			}
		}
		return nil
	}

	return handleError(fmt.Sprintf("unable to parse array default: %s", defaultStr), opts.ErrorMode)
}

// parseMapDefault 解析映射默认值
func parseMapDefault(vv reflect.Value, defaultStr string, opts *Options) error {
	// 尝试 JSON 解析
	if strings.HasPrefix(defaultStr, "{") && strings.HasSuffix(defaultStr, "}") {
		mapPtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), mapPtr.Interface()); err == nil {
			vv.Set(mapPtr.Elem())
			return nil
		}
	}

	return handleError(fmt.Sprintf("unable to parse map default: %s", defaultStr), opts.ErrorMode)
}

// shouldSetValue 判断是否应该设置值
func shouldSetValue(isZeroValue bool, defaultStr string, allowOverwrite bool) bool {
	return (isZeroValue || allowOverwrite) && defaultStr != ""
}

// handleError 处理错误
func handleError(msg string, mode ErrorMode) error {
	switch mode {
	case ErrorModePanic:
		panic(msg)
	case ErrorModeIgnore:
		return nil
	case ErrorModeReturn:
		return fmt.Errorf("%s", msg)
	default:
		panic(msg)
	}
}

// RegisterCustomDefault 注册自定义默认值函数
func RegisterCustomDefault(typeName string, fn DefaultFunc) {
	if defaultOptions.CustomDefaults == nil {
		defaultOptions.CustomDefaults = make(map[string]DefaultFunc)
	}
	defaultOptions.CustomDefaults[typeName] = fn
}

// ClearCustomDefaults 清除所有自定义默认值函数
func ClearCustomDefaults() {
	defaultOptions.CustomDefaults = make(map[string]DefaultFunc)
}
