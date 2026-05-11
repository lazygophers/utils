package validator

import (
	"reflect"
	"testing"
)

// newMockFieldLevel 创建mock FieldLevel用于测试
func newMockFieldLevel(value string) FieldLevel {
	return mockFieldLevel{
		field: reflect.ValueOf(value),
	}
}

// 生成测试银行卡号（固定种子保证可重复）
func genBankCards(n int) []string {
	cards := make([]string, n)
	// 有效的银行卡号（通过Luhn算法）
	validCards := []string{
		"6222021234567890",    // 16位
		"6228481234567890123", // 19位
		"4000001234567890",    // 16位
		"5555551234567890",    // 16位
		"378282246310005",     // 15位
		"6011111111111117",    // 16位
		"3530111333300000",    // 16位
		"1234567890123456789", // 19位
	}

	for i := 0; i < n; i++ {
		cards[i] = validCards[i%len(validCards)]
	}
	return cards
}

// 生成无效银行卡号（测试快速失败）
func genInvalidBankCards(n int) []string {
	cards := make([]string, n)
	invalidCases := []string{
		"abcd1234567890",     // 包含字母
		"123",                 // 太短
		"12345678901234567890", // 太长
		"6222021234567891",    // Luhn失败
		"",                    // 空字符串
		"62220@1234567890",    // 特殊字符
	}

	for i := 0; i < n; i++ {
		cards[i] = invalidCases[i%len(invalidCases)]
	}
	return cards
}

// Baseline: 当前实现
func BenchmarkValidateBankCard_Current_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCard(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Current_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCard(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Current_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCard(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案1: 字节级检查 + 手动Luhn（避免strconv.Atoi）
func BenchmarkValidateBankCard_Opt1_ByteManualLuhn_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt1ByteManualLuhn(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt1_ByteManualLuhn_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt1ByteManualLuhn(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt1_ByteManualLuhn_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt1ByteManualLuhn(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案2: 字节级 + 查找表优化Luhn
func BenchmarkValidateBankCard_Opt2_LookupTable_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt2LookupTable(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt2_LookupTable_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt2LookupTable(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt2_LookupTable_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt2LookupTable(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案3: 预计算Luhn双倍值表
func BenchmarkValidateBankCard_Opt3_PrecomputedDoubles_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt3PrecomputedDoubles(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt3_PrecomputedDoubles_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt3PrecomputedDoubles(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt3_PrecomputedDoubles_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt3PrecomputedDoubles(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案4: 快速失败优化（长度前置检查）
func BenchmarkValidateBankCard_Opt4_FastFail_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt4FastFail(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt4_FastFail_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt4FastFail(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt4_FastFail_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt4FastFail(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案5: 索引循环（避免range）
func BenchmarkValidateBankCard_Opt5_IndexLoop_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt5IndexLoop(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt5_IndexLoop_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt5IndexLoop(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt5_IndexLoop_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt5IndexLoop(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案6: ASCII范围检查优化
func BenchmarkValidateBankCard_Opt6_ASCII_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt6ASCII(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt6_ASCII_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt6ASCII(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt6_ASCII_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt6ASCII(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案7: 单次遍历优化
func BenchmarkValidateBankCard_Opt7_SinglePass_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt7SinglePass(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt7_SinglePass_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt7SinglePass(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt7_SinglePass_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt7SinglePass(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案8: 位运算优化
func BenchmarkValidateBankCard_Opt8_BitOps_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt8BitOps(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt8_BitOps_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt8BitOps(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt8_BitOps_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt8BitOps(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案9: 反向遍历优化
func BenchmarkValidateBankCard_Opt9_Reverse_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt9Reverse(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt9_Reverse_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt9Reverse(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt9_Reverse_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt9Reverse(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案10: 组合优化（字节+手动Luhn+快速失败+ASCII）
func BenchmarkValidateBankCard_Opt10_Combined_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10Combined(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt10_Combined_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10Combined(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt10_Combined_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10Combined(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案11: 无分支Luhn
func BenchmarkValidateBankCard_Opt11_Branchless_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt11Branchless(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt11_Branchless_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt11Branchless(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt11_Branchless_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt11Branchless(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 方案12: SIMD启发式（批量处理）
func BenchmarkValidateBankCard_Opt12_SimdInspired_Valid_Small(b *testing.B) {
	cards := genBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt12SimdInspired(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt12_SimdInspired_Valid_Medium(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt12SimdInspired(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt12_SimdInspired_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt12SimdInspired(newMockFieldLevel(cards[i%len(cards)]))
	}
}

// 内存分配测试
func BenchmarkValidateBankCard_Current_Allocs(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCard(newMockFieldLevel(cards[i%len(cards)]))
	}
}

func BenchmarkValidateBankCard_Opt10_Combined_Allocs(b *testing.B) {
	cards := genBankCards(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10Combined(newMockFieldLevel(cards[i%len(cards)]))
	}
}
