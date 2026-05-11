//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// 测试数据
var (
	validMobile   = "13812345678"
	invalidPrefix = "12812345678"
	invalidLen    = "138123456"
)

// ========== 方案1: 当前正则（基线） ==========
var mobileRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

func validateMobile01(mobile string) bool {
	if mobile == "" {
		return false
	}
	return mobileRegex.MatchString(mobile)
}

// ========== 方案2: 手动逐字符 ASCII 比较 ==========
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

// ========== 方案3: 字节切片遍历 ==========
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

// ========== 方案4: 前缀快速路径 ==========
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

// ========== 方案5: 查找表 ==========
var validSecondDigits = [256]bool{
	'3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
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

// ========== 方案9: 组合优化 ==========
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

// ========== 方案11: 循环展开 ==========
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

// ========== 方案12: 位运算优化 ==========
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

func benchmark(name string, fn func(string) bool, input string, iterations int) time.Duration {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn(input)
	}
	return time.Since(start)
}

func main() {
	iterations := 1000000

	functions := map[string]func(string) bool{
		"01_Regex基线":  validateMobile01,
		"02_手动逐字符":  validateMobile02,
		"03_字节切片":   validateMobile03,
		"04_快速路径":   validateMobile04,
		"05_查找表":    validateMobile05,
		"09_组合优化":   validateMobile09,
		"11_循环展开":   validateMobile11,
		"12_位运算优化":  validateMobile12,
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
			fmt.Printf("✗ %s: 失败\n", name)
		}
	}

	fmt.Println("\n=== 性能基准测试 ===")
	fmt.Printf("每个函数运行 %d 次\n\n", iterations)

	fmt.Printf("%-15s %15s %15s %15s\n", "方案", "有效(μs)", "无效前缀(μs)", "无效长度(μs)")
	fmt.Println(strings.Repeat("-", 65))

	results := make(map[string][3]time.Duration)

	for name, fn := range functions {
		validTime := benchmark(name, fn, validMobile, iterations)
		invalidPrefixTime := benchmark(name, fn, invalidPrefix, iterations)
		invalidLenTime := benchmark(name, fn, invalidLen, iterations)

		results[name] = [3]time.Duration{validTime, invalidPrefixTime, invalidLenTime}

		fmt.Printf("%-15s %15d %15d %15d\n",
			name,
			validTime.Microseconds(),
			invalidPrefixTime.Microseconds(),
			invalidLenTime.Microseconds())
	}

	// 找出最优方案
	fmt.Println("\n=== 最优方案分析 ===")

	bestValidName := ""
	var bestValidTime time.Duration
	for name, times := range results {
		if bestValidName == "" || times[0] < bestValidTime {
			bestValidName = name
			bestValidTime = times[0]
		}
	}

	fmt.Printf("有效手机号最快: %s (%d μs for %d iterations)\n",
		bestValidName, bestValidTime.Microseconds(), iterations)

	// 计算性能提升
	baseLineTime := results["01_Regex基线"][0].Microseconds()
	bestTime := results[bestValidName][0].Microseconds()
	improvement := float64(baseLineTime) / float64(bestTime)

	fmt.Printf("性能提升: %.2fx faster than baseline\n", improvement)
}
