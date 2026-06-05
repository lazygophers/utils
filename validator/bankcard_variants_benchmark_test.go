package validator

import (
	"reflect"
	"testing"
	"unicode"
)

func BenchmarkMobile01_Regex_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkMobile02_Manual_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkMobile03_Bytes_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile03(fl)
	}
}

func BenchmarkMobile04_FastPath_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile04(fl)
	}
}

func BenchmarkMobile05_Lookup_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile05(fl)
	}
}

func BenchmarkMobile06_LengthFirst_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile06(fl)
	}
}

func BenchmarkMobile09_Combined_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile09(fl)
	}
}

func BenchmarkMobile10_Range_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile10(fl)
	}
}

func BenchmarkMobile11_Unrolled_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

func BenchmarkMobile12_BitOps_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile12(fl)
	}
}

// 无效前缀测试
func BenchmarkMobile01_Regex_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("12812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkMobile02_Manual_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("12812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkMobile11_Unrolled_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("12812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

// 无效长度测试
func BenchmarkMobile01_Regex_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("138123456")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkMobile02_Manual_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("138123456")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkMobile11_Unrolled_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("138123456")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

// 专门的性能基准测试
func BenchmarkPattern_Email_Valid(b *testing.B) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("test@example.com")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_Email_Invalid(b *testing.B) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("invalid")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FixedLength_Valid(b *testing.B) {
	pattern := `^\d{5}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("12345")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_Literal_Valid(b *testing.B) {
	pattern := `^hello$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("hello")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

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
		"abcd1234567890",       // 包含字母
		"123",                  // 太短
		"12345678901234567890", // 太长
		"6222021234567891",     // Luhn失败
		"",                     // 空字符串
		"62220@1234567890",     // 特殊字符
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

type QuickStruct struct {
	Name string `validate:"required"`
}

func BenchmarkQuick(b *testing.B) {
	v, _ := New()
	s := QuickStruct{Name: "test"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(s)
	}
}

func BenchmarkDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i * 2
	}
}

func BenchmarkValidateIPv4_Original(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Original(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Original(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_ByteParse(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ByteParse(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ByteParse(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_StateMachine(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_StateMachine(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_StateMachine(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_Manual(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Manual(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Manual(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_NetParse(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_NetParse(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_NetParse(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_LookupTable(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_LookupTable(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_LookupTable(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_BitOps(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_BitOps(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_BitOps(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_PreAlloc(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_PreAlloc(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_PreAlloc(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_Hybrid(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Hybrid(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Hybrid(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_ZeroAlloc(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ZeroAlloc(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ZeroAlloc(invalidIP)
		}
	})
}

func BenchmarkValidateEmailRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailRegex(email)
		}
		for _, email := range invalidEmails {
			validateEmailRegex(email)
		}
	}
}

func BenchmarkValidateEmailSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailSplit(email)
		}
		for _, email := range invalidEmails {
			validateEmailSplit(email)
		}
	}
}

func BenchmarkValidateEmailStateMachine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailStateMachine(email)
		}
		for _, email := range invalidEmails {
			validateEmailStateMachine(email)
		}
	}
}

func BenchmarkValidateEmailLastIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailLastIndex(email)
		}
		for _, email := range invalidEmails {
			validateEmailLastIndex(email)
		}
	}
}

func BenchmarkValidateEmailFastPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailFastPath(email)
		}
		for _, email := range invalidEmails {
			validateEmailFastPath(email)
		}
	}
}

func BenchmarkValidateEmailSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailSegmented(email)
		}
		for _, email := range invalidEmails {
			validateEmailSegmented(email)
		}
	}
}

func BenchmarkValidateEmailRFCSimplified(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailRFCSimplified(email)
		}
		for _, email := range invalidEmails {
			validateEmailRFCSimplified(email)
		}
	}
}

func BenchmarkValidateEmailStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailStrings(email)
		}
		for _, email := range invalidEmails {
			validateEmailStrings(email)
		}
	}
}

func BenchmarkValidateEmailCombined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailCombined(email)
		}
		for _, email := range invalidEmails {
			validateEmailCombined(email)
		}
	}
}

func BenchmarkValidateEmailASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailASCII(email)
		}
		for _, email := range invalidEmails {
			validateEmailASCII(email)
		}
	}
}

func BenchmarkValidateEmailStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailStdLib(email)
		}
		for _, email := range invalidEmails {
			validateEmailStdLib(email)
		}
	}
}

func BenchmarkValidateEmailMinimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailMinimal(email)
		}
		for _, email := range invalidEmails {
			validateEmailMinimal(email)
		}
	}
}

func BenchmarkValidEmailsOnly(b *testing.B) {
	validators := map[string]func(string) bool{
		"方案1-正则":        validateEmailRegex,
		"方案4-LastIndex": validateEmailLastIndex,
		"方案9-组合优化":      validateEmailCombined,
		"方案10-ASCII":    validateEmailASCII,
	}

	for name, fn := range validators {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, email := range validEmails {
					fn(email)
				}
			}
		})
	}
}

func BenchmarkInvalidEmailsOnly(b *testing.B) {
	validators := map[string]func(string) bool{
		"方案1-正则":        validateEmailRegex,
		"方案4-LastIndex": validateEmailLastIndex,
		"方案9-组合优化":      validateEmailCombined,
		"方案10-ASCII":    validateEmailASCII,
	}

	for name, fn := range validators {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, email := range invalidEmails {
					fn(email)
				}
			}
		})
	}
}

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

func BenchmarkValidateUUID_Original_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDOriginal(uuid)
		}
	}
}

func BenchmarkValidateUUID_Original_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDOriginal(uuid)
		}
	}
}

func BenchmarkValidateUUID_Manual_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDManual(uuid)
		}
	}
}

func BenchmarkValidateUUID_Manual_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDManual(uuid)
		}
	}
}

func BenchmarkValidateUUID_Segmented_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDSegmented(uuid)
		}
	}
}

func BenchmarkValidateUUID_Segmented_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDSegmented(uuid)
		}
	}
}

func BenchmarkValidateUUID_LookupTable_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDLookupTable(uuid)
		}
	}
}

func BenchmarkValidateUUID_LookupTable_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDLookupTable(uuid)
		}
	}
}

func BenchmarkValidateUUID_Hybrid_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDHybrid(uuid)
		}
	}
}

func BenchmarkValidateUUID_Hybrid_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDHybrid(uuid)
		}
	}
}

func BenchmarkValidateUUID_ASCII_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDASCIICheck(uuid)
		}
	}
}

func BenchmarkValidateUUID_ASCII_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDASCIICheck(uuid)
		}
	}
}

func BenchmarkValidateUUID_ByteCompare_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDByteCompare(uuid)
		}
	}
}

func BenchmarkValidateUUID_ByteCompare_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDByteCompare(uuid)
		}
	}
}

func BenchmarkValidateUUID_BitOps_Valid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range validUUIDs {
			validateUUIDBitOps(uuid)
		}
	}
}

func BenchmarkValidateUUID_BitOps_Invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, uuid := range invalidUUIDs {
			validateUUIDBitOps(uuid)
		}
	}
}

func BenchmarkStrongPassword_Validator(b *testing.B) {
	v, _ := New()

	type Form struct {
		Password string `validate:"strong_password"`
	}

	validPasswords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}

	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validPasswords {
				form := Form{Password: pwd}
				_ = v.Struct(form)
			}
		}
	})
}

// 测试数据
var (
	validStrongPasswords = []string{
		"Abc123!@",        // 最小长度，包含所有类型
		"Password123!",    // 常见强密码
		"SecurePass#2024", // 包含数字
		"MyP@ssw0rd",      // 复杂密码
		"Test@1234",       // 简单但符合
		"ADMIN@123",       // 全大写+数字+特殊
		"student#123",     // 全小写+数字+特殊
		"User2024$Pass",   // 混合
		"1A2b3C4d!",       // 交替字符
		"P@ssw0rd123456",  // 较长密码
	}
	invalidStrongPasswords = []string{
		"",             // 空
		"short1A!",     // 太短
		"nocaps123!",   // 无大写
		"NOLOWER123!",  // 无小写
		"NoNumber!!",   // 无数字
		"NoSpecial123", // 无特殊字符
		"onlylower",    // 只小写
		"ONLYUPPER",    // 只大写
		"12345678",     // 只数字
		"!@#$%^&*",     // 只特殊字符
		"Abc123",       // 长度够但只2种类型
		"ABCDEFGH",     // 只1种类型
	}
	weakPasswords = []string{
		"Password1", // 无特殊字符（2种类型）
		"password!", // 无数字（2种类型）
		"PASSWORD1", // 无小写和特殊（2种类型）
		"12345678a", // 无大写和特殊（2种类型）
	}
)

// ============== 方案1：当前实现（基线） ==============
func validateStrongPasswordOriginal(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasNumber = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案2：使用字节循环替代 rune ==============
func validateStrongPasswordByteLoop(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		case (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~'):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案3：使用查找表优化 ASCII 范围检查 ==============
func validateStrongPasswordLookupTable(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 预定义字符类型查找表
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case (c >= 'A' && c <= 'Z'):
			hasUpper = true
		case (c >= 'a' && c <= 'z'):
			hasLower = true
		case (c >= '0' && c <= '9'):
			hasNumber = true
		default:
			// 其他 ASCII 字符视为特殊字符
			if c >= 32 && c <= 126 {
				hasSpecial = true
			}
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案4：快速失败优化 ==============
func validateStrongPasswordFastFail(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   uint8
		hasLower   uint8
		hasNumber  uint8
		hasSpecial uint8
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = 1
		case c >= 'a' && c <= 'z':
			hasLower = 1
		case c >= '0' && c <= '9':
			hasNumber = 1
		default:
			if c >= 32 && c <= 126 {
				hasSpecial = 1
			}
		}

		// 快速失败：已经找到3种类型且遍历了足够长度
		if hasUpper+hasLower+hasNumber+hasSpecial >= 3 {
			return true
		}
	}

	return hasUpper+hasLower+hasNumber+hasSpecial >= 3
}

// ============== 方案5：位掩码优化 ==============
func validateStrongPasswordBitMask(password string) bool {
	if len(password) < 8 {
		return false
	}

	const (
		upperMask   = 1 << 0
		lowerMask   = 1 << 1
		numberMask  = 1 << 2
		specialMask = 1 << 3
	)

	var mask uint8

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			mask |= upperMask
		case c >= 'a' && c <= 'z':
			mask |= lowerMask
		case c >= '0' && c <= '9':
			mask |= numberMask
		default:
			if c >= 32 && c <= 126 {
				mask |= specialMask
			}
		}
	}

	// 检查是否至少有3种类型
	return (mask&upperMask + mask&lowerMask + mask&numberMask + mask&specialMask) >= 3
}

// ============== 方案6：预计算特殊字符范围 ==============
func validateStrongPasswordPrecompute(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		case c >= 33 && c <= 47, c >= 58 && c <= 64, c >= 91 && c <= 96, c >= 123 && c <= 126:
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案7：分支消除优化 ==============
func validateStrongPasswordBranchless(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   uint8
		hasLower   uint8
		hasNumber  uint8
		hasSpecial uint8
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		upper := uint8(0)
		if c >= 'A' && c <= 'Z' {
			upper = 1
		}
		lower := uint8(0)
		if c >= 'a' && c <= 'z' {
			lower = 1
		}
		digit := uint8(0)
		if c >= '0' && c <= '9' {
			digit = 1
		}
		special := uint8(0)
		if c >= 33 && c <= 126 {
			if c < '0' || c > '9' {
				if c < 'A' || c > 'Z' {
					if c < 'a' || c > 'z' {
						special = 1
					}
				}
			}
		}

		hasUpper += upper
		hasLower += lower
		hasNumber += digit
		hasSpecial += special
	}

	count := uint8(0)
	if hasUpper > 0 {
		count++
	}
	if hasLower > 0 {
		count++
	}
	if hasNumber > 0 {
		count++
	}
	if hasSpecial > 0 {
		count++
	}

	return count >= 3
}

// ============== 方案8：ASCII 表查找 ==============
func validateStrongPasswordASCIITable(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 字符类型表：0=其他, 1=大写, 2=小写, 3=数字, 4=特殊
	var charType [128]uint8
	for c := 'A'; c <= 'Z'; c++ {
		charType[c] = 1
	}
	for c := 'a'; c <= 'z'; c++ {
		charType[c] = 2
	}
	for c := '0'; c <= '9'; c++ {
		charType[c] = 3
	}
	// 特殊字符
	for c := '!'; c <= '/'; c++ {
		charType[c] = 4
	}
	for c := ':'; c <= '@'; c++ {
		charType[c] = 4
	}
	for c := '['; c <= '`'; c++ {
		charType[c] = 4
	}
	for c := '{'; c <= '~'; c++ {
		charType[c] = 4
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		if c < 128 {
			switch charType[c] {
			case 1:
				hasUpper = true
			case 2:
				hasLower = true
			case 3:
				hasNumber = true
			case 4:
				hasSpecial = true
			}
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案9：混合 Unicode + ASCII 快速路径 ==============
func validateStrongPasswordHybrid(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 先检查是否为纯 ASCII
	isASCII := true
	for i := 0; i < len(password); i++ {
		if password[i] > 127 {
			isASCII = false
			break
		}
	}

	if isASCII {
		// 快速 ASCII 路径
		var (
			hasUpper   bool
			hasLower   bool
			hasNumber  bool
			hasSpecial bool
		)

		for i := 0; i < len(password); i++ {
			c := password[i]
			switch {
			case c >= 'A' && c <= 'Z':
				hasUpper = true
			case c >= 'a' && c <= 'z':
				hasLower = true
			case c >= '0' && c <= '9':
				hasNumber = true
			default:
				if c >= 32 && c <= 126 {
					hasSpecial = true
				}
			}
		}

		count := 0
		if hasUpper {
			count++
		}
		if hasLower {
			count++
		}
		if hasNumber {
			count++
		}
		if hasSpecial {
			count++
		}

		return count >= 3
	}

	// Unicode 路径（原始实现）
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案10：计数器累加优化 ==============
func validateStrongPasswordCounter(password string) bool {
	if len(password) < 8 {
		return false
	}

	count := 0

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			if !hasUpper {
				hasUpper = true
				count++
			}
		case c >= 'a' && c <= 'z':
			if !hasLower {
				hasLower = true
				count++
			}
		case c >= '0' && c <= '9':
			if !hasNumber {
				hasNumber = true
				count++
			}
		default:
			if c >= 32 && c <= 126 && !hasSpecial {
				hasSpecial = true
				count++
			}
		}

		// 快速失败
		if count >= 3 {
			return true
		}
	}

	return count >= 3
}

// ============== 方案11：SIMD 风格批量处理 ==============
func validateStrongPasswordSIMDStyle(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	i := 0
	// 批量处理 4 个字符
	for i+4 <= len(password) {
		c1, c2, c3, c4 := password[i], password[i+1], password[i+2], password[i+3]

		// 字符 1
		switch {
		case c1 >= 'A' && c1 <= 'Z':
			hasUpper = true
		case c1 >= 'a' && c1 <= 'z':
			hasLower = true
		case c1 >= '0' && c1 <= '9':
			hasNumber = true
		default:
			if c1 >= 32 && c1 <= 126 {
				hasSpecial = true
			}
		}

		// 字符 2
		switch {
		case c2 >= 'A' && c2 <= 'Z':
			hasUpper = true
		case c2 >= 'a' && c2 <= 'z':
			hasLower = true
		case c2 >= '0' && c2 <= '9':
			hasNumber = true
		default:
			if c2 >= 32 && c2 <= 126 {
				hasSpecial = true
			}
		}

		// 字符 3
		switch {
		case c3 >= 'A' && c3 <= 'Z':
			hasUpper = true
		case c3 >= 'a' && c3 <= 'z':
			hasLower = true
		case c3 >= '0' && c3 <= '9':
			hasNumber = true
		default:
			if c3 >= 32 && c3 <= 126 {
				hasSpecial = true
			}
		}

		// 字符 4
		switch {
		case c4 >= 'A' && c4 <= 'Z':
			hasUpper = true
		case c4 >= 'a' && c4 <= 'z':
			hasLower = true
		case c4 >= '0' && c4 <= '9':
			hasNumber = true
		default:
			if c4 >= 32 && c4 <= 126 {
				hasSpecial = true
			}
		}

		i += 4
	}

	// 处理剩余字符
	for i < len(password) {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		default:
			if c >= 32 && c <= 126 {
				hasSpecial = true
			}
		}
		i++
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案12：内联优化 ==============
func validateStrongPasswordInlined(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for i := 0; i < len(password); i++ {
		c := password[i]
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		} else if c >= 'a' && c <= 'z' {
			hasLower = true
		} else if c >= '0' && c <= '9' {
			hasNumber = true
		} else if c >= 32 && c <= 126 {
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 基准测试 ==============

// BenchmarkStrongPassword_Original 当前实现（基线）
func BenchmarkStrongPassword_Original(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordOriginal(pwd)
			}
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range invalidStrongPasswords {
				_ = validateStrongPasswordOriginal(pwd)
			}
		}
	})

	b.Run("Weak", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range weakPasswords {
				_ = validateStrongPasswordOriginal(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_ByteLoop 字节循环
func BenchmarkStrongPassword_ByteLoop(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordByteLoop(pwd)
			}
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range invalidStrongPasswords {
				_ = validateStrongPasswordByteLoop(pwd)
			}
		}
	})

	b.Run("Weak", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range weakPasswords {
				_ = validateStrongPasswordByteLoop(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_LookupTable 查找表优化
func BenchmarkStrongPassword_LookupTable(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordLookupTable(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_FastFail 快速失败
func BenchmarkStrongPassword_FastFail(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordFastFail(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_BitMask 位掩码
func BenchmarkStrongPassword_BitMask(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordBitMask(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Precompute 预计算
func BenchmarkStrongPassword_Precompute(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordPrecompute(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Branchless 分支消除
func BenchmarkStrongPassword_Branchless(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordBranchless(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_ASCIITable ASCII 表
func BenchmarkStrongPassword_ASCIITable(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordASCIITable(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Hybrid 混合
func BenchmarkStrongPassword_Hybrid(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordHybrid(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Counter 计数器
func BenchmarkStrongPassword_Counter(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordCounter(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_SIMDStyle SIMD 风格
func BenchmarkStrongPassword_SIMDStyle(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordSIMDStyle(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Inlined 内联优化
func BenchmarkStrongPassword_Inlined(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordInlined(pwd)
			}
		}
	})
}

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
