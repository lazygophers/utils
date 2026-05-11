package validator

import (
	"fmt"
	"testing"
)

// TestBankCardOptimizations 对比所有优化方案
func TestBankCardOptimizations(t *testing.T) {
	card := "4532015112830366" // 有效的16位银行卡号（Visa测试卡号）

	// 首先验证所有方案的正确性
	functions := map[string]func(string) bool{
		"Current":          validateBankCardCurrent,
		"Opt1-字节手动Luhn": validateBankCardOpt1,
		"Opt2-查找表":       validateBankCardOpt2,
		"Opt3-预计算双倍":    validateBankCardOpt3,
		"Opt4-快速失败":     validateBankCardOpt4,
		"Opt5-索引循环":     validateBankCardOpt5,
		"Opt6-ASCII优化":   validateBankCardOpt6,
		"Opt7-单次遍历":     validateBankCardOpt7,
		"Opt8-位运算":       validateBankCardOpt8,
		"Opt9-反向遍历":     validateBankCardOpt9,
		"Opt10-组合优化":    validateBankCardOpt10,
		"Opt11-无分支":      validateBankCardOpt11,
		"Opt12-SIMD启发":   validateBankCardOpt12,
	}

	// 验证正确性
	for name, fn := range functions {
		if !fn(card) {
			t.Fatalf("[%s] 正确性验证失败: 有效卡号被拒绝", name)
		}
		if fn("abcd1234567890") {
			t.Fatalf("[%s] 正确性验证失败: 无效卡号被接受", name)
		}
		if fn("123") {
			t.Fatalf("[%s] 正确性验证失败: 过短卡号被接受", name)
		}
	}

	// 基准测试当前实现
	baseline := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateBankCardCurrent(card)
		}
	})

	fmt.Printf("\n========== 银行卡验证性能优化结果 ==========\n")
	fmt.Printf("测试数据: %s (16位有效卡号)\n", card)
	fmt.Printf("\n当前实现 (Baseline):\n")
	fmt.Printf("  %s\n", baseline)
	fmt.Printf("  %.2f ns/op, %.2f MB/s\n\n", float64(baseline.NsPerOp()), 1000.0/float64(baseline.NsPerOp())*16)

	// 测试所有优化方案
	fmt.Printf("优化方案对比:\n")
	fmt.Printf("%-20s %12s %12s %8s\n", "方案", "ns/op", "相对性能", "提升%")
	fmt.Printf("%s\n", "--------------------------------------------------------------------------------")

	for name, fn := range functions {
		if name == "Current" {
			continue
		}

		result := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(card)
			}
		})

		speedup := float64(baseline.NsPerOp()) / float64(result.NsPerOp())
		improvement := (1 - float64(result.NsPerOp())/float64(baseline.NsPerOp())) * 100

		fmt.Printf("%-20s %12.2f %11.2fx %7.1f%%\n",
			name,
			float64(result.NsPerOp()),
			speedup,
			improvement,
		)
	}

	// 找出最优方案
	fmt.Printf("\n========== 性能分析 ==========\n")

	bestName := "Current"
	bestNs := baseline.NsPerOp()

	for name, fn := range functions {
		if name == "Current" {
			continue
		}
		result := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(card)
			}
		})
		if result.NsPerOp() < bestNs {
			bestNs = result.NsPerOp()
			bestName = name
		}
	}

	bestSpeedup := float64(baseline.NsPerOp()) / float64(bestNs)
	fmt.Printf("最优方案: %s\n", bestName)
	fmt.Printf("性能提升: %.2fx (%.1f%%)\n", bestSpeedup, (1-1.0/bestSpeedup)*100)

	// 内存分配对比
	fmt.Printf("\n========== 内存分配对比 ==========\n")

	allocBaseline := testing.AllocsPerRun(1000, func() {
		validateBankCardCurrent(card)
	})
	fmt.Printf("当前实现: %.0f allocs/op\n\n", allocBaseline)

	for name, fn := range functions {
		if name == "Current" {
			continue
		}
		allocs := testing.AllocsPerRun(1000, func() {
			fn(card)
		})
		fmt.Printf("%-20s %.0f allocs/op", name, allocs)
		if allocs < allocBaseline {
			fmt.Printf(" (减少 %.0f)\n", allocBaseline-allocs)
		} else if allocs > allocBaseline {
			fmt.Printf(" (增加 %.0f)\n", allocs-allocBaseline)
		} else {
			fmt.Printf(" (相同)\n")
		}
	}
}

// ========== 所有实现函数 ==========

// validateBankCardCurrent 当前实现
func validateBankCardCurrent(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	if len(cardNo) < 13 || len(cardNo) > 19 {
		return false
	}
	for _, r := range cardNo {
		if r < '0' || r > '9' {
			return false
		}
	}
	return luhnCheckCurrent(cardNo)
}

func luhnCheckCurrent(cardNo string) bool {
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

// validateBankCardOpt1 字节级 + 手动Luhn
func validateBankCardOpt1(cardNo string) bool {
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

// validateBankCardOpt2 查找表
func validateBankCardOpt2(cardNo string) bool {
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

// validateBankCardOpt3 预计算双倍值
func validateBankCardOpt3(cardNo string) bool {
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

// validateBankCardOpt4 快速失败
func validateBankCardOpt4(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	if len(cardNo) == 0 {
		return false
	}
	firstChar := cardNo[0]
	if firstChar < '0' || firstChar > '9' {
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

// validateBankCardOpt5 索引循环
func validateBankCardOpt5(cardNo string) bool {
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

// validateBankCardOpt6 ASCII优化
func validateBankCardOpt6(cardNo string) bool {
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

// validateBankCardOpt7 单次遍历
func validateBankCardOpt7(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	sum := 0
	double := false
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

// validateBankCardOpt8 位运算
func validateBankCardOpt8(cardNo string) bool {
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
	return sum%10 == 0
}

// validateBankCardOpt9 反向遍历
func validateBankCardOpt9(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	sum := 0
	double := false
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

// validateBankCardOpt10 组合优化（最优方案）
func validateBankCardOpt10(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	if len(cardNo) == 0 {
		return false
	}
	firstChar := cardNo[0]
	if firstChar < '0' || firstChar > '9' {
		return false
	}
	sum := 0
	double := false
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

// validateBankCardOpt11 无分支
func validateBankCardOpt11(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	// 查找表：前10个是不双倍，后10个是双倍后的值
	var lut = [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
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

// validateBankCardOpt12 SIMD启发式
func validateBankCardOpt12(cardNo string) bool {
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
	double := false
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
