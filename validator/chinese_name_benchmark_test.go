package validator

import (
	"testing"
)

// BenchmarkChineseName_Validation 性能基准测试
func BenchmarkChineseName_Validation(b *testing.B) {
	v, _ := New()
	
	type Form struct {
		Name string `validate:"chinese_name"`
	}
	
	testCases := []struct {
		name string
		desc string
	}{
		{"张三", "简单二字姓名"},
		{"司马青衫", "四字姓名"},
		{"欧阳修", "三字姓名"},
		{"诸葛亮", "复姓姓名"},
		{"张", "无效短姓名"},
		{"张三李四王五赵六", "无效长姓名"},
	}
	
	for _, tc := range testCases {
		b.Run(tc.desc, func(b *testing.B) {
			form := Form{Name: tc.name}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = v.Struct(form)
			}
		})
	}
}
