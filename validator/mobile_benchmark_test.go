package validator

import (
	"reflect"
	"strings"
	"testing"
)

// 测试数据
var (
	validMobile   = "13812345678"
	invalidPrefix = "12812345678" // 第二位不是3-9
	invalidLen    = "138123456"   // 长度不足
	invalidChar   = "1381234567a" // 包含非数字
	emptyString   = ""
)

// ========== 方案1: 当前正则（基线） ==========
func validateMobile01(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}
	return mobileRegex.MatchString(mobile)
}

// ========== 方案2: 手动逐字符 ASCII 比较 ==========
func validateMobile02(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 长度检查
	if len(mobile) != 11 {
		return false
	}

	// 第一位必须是'1'
	if mobile[0] != '1' {
		return false
	}

	// 第二位必须是3-9
	if mobile[1] < '3' || mobile[1] > '9' {
		return false
	}

	// 后9位必须是数字
	for i := 2; i < 11; i++ {
		if mobile[i] < '0' || mobile[i] > '9' {
			return false
		}
	}

	return true
}

// ========== 方案3: 字节切片遍历（优化索引循环） ==========
func validateMobile03(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	s := []byte(mobile)
	l := len(s)

	// 长度检查
	if l != 11 {
		return false
	}

	// 第一位必须是'1'
	if s[0] != '1' {
		return false
	}

	// 第二位必须是3-9
	if s[1] < '3' || s[1] > '9' {
		return false
	}

	// 后9位必须是数字
	for i := 2; i < l; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

// ========== 方案4: 前缀快速路径 + 后续验证 ==========
func validateMobile04(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 快速失败：长度和前缀
	if len(mobile) != 11 || mobile[0] != '1' {
		return false
	}

	// 第二位检查（使用查表法优化）
	secondDigit := mobile[1]
	if secondDigit < '3' || secondDigit > '9' {
		return false
	}

	// 后续数字检查
	for i := 2; i < 11; i++ {
		c := mobile[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

// ========== 方案5: 查找表（第二位数字3-9） ==========
var validSecondDigits = [256]bool{
	'3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
}

func validateMobile05(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
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

// ========== 方案6: 长度检查 + 字符类检查（提前计算） ==========
func validateMobile06(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	l := len(mobile)
	if l != 11 {
		return false
	}

	// 检查所有字符是否为数字
	for i := 0; i < l; i++ {
		c := mobile[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 第一位必须是1
	if mobile[0] != '1' {
		return false
	}

	// 第二位必须是3-9
	if mobile[1] < '3' || mobile[1] > '9' {
		return false
	}

	return true
}

// ========== 方案7: 使用 strings.Index 快速失败 ==========
func validateMobile07(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 快速检查是否包含非数字字符
	if strings.IndexAny(mobile, "0123456789") != -1 {
		// 这个逻辑有问题，换一个思路
		// 直接检查每个字符
	}

	// 回退到简单方案
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

// ========== 方案8: 避免正则编译，使用 regexp.MustCompile 缓存（已是最优）==========
// 这个方案实际上就是当前方案，因为已经使用了 MustCompile

// ========== 方案9: 组合优化（长度+前缀+数字检查） ==========
func validateMobile09(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 组合所有检查
	if len(mobile) != 11 ||
		mobile[0] != '1' ||
		mobile[1] < '3' ||
		mobile[1] > '9' {
		return false
	}

	// 检查后9位
	for i := 2; i < 11; i++ {
		c := mobile[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

// ========== 方案10: 纯 ASCII 数字范围检查 ==========
func validateMobile10(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 使用 range 遍历字符串（UTF-8安全）
	count := 0
	for _, c := range mobile {
		if c < '0' || c > '9' {
			return false
		}
		count++
	}

	// 检查长度
	if count != 11 {
		return false
	}

	// 第一位必须是1
	if mobile[0] != '1' {
		return false
	}

	// 第二位必须是3-9
	if mobile[1] < '3' || mobile[1] > '9' {
		return false
	}

	return true
}

// ========== 方案11: 手动展开循环（循环展开优化） ==========
func validateMobile11(fl FieldLevel) bool {
	mobile := fl.Field().String()
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

// ========== 方案12: 位运算优化（检查数字） ==========
func validateMobile12(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	if len(mobile) != 11 {
		return false
	}

	if mobile[0] != '1' {
		return false
	}

	// 第二位：'3'(51) 到 '9'(57)
	if mobile[1] < 51 || mobile[1] > 57 {
		return false
	}

	// 使用位运算检查数字: '0'(48) 到 '9'(57)
	for i := 2; i < 11; i++ {
		c := mobile[i] - 48
		if c > 9 {
			return false
		}
	}

	return true
}

// ========== 基准测试 ==========

// 基准测试 - 有效手机号
func BenchmarkValidateMobile01_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkValidateMobile02_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkValidateMobile03_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile03(fl)
	}
}

func BenchmarkValidateMobile04_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile04(fl)
	}
}

func BenchmarkValidateMobile05_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile05(fl)
	}
}

func BenchmarkValidateMobile06_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile06(fl)
	}
}

func BenchmarkValidateMobile07_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile07(fl)
	}
}

func BenchmarkValidateMobile09_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile09(fl)
	}
}

func BenchmarkValidateMobile10_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile10(fl)
	}
}

func BenchmarkValidateMobile11_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

func BenchmarkValidateMobile12_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(validMobile)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile12(fl)
	}
}

// 基准测试 - 无效前缀
func BenchmarkValidateMobile01_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkValidateMobile02_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkValidateMobile03_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile03(fl)
	}
}

func BenchmarkValidateMobile04_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile04(fl)
	}
}

func BenchmarkValidateMobile05_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile05(fl)
	}
}

func BenchmarkValidateMobile06_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile06(fl)
	}
}

func BenchmarkValidateMobile09_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile09(fl)
	}
}

func BenchmarkValidateMobile11_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

func BenchmarkValidateMobile12_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidPrefix)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile12(fl)
	}
}

// 基准测试 - 无效长度
func BenchmarkValidateMobile01_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkValidateMobile02_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkValidateMobile03_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile03(fl)
	}
}

func BenchmarkValidateMobile04_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile04(fl)
	}
}

func BenchmarkValidateMobile05_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile05(fl)
	}
}

func BenchmarkValidateMobile06_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile06(fl)
	}
}

func BenchmarkValidateMobile09_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile09(fl)
	}
}

func BenchmarkValidateMobile11_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

func BenchmarkValidateMobile12_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf(invalidLen)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile12(fl)
	}
}
