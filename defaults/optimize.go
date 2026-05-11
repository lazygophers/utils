package defaults

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// 缓存结构
type cachedFieldInfo struct {
	index       int
	kind        reflect.Kind
	defaultTag  string
	stringValue string
	intValue    int64
	uintValue   uint64
	floatValue  float64
	boolValue   bool
	isTime      bool
}

type cachedTypeInfo struct {
	fields []cachedFieldInfo
}

var (
	typeCache      sync.Map
	typeCacheMutex sync.RWMutex
)

// setDefaultOptimized 优化版本 - 使用反射信息缓存和预解析值
func setDefaultOptimized(vv reflect.Value, defaultStr string, opts *Options) error {
	if !vv.IsValid() {
		return handleError("invalid reflect value", opts.ErrorMode)
	}

	// 处理结构体的优化路径
	if vv.Kind() == reflect.Struct && vv.CanAddr() {
		typ := vv.Type()

		// 处理 time.Time 特殊类型
		if typ == reflect.TypeOf(time.Time{}) {
			return setTimeDefault(vv, defaultStr, opts)
		}

		// 从缓存获取类型信息
		if cached, ok := typeCache.Load(typ); ok {
			ct := cached.(cachedTypeInfo)
			return setFieldsFromCache(vv, ct, opts)
		}

		// 首次访问，构建缓存
		ct := buildTypeCache(vv, typ)
		typeCache.Store(typ, ct)

		// 首次设置（使用缓存）
		return setFieldsFromCache(vv, ct, opts)
	}

	// 其他类型使用原始逻辑
	return setDefaultWithOptions(vv, defaultStr, opts)
}

// buildTypeCache 构建类型缓存
func buildTypeCache(vv reflect.Value, typ reflect.Type) cachedTypeInfo {
	fields := make([]cachedFieldInfo, 0, vv.NumField())

	for i := 0; i < vv.NumField(); i++ {
		field := vv.Field(i)
		fieldType := typ.Field(i)

		if !field.CanSet() {
			continue
		}

		defaultTag := fieldType.Tag.Get("default")
		if defaultTag == "" {
			continue
		}

		info := cachedFieldInfo{
			index:      i,
			kind:       field.Kind(),
			defaultTag: defaultTag,
		}

		// 预解析值
		switch field.Kind() {
		case reflect.String:
			info.stringValue = defaultTag
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Sscanf(defaultTag, "%d", &info.intValue)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Sscanf(defaultTag, "%d", &info.uintValue)
		case reflect.Float32, reflect.Float64:
			fmt.Sscanf(defaultTag, "%f", &info.floatValue)
		case reflect.Bool:
			info.boolValue = defaultTag == "true"
		case reflect.Struct:
			info.isTime = fieldType.Type == reflect.TypeOf(time.Time{})
		}

		fields = append(fields, info)
	}

	return cachedTypeInfo{fields: fields}
}

// setFieldsFromCache 从缓存设置字段值
func setFieldsFromCache(vv reflect.Value, ct cachedTypeInfo, opts *Options) error {
	for _, info := range ct.fields {
		field := vv.Field(info.index)

		if !field.CanSet() {
			continue
		}

		// 处理条件默认值
		if isConditionalDefault(info.defaultTag) {
			handleConditionalDefault(field, info.defaultTag, vv, opts)
			continue
		}

		// 使用预解析的值快速设置
		switch info.kind {
		case reflect.String:
			if shouldSetValue(field.String() == "", info.defaultTag, opts.AllowOverwrite) {
				field.SetString(info.stringValue)
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if shouldSetValue(field.Int() == 0, info.defaultTag, opts.AllowOverwrite) {
				field.SetInt(info.intValue)
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if shouldSetValue(field.Uint() == 0, info.defaultTag, opts.AllowOverwrite) {
				field.SetUint(info.uintValue)
			}

		case reflect.Float32, reflect.Float64:
			if shouldSetValue(field.Float() == 0, info.defaultTag, opts.AllowOverwrite) {
				field.SetFloat(info.floatValue)
			}

		case reflect.Bool:
			if shouldSetValue(!field.Bool(), info.defaultTag, opts.AllowOverwrite) {
				field.SetBool(info.boolValue)
			}

		case reflect.Ptr:
			if field.IsNil() && info.defaultTag != "" {
				_ = setPtrDefault(field, info.defaultTag, opts)
			}

		case reflect.Struct:
			if info.isTime {
				_ = setTimeDefault(field, info.defaultTag, opts)
			} else {
				// 嵌套结构体
				_ = setDefaultOptimized(field, info.defaultTag, opts)
			}

		case reflect.Interface:
			if field.IsNil() && info.defaultTag != "" {
				_ = setInterfaceDefault(field, info.defaultTag, opts)
			}

		case reflect.Slice:
			if field.IsNil() && info.defaultTag != "" {
				_ = setSliceDefault(field, info.defaultTag, opts)
			}

		case reflect.Map:
			if field.IsNil() && info.defaultTag != "" {
				_ = setMapDefault(field, info.defaultTag, opts)
			}

		case reflect.Chan:
			if field.IsNil() && info.defaultTag != "" {
				_ = setChanDefault(field, info.defaultTag, opts)
			}

		case reflect.Func:
			if field.IsNil() && info.defaultTag != "" {
				_ = setFuncDefault(field, info.defaultTag, opts)
			}
		}
	}

	return nil
}

// handleConditionalDefault 处理条件默认值
func handleConditionalDefault(field reflect.Value, defaultTag string, vv reflect.Value, opts *Options) {
	conditionalValues := parseConditionalDefault(defaultTag)
	selectedValue := ""

	for condition, value := range conditionalValues {
		if matchCondition(condition, vv) {
			selectedValue = value
			break
		}
	}

	if selectedValue != "" && isZero(field) {
		_ = setValueFromString(field, selectedValue)
	}
}

// SetDefaultsOptimized 使用优化算法设置默认值
func SetDefaultsOptimized(value interface{}) {
	err := setDefaultOptimized(reflect.ValueOf(value), "", defaultOptions)
	if err != nil {
		panic(err)
	}
}

// SetDefaultsWithOptionsOptimized 使用优化算法和自定义选项设置默认值
func SetDefaultsWithOptionsOptimized(value interface{}, opts *Options) error {
	if opts == nil {
		opts = defaultOptions
	}

	return setDefaultOptimized(reflect.ValueOf(value), "", opts)
}

// 性能统计
type PerfStats struct {
	HitCount  int64
	MissCount int64
}

var perfStats PerfStats

// GetPerfStats 获取性能统计
func GetPerfStats() PerfStats {
	return perfStats
}

// ClearCache 清除缓存（用于测试）
func ClearCache() {
	typeCache = sync.Map{}
}
