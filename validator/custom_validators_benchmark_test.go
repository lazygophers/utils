package validator

import (
	"testing"
)

// BenchmarkValidateBankCard 性能基准测试
func BenchmarkValidateBankCard_Valid16(b *testing.B) {
	fl := newMockFieldLevel("4532015112830366") // 有效16位Visa
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCard(fl)
	}
}

func BenchmarkValidateBankCard_Valid15(b *testing.B) {
	fl := newMockFieldLevel("378282246310005") // 有效15位Amex
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCard(fl)
	}
}

func BenchmarkValidateBankCard_Invalid_Short(b *testing.B) {
	fl := newMockFieldLevel("123") // 太短
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCard(fl)
	}
}

func BenchmarkValidateBankCard_Invalid_NonDigit(b *testing.B) {
	fl := newMockFieldLevel("abcd1234567890") // 包含字母
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCard(fl)
	}
}

func BenchmarkLuhnCheck_Valid(b *testing.B) {
	cardNo := "4532015112830366"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		luhnCheck(cardNo)
	}
}

// BenchmarkValidateBankCard_Mixed 混合场景（更接近实际使用）
func BenchmarkValidateBankCard_Mixed(b *testing.B) {
	cards := []string{
		"4532015112830366", // 有效16位
		"378282246310005",  // 有效15位
		"6011111111111117", // 有效16位
		"123",              // 无效：太短
		"abcd1234567890",   // 无效：非数字
		"",                 // 无效：空
	}

	fls := make([]FieldLevel, len(cards))
	for i, card := range cards {
		fls[i] = newMockFieldLevel(card)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCard(fls[i%len(fls)])
	}
}

func BenchmarkIDCard18_Optimized_Valid(b *testing.B) {
	testCard := "110101199003072273"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18(testCard)
	}
}

func BenchmarkIDCard18_Optimized_Invalid(b *testing.B) {
	testCard := "invalid"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18(testCard)
	}
}

func BenchmarkIDCard18_Optimized_Alloc(b *testing.B) {
	testCard := "110101199003072273"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateIDCard18(testCard)
	}
}
