package validator

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// BenchmarkValidate_SingleField 单字段验证性能测试
func BenchmarkValidate_SingleField(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name string `validate:"required"`
	}
	user := User{Name: "test"}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(user)
	}
}

// BenchmarkValidate_SingleField_Invalid 单字段验证失败性能测试
func BenchmarkValidate_SingleField_Invalid(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name string `validate:"required"`
	}
	user := User{Name: ""}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(user)
	}
}

// BenchmarkValidate_StringValidation email/url 验证性能测试
func BenchmarkValidate_StringValidation(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("Email", func(b *testing.B) {
		type Form struct {
			Email string `validate:"email"`
		}
		form := Form{Email: "user@example.com"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("URL", func(b *testing.B) {
		type Form struct {
			URL string `validate:"url"`
		}
		form := Form{URL: "https://example.com/path?query=value"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("EmailAndURL", func(b *testing.B) {
		type Form struct {
			Email string `validate:"required,email"`
			URL   string `validate:"required,url"`
		}
		form := Form{
			Email: "user@example.com",
			URL:   "https://example.com",
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})
}

// BenchmarkValidate_MultiField 多字段结构体验证性能测试
func BenchmarkValidate_MultiField(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("5Fields", func(b *testing.B) {
		type Form struct {
			Name     string `json:"name" validate:"required,min=2"`
			Email    string `json:"email" validate:"required,email"`
			Age      int    `json:"age" validate:"required,min=1,max=150"`
			Phone    string `json:"phone" validate:"required"`
			Password string `json:"password" validate:"required,min=6"`
		}
		form := Form{
			Name:     "Zhang San",
			Email:    "test@example.com",
			Age:      25,
			Phone:    "13812345678",
			Password: "MyPass123!",
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("10Fields", func(b *testing.B) {
		type Form struct {
			Name     string `json:"name" validate:"required,min=2"`
			Email    string `json:"email" validate:"required,email"`
			Age      int    `json:"age" validate:"required,min=1,max=150"`
			Phone    string `json:"phone" validate:"required"`
			Password string `json:"password" validate:"required,min=6"`
			Address  string `json:"address" validate:"required,min=5"`
			City     string `json:"city" validate:"required"`
			Country  string `json:"country" validate:"required,min=2"`
			ZipCode  string `json:"zip_code" validate:"required,len=6"`
			Website  string `json:"website" validate:"url"`
		}
		form := Form{
			Name:     "Zhang San",
			Email:    "test@example.com",
			Age:      25,
			Phone:    "13812345678",
			Password: "MyPass123!",
			Address:  "No.1 Main Street",
			City:     "Beijing",
			Country:  "China",
			ZipCode:  "100000",
			Website:  "https://example.com",
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})
}

// BenchmarkValidate_CustomValidator 自定义验证器性能测试
func BenchmarkValidate_CustomValidator(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("StrongPassword", func(b *testing.B) {
		type Form struct {
			Password string `validate:"strong_password"`
		}
		form := Form{Password: "MyPass123!"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})
}

// BenchmarkValidate_NestedStruct 嵌套结构体验证性能测试
func BenchmarkValidate_NestedStruct(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type Address struct {
		Street  string `json:"street" validate:"required,min=5"`
		City    string `json:"city" validate:"required"`
		ZipCode string `json:"zip_code" validate:"required,len=6"`
	}

	type Company struct {
		Name    string  `json:"name" validate:"required,min=2"`
		Address Address `json:"address"`
	}

	type Employee struct {
		Name    string  `json:"name" validate:"required,min=2"`
		Email   string  `json:"email" validate:"required,email"`
		Age     int     `json:"age" validate:"required,min=18,max=65"`
		Company Company `json:"company"`
		Home    Address `json:"home"`
	}

	employee := Employee{
		Name:  "Zhang San",
		Email: "zhangsan@example.com",
		Age:   30,
		Company: Company{
			Name: "Tech Corp",
			Address: Address{
				Street:  "No.100 Innovation Road",
				City:    "Beijing",
				ZipCode: "100000",
			},
		},
		Home: Address{
			Street:  "No.1 Residential Area",
			City:    "Beijing",
			ZipCode: "100001",
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(employee)
	}
}

// BenchmarkValidate_Batch 批量验证性能测试
func BenchmarkValidate_Batch(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name  string `json:"name" validate:"required,min=2"`
		Email string `json:"email" validate:"required,email"`
		Age   int    `json:"age" validate:"required,min=1"`
	}

	// 预构建 100 个对象
	users := make([]User, 100)
	for i := range users {
		users[i] = User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + i%50,
		}
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := range users {
			_ = v.Struct(users[j])
		}
	}
}

// BenchmarkValidate_ChineseMessage 中文错误消息性能测试
func BenchmarkValidate_ChineseMessage(b *testing.B) {
	v, err := New(WithLocale("zh"))
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name  string `json:"name" validate:"required,min=2"`
		Email string `json:"email" validate:"required,email"`
	}

	// 使用无效数据触发中文错误消息
	user := User{
		Name:  "",
		Email: "invalid",
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(user)
	}
}

// BenchmarkValidate_Var 单变量验证性能测试
func BenchmarkValidate_Var(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("Required", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Var("hello", "required")
		}
	})

	b.Run("Email", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Var("user@example.com", "email")
		}
	})

	b.Run("MinMax", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Var(25, "min=1,max=100")
		}
	})
}

// BenchmarkValidate_Parallel 并发验证性能测试
func BenchmarkValidate_Parallel(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name  string `json:"name" validate:"required,min=2"`
		Email string `json:"email" validate:"required,email"`
	}

	user := User{
		Name:  "Zhang San",
		Email: "test@example.com",
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = v.Struct(user)
		}
	})
}

// ============================================================================
// 消息格式化性能优化基准测试
// ============================================================================

var benchmarkFieldError = &FieldError{
	Field: "Email",
	Tag:   "email",
	Param: "",
	Value: "invalid-email",
}

var benchmarkTemplates = map[string]string{
	"simple":  "{field} 格式无效",
	"param":   "{field} 长度不能小于 {param}",
	"value":   "{field} 的值 {value} 无效",
	"complex": "{field} 验证失败: {tag} (参数: {param}, 值: {value})",
	"noplace": "验证失败",
}

// 方案1: 基准 - 当前实现 (strings.ReplaceAll)
func benchmarkFmt_Baseline(template string, err *FieldError) string {
	msg := template
	msg = strings.ReplaceAll(msg, "{field}", err.Field)
	msg = strings.ReplaceAll(msg, "{tag}", err.Tag)
	msg = strings.ReplaceAll(msg, "{param}", err.Param)
	if err.Value != nil {
		msg = strings.ReplaceAll(msg, "{value}", fmt.Sprintf("%v", err.Value))
	} else {
		msg = strings.ReplaceAll(msg, "{value}", "")
	}
	return msg
}

// 方案2: strings.Builder 单次遍历
func benchmarkFmt_Builder(template string, err *FieldError) string {
	var b strings.Builder
	b.Grow(len(template) + 50)

	repls := []struct{ p, v string }{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
	}
	vstr := ""
	if err.Value != nil {
		vstr = fmt.Sprintf("%v", err.Value)
	}
	repls = append(repls, struct{ p, v string }{"{value}", vstr})

	i := 0
	for i < len(template) {
		matched := false
		for _, r := range repls {
			if strings.HasPrefix(template[i:], r.p) {
				b.WriteString(r.v)
				i += len(r.p)
				matched = true
				break
			}
		}
		if !matched {
			b.WriteByte(template[i])
			i++
		}
	}
	return b.String()
}

// 方案3: 字节切片预分配
func benchmarkFmt_ByteSlice(template string, err *FieldError) string {
	result := make([]byte, 0, len(template)+50)

	repls := []struct{ p, v string }{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
	}
	vstr := ""
	if err.Value != nil {
		vstr = fmt.Sprintf("%v", err.Value)
	}
	repls = append(repls, struct{ p, v string }{"{value}", vstr})

	i := 0
	for i < len(template) {
		matched := false
		for _, r := range repls {
			if strings.HasPrefix(template[i:], r.p) {
				result = append(result, r.v...)
				i += len(r.p)
				matched = true
				break
			}
		}
		if !matched {
			result = append(result, template[i])
			i++
		}
	}
	return string(result)
}

// 方案4: 内联优化
func benchmarkFmt_Inlined(template string, err *FieldError) string {
	result := make([]byte, 0, len(template)+50)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) {
			if template[i:i+7] == "{field}" {
				result = append(result, err.Field...)
				i += 7
				continue
			}
			if template[i:i+7] == "{param}" {
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}
		if i+5 <= len(template) {
			if template[i:i+5] == "{tag}" {
				result = append(result, err.Tag...)
				i += 5
				continue
			}
		}
		if i+6 <= len(template) {
			if template[i:i+6] == "{value}" {
				if err.Value != nil {
					result = append(result, fmt.Sprintf("%v", err.Value)...)
				}
				i += 6
				continue
			}
		}
		result = append(result, template[i])
		i++
	}
	return string(result)
}

// 方案5: 快速路径优化
func benchmarkFmt_FastPath(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}
	result := make([]byte, 0, len(template)+50)

	i := 0
	repls := []struct{ p, v string }{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
	}
	vstr := ""
	if err.Value != nil {
		vstr = fmt.Sprintf("%v", err.Value)
	}
	repls = append(repls, struct{ p, v string }{"{value}", vstr})

	for i < len(template) {
		matched := false
		for _, r := range repls {
			if strings.HasPrefix(template[i:], r.p) {
				result = append(result, r.v...)
				i += len(r.p)
				matched = true
				break
			}
		}
		if !matched {
			result = append(result, template[i])
			i++
		}
	}
	return string(result)
}

// 方案6: sync.Pool 复用
var fmtBuilderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func benchmarkFmt_Pool(template string, err *FieldError) string {
	b := fmtBuilderPool.Get().(*strings.Builder)
	defer func() {
		b.Reset()
		fmtBuilderPool.Put(b)
	}()

	b.Grow(len(template) + 50)

	repls := []struct{ p, v string }{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
	}
	vstr := ""
	if err.Value != nil {
		vstr = fmt.Sprintf("%v", err.Value)
	}
	repls = append(repls, struct{ p, v string }{"{value}", vstr})

	i := 0
	for i < len(template) {
		matched := false
		for _, r := range repls {
			if strings.HasPrefix(template[i:], r.p) {
				b.WriteString(r.v)
				i += len(r.p)
				matched = true
				break
			}
		}
		if !matched {
			b.WriteByte(template[i])
			i++
		}
	}
	return b.String()
}

// 方案7: 容量预估优化
func benchmarkFmt_Estimated(template string, err *FieldError) string {
	estSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	if err.Value != nil {
		estSize += 20
	}

	result := make([]byte, 0, estSize)

	repls := []struct{ p, v string }{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
	}
	vstr := ""
	if err.Value != nil {
		vstr = fmt.Sprintf("%v", err.Value)
	}
	repls = append(repls, struct{ p, v string }{"{value}", vstr})

	i := 0
	for i < len(template) {
		matched := false
		for _, r := range repls {
			if strings.HasPrefix(template[i:], r.p) {
				result = append(result, r.v...)
				i += len(r.p)
				matched = true
				break
			}
		}
		if !matched {
			result = append(result, template[i])
			i++
		}
	}
	return string(result)
}

// 方案8: ReplaceAll 优化版
func benchmarkFmt_ReplaceAllOpt(template string, err *FieldError) string {
	msg := template
	if strings.Contains(msg, "{field}") {
		msg = strings.ReplaceAll(msg, "{field}", err.Field)
	}
	if strings.Contains(msg, "{tag}") {
		msg = strings.ReplaceAll(msg, "{tag}", err.Tag)
	}
	if strings.Contains(msg, "{param}") {
		msg = strings.ReplaceAll(msg, "{param}", err.Param)
	}
	if strings.Contains(msg, "{value}") {
		if err.Value != nil {
			msg = strings.ReplaceAll(msg, "{value}", fmt.Sprintf("%v", err.Value))
		} else {
			msg = strings.ReplaceAll(msg, "{value}", "")
		}
	}
	return msg
}

// 方案9: 混合优化（快速路径 + 内联）
func benchmarkFmt_Hybrid(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	estSize := len(template) + len(err.Field) + len(err.Tag) + len(err.Param) + 50
	if err.Value != nil {
		estSize += 20
	}

	result := make([]byte, 0, estSize)

	i := 0
	for i < len(template) {
		if i+7 <= len(template) {
			if template[i:i+7] == "{field}" {
				result = append(result, err.Field...)
				i += 7
				continue
			}
			if template[i:i+7] == "{param}" {
				result = append(result, err.Param...)
				i += 7
				continue
			}
		}
		if i+5 <= len(template) {
			if template[i:i+5] == "{tag}" {
				result = append(result, err.Tag...)
				i += 5
				continue
			}
		}
		if i+6 <= len(template) {
			if template[i:i+6] == "{value}" {
				if err.Value != nil {
					result = append(result, fmt.Sprintf("%v", err.Value)...)
				}
				i += 6
				continue
			}
		}
		result = append(result, template[i])
		i++
	}
	return string(result)
}

// 方案10: 最小分配优化
func benchmarkFmt_MinAlloc(template string, err *FieldError) string {
	if !strings.Contains(template, "{") {
		return template
	}

	repls := [4][2]string{
		{"{field}", err.Field},
		{"{tag}", err.Tag},
		{"{param}", err.Param},
		{"{value}", ""},
	}
	if err.Value != nil {
		repls[3][1] = fmt.Sprintf("%v", err.Value)
	}

	result := make([]byte, 0, len(template)+64)

	i := 0
	for i < len(template) {
		matched := false
		for j := 0; j < 4; j++ {
			if strings.HasPrefix(template[i:], repls[j][0]) {
				result = append(result, repls[j][1]...)
				i += len(repls[j][0])
				matched = true
				break
			}
		}
		if !matched {
			result = append(result, template[i])
			i++
		}
	}
	return string(result)
}

// ============================================================================
// 基准测试入口
// ============================================================================

func BenchmarkFormatMessage_Simple(b *testing.B) {
	tmpl := benchmarkTemplates["simple"]
	b.Run("Baseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Baseline(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Builder", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Builder(tmpl, benchmarkFieldError)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_ByteSlice(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Inlined(tmpl, benchmarkFieldError)
		}
	})
	b.Run("FastPath", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_FastPath(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Pool(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Estimated", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Estimated(tmpl, benchmarkFieldError)
		}
	})
	b.Run("ReplaceAllOpt", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_ReplaceAllOpt(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Hybrid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Hybrid(tmpl, benchmarkFieldError)
		}
	})
	b.Run("MinAlloc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_MinAlloc(tmpl, benchmarkFieldError)
		}
	})
}

func BenchmarkFormatMessage_Complex(b *testing.B) {
	tmpl := benchmarkTemplates["complex"]
	b.Run("Baseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Baseline(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Builder", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Builder(tmpl, benchmarkFieldError)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_ByteSlice(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Inlined(tmpl, benchmarkFieldError)
		}
	})
	b.Run("FastPath", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_FastPath(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Pool(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Estimated", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Estimated(tmpl, benchmarkFieldError)
		}
	})
	b.Run("ReplaceAllOpt", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_ReplaceAllOpt(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Hybrid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Hybrid(tmpl, benchmarkFieldError)
		}
	})
	b.Run("MinAlloc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_MinAlloc(tmpl, benchmarkFieldError)
		}
	})
}

func BenchmarkFormatMessage_NoPlaceholder(b *testing.B) {
	tmpl := benchmarkTemplates["noplace"]
	b.Run("Baseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Baseline(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Builder", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Builder(tmpl, benchmarkFieldError)
		}
	})
	b.Run("ByteSlice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_ByteSlice(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Inlined(tmpl, benchmarkFieldError)
		}
	})
	b.Run("FastPath", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_FastPath(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Pool(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Estimated", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Estimated(tmpl, benchmarkFieldError)
		}
	})
	b.Run("ReplaceAllOpt", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_ReplaceAllOpt(tmpl, benchmarkFieldError)
		}
	})
	b.Run("Hybrid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_Hybrid(tmpl, benchmarkFieldError)
		}
	})
	b.Run("MinAlloc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = benchmarkFmt_MinAlloc(tmpl, benchmarkFieldError)
		}
	})
}
