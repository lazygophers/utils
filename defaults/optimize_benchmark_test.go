package defaults

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func BenchmarkPerfSimple_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s PerfTestSimple
		_ = setDefaultWithOptions(reflect.ValueOf(&s), "", defaultOptions)
	}
}

func BenchmarkPerfSimple_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s PerfTestSimple
		_ = setDefaultOptimized(reflect.ValueOf(&s), "", defaultOptions)
	}
}

func BenchmarkPerfComplex_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c PerfTestComplex
		_ = setDefaultWithOptions(reflect.ValueOf(&c), "", defaultOptions)
	}
}

func BenchmarkPerfComplex_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c PerfTestComplex
		_ = setDefaultOptimized(reflect.ValueOf(&c), "", defaultOptions)
	}
}

// ========== 测试数据结构 ==========

type TimeTestStruct struct {
	TimeField time.Time `default:"2024-01-15T10:30:00Z"`
}

type PtrTestStruct struct {
	PtrField  *string          `default:"hello"`
	NestedPtr *struct{ X int } `default:"{}"`
}

type InterfaceTestStruct struct {
	InterfaceField interface{} `default:"test"`
	JSONField      interface{} `default:"{"key":"value"}"`
}

type ChanTestStruct struct {
	ChanField  chan int    `default:"10"`
	ChanField2 chan string `default:"0"`
}

type FuncTestStruct struct {
	FuncField func() `default:""`
}

// ========== 原始实现（基线） ==========

// setTimeDefault_Original 原始实现
func setTimeDefault_Original(vv reflect.Value, defaultStr string, opts *Options) error {
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

// ========== 优化方案 1：Unix 时间戳快速路径 ==========
func setTimeDefault_Opt1_UnixFastPath(vv reflect.Value, defaultStr string, opts *Options) error {
	if shouldSetValue(vv.Interface().(time.Time).IsZero(), defaultStr, opts.AllowOverwrite) {
		if defaultStr == "now" {
			vv.Set(reflect.ValueOf(time.Now()))
		} else if defaultStr != "" {
			// 快速路径：检查是否为纯数字（Unix 时间戳）
			if len(defaultStr) > 0 && defaultStr[0] >= '0' && defaultStr[0] <= '9' {
				if sec, err := strconv.ParseInt(defaultStr, 10, 64); err == nil {
					vv.Set(reflect.ValueOf(time.Unix(sec, 0)))
					return nil
				}
			}

			// 标准路径：尝试各种时间格式
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

// ========== 优化方案 2：减少反射调用 ==========
func setTimeDefault_Opt2_LessReflection(vv reflect.Value, defaultStr string, opts *Options) error {
	if !shouldSetValue(vv.Interface().(time.Time).IsZero(), defaultStr, opts.AllowOverwrite) {
		return nil
	}

	if defaultStr == "now" {
		vv.Set(reflect.ValueOf(time.Now()))
		return nil
	}

	if defaultStr == "" {
		return nil
	}

	// 直接内联时间解析，避免循环
	var t time.Time
	var err error

	// Unix 时间戳快速路径
	if t, err = parseUnixTimestamp(defaultStr); err == nil {
		vv.Set(reflect.ValueOf(t))
		return nil
	}

	// 尝试各种格式
	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05",
		"2006-01-02",
		"15:04:05",
	}

	for _, layout := range layouts {
		if t, err = time.Parse(layout, defaultStr); err == nil {
			vv.Set(reflect.ValueOf(t))
			return nil
		}
	}

	return handleError(fmt.Sprintf("invalid time format: %s", defaultStr), opts.ErrorMode)
}

// parseUnixTimestamp 辅助函数
func parseUnixTimestamp(s string) (time.Time, error) {
	if len(s) == 0 || s[0] < '0' || s[0] > '9' {
		return time.Time{}, fmt.Errorf("not a timestamp")
	}
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}

// ========== 优化方案 3：预检查优化 ==========
func setTimeDefault_Opt3_PreCheck(vv reflect.Value, defaultStr string, opts *Options) error {
	// 提前返回优化
	if defaultStr == "" {
		return nil
	}

	currentTime := vv.Interface().(time.Time)
	if !shouldSetValue(currentTime.IsZero(), defaultStr, opts.AllowOverwrite) {
		return nil
	}

	if defaultStr == "now" {
		vv.Set(reflect.ValueOf(time.Now()))
		return nil
	}

	// Unix 快速路径
	if len(defaultStr) > 0 && defaultStr[0] >= '0' && defaultStr[0] <= '9' {
		if sec, err := strconv.ParseInt(defaultStr, 10, 64); err == nil {
			vv.Set(reflect.ValueOf(time.Unix(sec, 0)))
			return nil
		}
	}

	// 标准格式
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
		if t, err = time.Parse(layout, defaultStr); err == nil {
			vv.Set(reflect.ValueOf(t))
			return nil
		}
	}

	return handleError(fmt.Sprintf("invalid time format: %s", defaultStr), opts.ErrorMode)
}

// ========== 优化方案 4：短字符串优化 ==========
func setTimeDefault_Opt4_ShortString(vv reflect.Value, defaultStr string, opts *Options) error {
	if !shouldSetValue(vv.Interface().(time.Time).IsZero(), defaultStr, opts.AllowOverwrite) {
		return nil
	}

	if defaultStr == "now" {
		vv.Set(reflect.ValueOf(time.Now()))
		return nil
	}

	if defaultStr == "" {
		return nil
	}

	// 短字符串快速路径
	if len(defaultStr) <= 10 {
		// 可能是日期或时间
		if t, err := time.Parse("2006-01-02", defaultStr); err == nil {
			vv.Set(reflect.ValueOf(t))
			return nil
		}
		if t, err := time.Parse("15:04:05", defaultStr); err == nil {
			vv.Set(reflect.ValueOf(t))
			return nil
		}
	}

	// Unix 时间戳
	if len(defaultStr) > 0 && defaultStr[0] >= '0' && defaultStr[0] <= '9' {
		if sec, err := strconv.ParseInt(defaultStr, 10, 64); err == nil {
			vv.Set(reflect.ValueOf(time.Unix(sec, 0)))
			return nil
		}
	}

	// 标准格式
	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05",
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		if t, err = time.Parse(layout, defaultStr); err == nil {
			vv.Set(reflect.ValueOf(t))
			return nil
		}
	}

	return handleError(fmt.Sprintf("invalid time format: %s", defaultStr), opts.ErrorMode)
}

// ========== 优化方案 5：缓存 layouts ==========
func setTimeDefault_Opt5_CachedLayouts(vv reflect.Value, defaultStr string, opts *Options) error {
	if !shouldSetValue(vv.Interface().(time.Time).IsZero(), defaultStr, opts.AllowOverwrite) {
		return nil
	}

	if defaultStr == "now" {
		vv.Set(reflect.ValueOf(time.Now()))
		return nil
	}

	if defaultStr == "" {
		return nil
	}

	// Unix 快速路径
	if len(defaultStr) > 0 && defaultStr[0] >= '0' && defaultStr[0] <= '9' {
		if sec, err := strconv.ParseInt(defaultStr, 10, 64); err == nil {
			vv.Set(reflect.ValueOf(time.Unix(sec, 0)))
			return nil
		}
	}

	// 使用缓存的 layouts（从 default.go 导入）
	var t time.Time
	var err error
	for _, layout := range timeLayouts {
		if t, err = time.Parse(layout, defaultStr); err == nil {
			vv.Set(reflect.ValueOf(t))
			return nil
		}
	}

	return handleError(fmt.Sprintf("invalid time format: %s", defaultStr), opts.ErrorMode)
}

// ========== 指针优化方案 ==========

// setPtrDefault_Original 原始实现
func setPtrDefault_Original(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() {
		vv.Set(reflect.New(vv.Type().Elem()))
	}

	// 处理多层指针，确保每一层都被初始化
	current := vv
	for current.Kind() == reflect.Ptr {
		if current.IsNil() {
			current.Set(reflect.New(current.Type().Elem()))
		}
		current = current.Elem()
	}

	return setDefaultWithOptions(vv.Elem(), defaultStr, opts)
}

// 优化方案 6：减少循环次数
func setPtrDefault_Opt6_LessLoop(vv reflect.Value, defaultStr string, opts *Options) error {
	// 快速路径：单层指针
	if vv.IsNil() {
		vv.Set(reflect.New(vv.Type().Elem()))
	}

	elem := vv.Elem()
	if elem.Kind() != reflect.Ptr {
		return setDefaultWithOptions(elem, defaultStr, opts)
	}

	// 多层指针处理
	current := vv
	for current.Kind() == reflect.Ptr {
		if current.IsNil() {
			current.Set(reflect.New(current.Type().Elem()))
		}
		current = current.Elem()
	}

	return setDefaultWithOptions(vv.Elem(), defaultStr, opts)
}

// 优化方案 7：提前返回
func setPtrDefault_Opt7_EarlyReturn(vv reflect.Value, defaultStr string, opts *Options) error {
	if !vv.IsNil() {
		// 已初始化，直接处理元素
		elem := vv.Elem()
		if elem.Kind() != reflect.Ptr {
			return setDefaultWithOptions(elem, defaultStr, opts)
		}
	}

	// 需要初始化
	current := vv
	for current.Kind() == reflect.Ptr {
		if current.IsNil() {
			current.Set(reflect.New(current.Type().Elem()))
		}
		if current.Elem().Kind() != reflect.Ptr {
			return setDefaultWithOptions(current.Elem(), defaultStr, opts)
		}
		current = current.Elem()
	}

	return setDefaultWithOptions(current, defaultStr, opts)
}

// 优化方案 8：批处理反射操作
func setPtrDefault_Opt8_BatchedReflection(vv reflect.Value, defaultStr string, opts *Options) error {
	// 收集所有需要初始化的指针层级
	ptrs := []reflect.Value{vv}
	current := vv

	for current.Kind() == reflect.Ptr {
		if current.IsNil() {
			ptrs = append(ptrs, reflect.New(current.Type().Elem()))
			current = ptrs[len(ptrs)-1]
		} else {
			current = current.Elem()
		}
	}

	// 批量设置
	for i := 0; i < len(ptrs)-1; i++ {
		ptrs[i].Set(ptrs[i+1])
	}

	return setDefaultWithOptions(current, defaultStr, opts)
}

// ========== 接口优化方案 ==========

// setInterfaceDefault_Original 原始实现
func setInterfaceDefault_Original(vv reflect.Value, defaultStr string, opts *Options) error {
	if defaultStr != "" && vv.IsNil() {
		if strings.Contains(defaultStr, "{") || strings.Contains(defaultStr, "[") {
			var result interface{}
			if err := jsonUnmarshal([]byte(defaultStr), &result); err == nil {
				vv.Set(reflect.ValueOf(result))
			}
		} else {
			vv.Set(reflect.ValueOf(defaultStr))
		}
	}
	return nil
}

// 辅助函数（避免 import encoding/json）
func jsonUnmarshal(data []byte, v *interface{}) error {
	return fmt.Errorf("not implemented")
}

// 优化方案 9：快速类型判断
func setInterfaceDefault_Opt9_FastTypeCheck(vv reflect.Value, defaultStr string, opts *Options) error {
	if defaultStr == "" || !vv.IsNil() {
		return nil
	}

	// 快速路径：简单字符串
	if len(defaultStr) == 0 {
		return nil
	}

	firstChar := defaultStr[0]
	if firstChar != '{' && firstChar != '[' {
		vv.Set(reflect.ValueOf(defaultStr))
		return nil
	}

	// JSON 格式（简化处理）
	var result interface{}
	if err := jsonUnmarshal([]byte(defaultStr), &result); err == nil {
		vv.Set(reflect.ValueOf(result))
	}

	return nil
}

// 优化方案 10：减少字符串操作
func setInterfaceDefault_Opt10_LessStringOps(vv reflect.Value, defaultStr string, opts *Options) error {
	if defaultStr == "" || !vv.IsNil() {
		return nil
	}

	// 单次字符检查代替 Contains
	if len(defaultStr) > 0 && (defaultStr[0] == '{' || defaultStr[0] == '[') {
		var result interface{}
		if err := jsonUnmarshal([]byte(defaultStr), &result); err == nil {
			vv.Set(reflect.ValueOf(result))
		}
	} else {
		vv.Set(reflect.ValueOf(defaultStr))
	}

	return nil
}

// ========== Channel 优化方案 ==========

// setChanDefault_Original 原始实现
func setChanDefault_Original(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() && defaultStr != "" {
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

// 优化方案 11：预解析缓冲区大小
func setChanDefault_Opt11_PreParsed(vv reflect.Value, defaultStr string, opts *Options) error {
	if !vv.IsNil() || defaultStr == "" {
		return nil
	}

	// 快速路径：0
	if defaultStr == "0" {
		vv.Set(reflect.MakeChan(vv.Type(), 0))
		return nil
	}

	// 解析缓冲区大小
	bufSize, err := strconv.Atoi(defaultStr)
	if err != nil {
		return handleError(fmt.Sprintf("invalid channel buffer size: %s", defaultStr), opts.ErrorMode)
	}

	vv.Set(reflect.MakeChan(vv.Type(), bufSize))
	return nil
}

// 优化方案 12：内联优化
func setChanDefault_Opt12_Inlined(vv reflect.Value, defaultStr string, opts *Options) error {
	if vv.IsNil() && defaultStr != "" {
		var bufSize int
		var err error
		if defaultStr == "0" {
			bufSize = 0
		} else {
			bufSize, err = strconv.Atoi(defaultStr)
			if err != nil {
				return handleError(fmt.Sprintf("invalid channel buffer size: %s", defaultStr), opts.ErrorMode)
			}
		}
		vv.Set(reflect.MakeChan(vv.Type(), bufSize))
	}
	return nil
}

// ========== Benchmark 测试 ==========

// Time benchmarks
func BenchmarkTimeDefault_Original(b *testing.B) {
	var tt TimeTestStruct
	vv := reflect.ValueOf(&tt.TimeField).Elem()
	for i := 0; i < b.N; i++ {
		setTimeDefault_Original(vv, "2024-01-15T10:30:00Z", &Options{})
	}
}

func BenchmarkTimeDefault_Opt1_UnixFastPath(b *testing.B) {
	var tt TimeTestStruct
	vv := reflect.ValueOf(&tt.TimeField).Elem()
	for i := 0; i < b.N; i++ {
		setTimeDefault_Opt1_UnixFastPath(vv, "1705318200", &Options{})
	}
}

func BenchmarkTimeDefault_Opt2_LessReflection(b *testing.B) {
	var tt TimeTestStruct
	vv := reflect.ValueOf(&tt.TimeField).Elem()
	for i := 0; i < b.N; i++ {
		setTimeDefault_Opt2_LessReflection(vv, "2024-01-15T10:30:00Z", &Options{})
	}
}

func BenchmarkTimeDefault_Opt3_PreCheck(b *testing.B) {
	var tt TimeTestStruct
	vv := reflect.ValueOf(&tt.TimeField).Elem()
	for i := 0; i < b.N; i++ {
		setTimeDefault_Opt3_PreCheck(vv, "2024-01-15T10:30:00Z", &Options{})
	}
}

func BenchmarkTimeDefault_Opt4_ShortString(b *testing.B) {
	var tt TimeTestStruct
	vv := reflect.ValueOf(&tt.TimeField).Elem()
	for i := 0; i < b.N; i++ {
		setTimeDefault_Opt4_ShortString(vv, "2024-01-15", &Options{})
	}
}

func BenchmarkTimeDefault_Opt5_CachedLayouts(b *testing.B) {
	var tt TimeTestStruct
	vv := reflect.ValueOf(&tt.TimeField).Elem()
	for i := 0; i < b.N; i++ {
		setTimeDefault_Opt5_CachedLayouts(vv, "2024-01-15T10:30:00Z", &Options{})
	}
}

// Ptr benchmarks
func BenchmarkPtrDefault_Original(b *testing.B) {
	var tt PtrTestStruct
	vv := reflect.ValueOf(&tt.PtrField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setPtrDefault_Original(vv, "hello", &Options{})
	}
}

func BenchmarkPtrDefault_Opt6_LessLoop(b *testing.B) {
	var tt PtrTestStruct
	vv := reflect.ValueOf(&tt.PtrField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setPtrDefault_Opt6_LessLoop(vv, "hello", &Options{})
	}
}

func BenchmarkPtrDefault_Opt7_EarlyReturn(b *testing.B) {
	var tt PtrTestStruct
	vv := reflect.ValueOf(&tt.PtrField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setPtrDefault_Opt7_EarlyReturn(vv, "hello", &Options{})
	}
}

func BenchmarkPtrDefault_Opt8_BatchedReflection(b *testing.B) {
	var tt PtrTestStruct
	vv := reflect.ValueOf(&tt.PtrField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setPtrDefault_Opt8_BatchedReflection(vv, "hello", &Options{})
	}
}

// Interface benchmarks
func BenchmarkInterfaceDefault_Original(b *testing.B) {
	var tt InterfaceTestStruct
	vv := reflect.ValueOf(&tt.InterfaceField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setInterfaceDefault_Original(vv, "test", &Options{})
	}
}

func BenchmarkInterfaceDefault_Opt9_FastTypeCheck(b *testing.B) {
	var tt InterfaceTestStruct
	vv := reflect.ValueOf(&tt.InterfaceField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setInterfaceDefault_Opt9_FastTypeCheck(vv, "test", &Options{})
	}
}

func BenchmarkInterfaceDefault_Opt10_LessStringOps(b *testing.B) {
	var tt InterfaceTestStruct
	vv := reflect.ValueOf(&tt.InterfaceField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setInterfaceDefault_Opt10_LessStringOps(vv, "test", &Options{})
	}
}

// Channel benchmarks
func BenchmarkChanDefault_Original(b *testing.B) {
	var tt ChanTestStruct
	vv := reflect.ValueOf(&tt.ChanField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setChanDefault_Original(vv, "10", &Options{})
	}
}

func BenchmarkChanDefault_Opt11_PreParsed(b *testing.B) {
	var tt ChanTestStruct
	vv := reflect.ValueOf(&tt.ChanField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setChanDefault_Opt11_PreParsed(vv, "10", &Options{})
	}
}

func BenchmarkChanDefault_Opt12_Inlined(b *testing.B) {
	var tt ChanTestStruct
	vv := reflect.ValueOf(&tt.ChanField).Elem()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		setChanDefault_Opt12_Inlined(vv, "10", &Options{})
	}
}

func BenchmarkOriginalSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s BenchmarkSimple
		SetDefaults(&s)
	}
}

func BenchmarkOriginalComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var c BenchmarkComplex
		SetDefaults(&c)
	}
}

func BenchmarkOriginalVeryComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v BenchmarkVeryComplex
		SetDefaults(&v)
	}
}

func BenchmarkOptimizedSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s BenchmarkSimple
		SetDefaultsOptimized(&s)
	}
}

func BenchmarkOptimizedComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var c BenchmarkComplex
		SetDefaultsOptimized(&c)
	}
}

func BenchmarkOptimizedVeryComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v BenchmarkVeryComplex
		SetDefaultsOptimized(&v)
	}
}

func BenchmarkCompareSimple(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		var s BenchmarkSimple
		for i := 0; i < b.N; i++ {
			SetDefaults(&s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		var s BenchmarkSimple
		for i := 0; i < b.N; i++ {
			SetDefaultsOptimized(&s)
		}
	})
}

func BenchmarkCompareComplex(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		var c BenchmarkComplex
		for i := 0; i < b.N; i++ {
			SetDefaults(&c)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		var c BenchmarkComplex
		for i := 0; i < b.N; i++ {
			SetDefaultsOptimized(&c)
		}
	})
}

func BenchmarkCompareVeryComplex(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		var v BenchmarkVeryComplex
		for i := 0; i < b.N; i++ {
			SetDefaults(&v)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		var v BenchmarkVeryComplex
		for i := 0; i < b.N; i++ {
			SetDefaultsOptimized(&v)
		}
	})
}
