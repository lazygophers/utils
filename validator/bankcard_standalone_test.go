package validator

import (
	"testing"
)

// 独立的银行卡基准测试 - 不依赖包内其他文件

// 生成测试银行卡号
func genBankCardsStandalone(n int) []string {
	cards := make([]string, n)
	validCards := []string{
		"6222021234567890",
		"6228481234567890123",
		"4000001234567890",
		"5555551234567890",
		"378282246310005",
		"6011111111111117",
		"3530111333300000",
		"1234567890123456789",
	}

	for i := 0; i < n; i++ {
		cards[i] = validCards[i%len(validCards)]
	}
	return cards
}

// 生成无效银行卡号
func genInvalidBankCardsStandalone(n int) []string {
	cards := make([]string, n)
	invalidCases := []string{
		"abcd1234567890",
		"123",
		"12345678901234567890",
		"6222021234567891",
		"",
		"62220@1234567890",
	}

	for i := 0; i < n; i++ {
		cards[i] = invalidCases[i%len(invalidCases)]
	}
	return cards
}

// 当前实现 (Baseline)
func BenchmarkBankCard_Current_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Current_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Current_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardStandalone(cards[i%len(cards)])
	}
}

// 方案1: 字节级 + 手动Luhn
func BenchmarkBankCard_Opt1_ByteManualLuhn_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt1ByteManualLuhnStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt1_ByteManualLuhn_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt1ByteManualLuhnStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt1_ByteManualLuhn_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt1ByteManualLuhnStandalone(cards[i%len(cards)])
	}
}

// 方案2: 查找表
func BenchmarkBankCard_Opt2_LookupTable_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt2LookupTableStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt2_LookupTable_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt2LookupTableStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt2_LookupTable_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt2LookupTableStandalone(cards[i%len(cards)])
	}
}

// 方案3: 预计算双倍值
func BenchmarkBankCard_Opt3_PrecomputedDoubles_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt3PrecomputedDoublesStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt3_PrecomputedDoubles_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt3PrecomputedDoublesStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt3_PrecomputedDoubles_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt3PrecomputedDoublesStandalone(cards[i%len(cards)])
	}
}

// 方案4: 快速失败
func BenchmarkBankCard_Opt4_FastFail_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt4FastFailStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt4_FastFail_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt4FastFailStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt4_FastFail_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt4FastFailStandalone(cards[i%len(cards)])
	}
}

// 方案5: 索引循环
func BenchmarkBankCard_Opt5_IndexLoop_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt5IndexLoopStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt5_IndexLoop_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt5IndexLoopStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt5_IndexLoop_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt5IndexLoopStandalone(cards[i%len(cards)])
	}
}

// 方案6: ASCII优化
func BenchmarkBankCard_Opt6_ASCII_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt6ASCIIStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt6_ASCII_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt6ASCIIStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt6_ASCII_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt6ASCIIStandalone(cards[i%len(cards)])
	}
}

// 方案7: 单次遍历
func BenchmarkBankCard_Opt7_SinglePass_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt7SinglePassStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt7_SinglePass_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt7SinglePassStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt7_SinglePass_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt7SinglePassStandalone(cards[i%len(cards)])
	}
}

// 方案8: 位运算
func BenchmarkBankCard_Opt8_BitOps_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt8BitOpsStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt8_BitOps_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt8BitOpsStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt8_BitOps_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt8BitOpsStandalone(cards[i%len(cards)])
	}
}

// 方案9: 反向遍历
func BenchmarkBankCard_Opt9_Reverse_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt9ReverseStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt9_Reverse_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt9ReverseStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt9_Reverse_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt9ReverseStandalone(cards[i%len(cards)])
	}
}

// 方案10: 组合优化
func BenchmarkBankCard_Opt10_Combined_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10CombinedStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt10_Combined_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10CombinedStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt10_Combined_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10CombinedStandalone(cards[i%len(cards)])
	}
}

// 方案11: 无分支
func BenchmarkBankCard_Opt11_Branchless_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt11BranchlessStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt11_Branchless_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt11BranchlessStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt11_Branchless_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt11BranchlessStandalone(cards[i%len(cards)])
	}
}

// 方案12: SIMD启发式
func BenchmarkBankCard_Opt12_SimdInspired_Valid_Small(b *testing.B) {
	cards := genBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt12SimdInspiredStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt12_SimdInspired_Valid_Medium(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt12SimdInspiredStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt12_SimdInspired_Invalid_Small(b *testing.B) {
	cards := genInvalidBankCardsStandalone(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt12SimdInspiredStandalone(cards[i%len(cards)])
	}
}

// 内存分配对比
func BenchmarkBankCard_Current_Allocs(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCardStandalone(cards[i%len(cards)])
	}
}

func BenchmarkBankCard_Opt10_Combined_Allocs(b *testing.B) {
	cards := genBankCardsStandalone(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateBankCardOpt10CombinedStandalone(cards[i%len(cards)])
	}
}

// ========== 独立实现函数 ==========

// 当前实现
func validateBankCardStandalone(cardNo string) bool {
	if cardNo == "" {
		return false
	}

	if len(cardNo) < 13 || len(cardNo) > 19 {
		return false
	}

	// 检查数字
	for _, r := range cardNo {
		if r < '0' || r > '9' {
			return false
		}
	}

	return luhnCheckStandalone(cardNo)
}

func luhnCheckStandalone(cardNo string) bool {
	sum := 0
	alternate := false

	for i := len(cardNo) - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}
		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}

// 优化方案1-12的独立实现（与bankcard_variants.go相同，但独立）
func validateBankCardOpt1ByteManualLuhnStandalone(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

func validateBankCardOpt2LookupTableStandalone(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	var doubled = [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit = doubled[digit]
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

func validateBankCardOpt3PrecomputedDoublesStandalone(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	sum := 0
	double := false
	for i := l - 1; i >= 0; i-- {
		d := int(cardNo[i] - '0')
		if double {
			d = d*2 - 9*(d/5)
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

func validateBankCardOpt4FastFailStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	if cardNo[0] < '1' || cardNo[0] > '9' {
		return false
	}

	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		digit := int(c - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

func validateBankCardOpt5IndexLoopStandalone(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit <<= 1
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

func validateBankCardOpt6ASCIIStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	sum := 0
	odd := (l & 1) == 0
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		digit := int(c - '0')
		if odd {
			digit <<= 1
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		odd = !odd
	}
	return sum%10 == 0
}

func validateBankCardOpt7SinglePassStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	sum := 0
	double := true
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d += d
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

func validateBankCardOpt8BitOpsStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	sum := 0
	alt := false
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if alt {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		alt = !alt
	}
	return (sum & 0xF) == 0
}

func validateBankCardOpt9ReverseStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	sum := 0
	double := true
	i := l - 1
	for i >= 0 {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
		i--
	}
	return sum%10 == 0
}

func validateBankCardOpt10CombinedStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	firstChar := cardNo[0]
	if firstChar < '1' || firstChar > '9' {
		return false
	}

	sum := 0
	double := true
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

func validateBankCardOpt11BranchlessStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	var lut = [20]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	sum := 0
	double := 0
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		sum += lut[d+double*10]
		double ^= 1
	}
	return sum%10 == 0
}

func validateBankCardOpt12SimdInspiredStandalone(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	i := 0
	for ; i+4 <= l; i += 4 {
		c0, c1, c2, c3 := cardNo[i], cardNo[i+1], cardNo[i+2], cardNo[i+3]
		if c0 < '0' || c0 > '9' ||
			c1 < '0' || c1 > '9' ||
			c2 < '0' || c2 > '9' ||
			c3 < '0' || c3 > '9' {
			return false
		}
	}

	for ; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	sum := 0
	double := true
	for i := l - 1; i >= 0; i-- {
		d := int(cardNo[i] - '0')
		if double {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}
