package validator

import (
	"math/rand"
	"testing"
)

// 生成测试用的身份证号（固定种子保证可重复）
func genIDCards(n int) []string {
	r := rand.New(rand.NewSource(42))
	cards := make([]string, n)

	// 有效的18位身份证号（带正确校验码）
	validCards := []string{
		"110101199003072273", // 北京
		"310104199010017834", // 上海
		"44030819910403921X", // 广东（大X）
		"44030819910403921x", // 广东（小x）
		"42010619800101001X", // 湖北
		"500101198501011234", // 重庆
	}

	for i := 0; i < n; i++ {
		cards[i] = validCards[r.Intn(len(validCards))]
	}
	return cards
}

// 生成无效身份证号
func genInvalidIDCards(n int) []string {
	r := rand.New(rand.NewSource(42))
	cards := make([]string, n)

	invalidPatterns := []string{
		"12345678901234567",   // 17位
		"1234567890123456789", // 19位
		"abcdefghijklmnopqr",  // 全字母
		"11010119900307227",   // 错误校验码
		"110101199003072274",  // 错误校验码
		"11010119900307227Y",  // 非法字符
		"",                    // 空字符串
		"11010119900307227 ",  // 包含空格
		" 110101199003072273", // 前导空格
		"110101-199003072273", // 包含连字符
		"110101 199003072273", // 包含空格
		"110101A01003072273",  // 包含字母
	}

	for i := 0; i < n; i++ {
		cards[i] = invalidPatterns[r.Intn(len(invalidPatterns))]
	}
	return cards
}

// ========== 当前实现（基线） ==========

// Baseline: 当前正则实现
func Benchmark_IDCard18_Current_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Current(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Current_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Current(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Current_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Current(data[i%len(data)])
	}
}

// ========== 优化方案1：纯字节检查（最快） ==========

func Benchmark_IDCard18_Opt1_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt1(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt1_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt1(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt1_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt1(data[i%len(data)])
	}
}

// ========== 优化方案2：ASCII 快速路径 ==========

func Benchmark_IDCard18_Opt2_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt2(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt2_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt2(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt2_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt2(data[i%len(data)])
	}
}

// ========== 优化方案3：提前返回优化 ==========

func Benchmark_IDCard18_Opt3_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt3(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt3_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt3(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt3_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt3(data[i%len(data)])
	}
}

// ========== 优化方案4：查表法 ==========

func Benchmark_IDCard18_Opt4_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt4(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt4_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt4(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt4_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt4(data[i%len(data)])
	}
}

// ========== 优化方案5：混合策略 ==========

func Benchmark_IDCard18_Opt5_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt5(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt5_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt5(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt5_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt5(data[i%len(data)])
	}
}

// ========== 优化方案6：完全展开循环 ==========

func Benchmark_IDCard18_Opt6_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt6(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt6_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt6(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt6_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt6(data[i%len(data)])
	}
}

// ========== 优化方案7：SIMD 风格批量检查 ==========

func Benchmark_IDCard18_Opt7_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt7(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt7_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt7(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt7_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt7(data[i%len(data)])
	}
}

// ========== 优化方案8：双重检查锁定模式 ==========

func Benchmark_IDCard18_Opt8_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt8(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt8_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt8(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt8_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt8(data[i%len(data)])
	}
}

// ========== 优化方案9：边界内联 ==========

func Benchmark_IDCard18_Opt9_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt9(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt9_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt9(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt9_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt9(data[i%len(data)])
	}
}

// ========== 优化方案10：最小分支 ==========

func Benchmark_IDCard18_Opt10_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt10(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt10_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt10(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt10_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt10(data[i%len(data)])
	}
}

// ========== 带校验码验证的方案 ==========

// 方案11：包含完整校验码验证
func Benchmark_IDCard18_Opt11_WithChecksum_Valid_Small(b *testing.B) {
	data := genIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt11_WithChecksum(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt11_WithChecksum_Valid_Medium(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt11_WithChecksum(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt11_WithChecksum_Invalid_Small(b *testing.B) {
	data := genInvalidIDCards(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt11_WithChecksum(data[i%len(data)])
	}
}

// ========== 内存分配报告 ==========

func Benchmark_IDCard18_Current_Alloc(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Current(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt1_Alloc(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt1(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt2_Alloc(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt2(data[i%len(data)])
	}
}

func Benchmark_IDCard18_Opt11_Alloc(b *testing.B) {
	data := genIDCards(100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateIDCard18_Opt11_WithChecksum(data[i%len(data)])
	}
}
