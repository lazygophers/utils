package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/lazygophers/utils/validator"
)

// 模拟 FieldLevel 接口
type testFieldLevel struct {
	value reflect.Value
}

func (f *testFieldLevel) Field() reflect.Value {
	return f.value
}

func (f *testFieldLevel) Top() reflect.Value {
	return f.value
}

func (f *testFieldLevel) Parent() reflect.Value {
	return f.value
}

func (f *testFieldLevel) FieldName() string {
	return "test"
}

func (f *testFieldLevel) StructFieldName() string {
	return "test"
}

func (f *testFieldLevel) Param() string {
	return ""
}

func (f *testFieldLevel) GetTag(key string) string {
	return ""
}

// 定义所有优化方案
func validateMobile01(mobile string) bool {
	if mobile == "" {
		return false
	}
	return validator.MobileRegex().MatchString(mobile)
}

func validateMobile02(mobile string) bool {
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
	}

	if mobile[1] < '3' || mobile[1] > '9' {
		return false
	}

	for i := 2; i < 11; i++ {
		if mobile[i] < '0' || mobile[i] > '9' {
			return false
		}
	}

	return true
}

func validateMobile03(mobile string) bool {
	if mobile == "" {
		return false
	}

	s := []byte(mobile)
	l := len(s)

	if l != 11 {
		return false
	}

	if s[0] != '1' {
		return false
	}

	if s[1] < '3' || s[1] > '9' {
		return false
	}

	for i := 2; i < l; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

func validateMobile04(mobile string) bool {
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 || mobile[0] != '1' {
		return false
	}

	secondDigit := mobile[1]
	if secondDigit < '3' || secondDigit > '9' {
		return false
	}

	for i := 2; i < 11; i++ {
		c := mobile[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

func validateMobile05(mobile string) bool {
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
	}

	// 查找表
	validSecondDigits := [256]bool{
		'3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
	}

	if !validSecondDigits[mobile[1]] {
		return false
	}

	for i := 2; i < 11; i++ {
		if mobile[i] < '0' || mobile[i] > '9' {
			return false
		}
	}

	return true
}

func validateMobile09(mobile string) bool {
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 ||
		mobile[0] != '1' ||
		mobile[1] < '3' ||
		mobile[1] > '9' {
		return false
	}

	for i := 2; i < 11; i++ {
		c := mobile[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

func validateMobile10(mobile string) bool {
	if mobile == "" {
		return false
	}

	count := 0
	for _, c := range mobile {
		if c < '0' || c > '9' {
			return false
		}
		count++
	}

	if count != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
	}

	if mobile[1] < '3' || mobile[1] > '9' {
		return false
	}

	return true
}

func validateMobile11(mobile string) bool {
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
	}

	if mobile[1] < '3' || mobile[1] > '9' {
		return false
	}

	// 手动展开循环
	if mobile[2] < '0' || mobile[2] > '9' {
		return false
	}
	if mobile[3] < '0' || mobile[3] > '9' {
		return false
	}
	if mobile[4] < '0' || mobile[4] > '9' {
		return false
	}
	if mobile[5] < '0' || mobile[5] > '9' {
		return false
	}
	if mobile[6] < '0' || mobile[6] > '9' {
		return false
	}
	if mobile[7] < '0' || mobile[7] > '9' {
		return false
	}
	if mobile[8] < '0' || mobile[8] > '9' {
		return false
	}
	if mobile[9] < '0' || mobile[9] > '9' {
		return false
	}
	if mobile[10] < '0' || mobile[10] > '9' {
		return false
	}

	return true
}

func validateMobile12(mobile string) bool {
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
	}

	if mobile[1] < 51 || mobile[1] > 57 {
		return false
	}

	for i := 2; i < 11; i++ {
		c := mobile[i] - 48
		if c > 9 {
			return false
		}
	}

	return true
}

// 基准测试函数
func benchmarkMobile(b *testing.B, fn func(string) string, input string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(input)
	}
}

func main() {
	validMobile := "13812345678"
	invalidPrefix := "12812345678"
	invalidLen := "138123456"

	// 验证所有函数的正确性
	functions := map[string]func(string) bool{
		"01_Regex基线":    validateMobile01,
		"02_手动逐字符":    validateMobile02,
		"03_字节切片":     validateMobile03,
		"04_快速路径":     validateMobile04,
		"05_查找表":      validateMobile05,
		"09_组合优化":     validateMobile09,
		"10_Range遍历":  validateMobile10,
		"11_循环展开":     validateMobile11,
		"12_位运算优化":    validateMobile12,
	}

	fmt.Println("=== 验证正确性 ===")
	for name, fn := range functions {
		valid := fn(validMobile)
		invalid1 := fn(invalidPrefix)
		invalid2 := fn(invalidLen)
		empty := fn("")

		if valid && !invalid1 && !invalid2 && !empty {
			fmt.Printf("✓ %s: 通过\n", name)
		} else {
			fmt.Printf("✗ %s: 失败 (valid=%v, invalidPrefix=%v, invalidLen=%v, empty=%v)\n",
				name, valid, invalid1, invalid2, empty)
		}
	}

	fmt.Println("\n=== 运行性能基准测试 ===")
	results := make(map[string][]testing.BenchmarkResult)

	for name, fn := range functions {
		fmt.Printf("运行 %s...\n", name)

		// 测试有效手机号
		r1 := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(validMobile)
			}
		})

		// 测试无效前缀
		r2 := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(invalidPrefix)
			}
		})

		// 测试无效长度
		r3 := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(invalidLen)
			}
		})

		results[name] = []testing.BenchmarkResult{r1, r2, r3}
	}

	// 输出结果
	fmt.Println("\n=== 基准测试结果 ===")
	fmt.Printf("%-15s %15s %15s %15s %15s %15s\n",
		"方案", "有效(ns/op)", "无效前缀(ns/op)", "无效长度(ns/op)", "有效(B/op)", "有效(allocs/op)")
	fmt.Println(strings.Repeat("-", 95))

	for name, res := range results {
		fmt.Printf("%-15s %15d %15d %15d %15d %15d\n",
			name,
			res[0].NsPerOp(),
			res[1].NsPerOp(),
			res[2].NsPerOp(),
			res[0].AllocsPerOp(),
			res[0].MemPerOp())
	}

	// 找出最优方案
	fmt.Println("\n=== 最优方案分析 ===")
	bestValid := results["02_手动逐字符"][0]
	bestValidName := "02_手动逐字符"

	for name, res := range results {
		if res[0].NsPerOp() < bestValid.NsPerOp() {
			bestValid = res[0]
			bestValidName = name
		}
	}

	fmt.Printf("有效手机号最快: %s (%d ns/op, %d B/op, %d allocs/op)\n",
		bestValidName, bestValid.NsPerOp(), bestValid.AllocsPerOp(), bestValid.MemPerOp())
}
