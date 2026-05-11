package validator

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

// 测试数据
var (
	validUUIDs = []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b811-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b812-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b814-9dad-11d1-80b4-00c04fd430c8",
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
		"01234567-89ab-cdef-0123-456789abcdef",
	}
	invalidUUIDs = []string{
		"550e8400-e29b-41d4-a716-44665544000",  // 太短
		"550e8400-e29b-41d4-a716-4466554400000", // 太长
		"550e8400-e29b-41d4-a716-44665544000G", // 无效字符
		"550e8400-e29b-41d4-a716-44665544000 ", // 尾部空格
		" 50e8400-e29b-41d4-a716-446655440000", // 头部空格
		"550e8400e29b-41d4-a716-446655440000",  // 缺少分隔符
		"550e8400-e29b-41d4-a716-44665544000",  // 格式错误
		"G50e8400-e29b-41d4-a716-446655440000", // 大写G
		"not-a-uuid",                           // 明显错误
		"",                                     // 空字符串
	}
)

// ============== 方案1：当前实现（基线） ==============
func validateUUIDOriginal(uuid string) bool {
	if uuid == "" {
		return false
	}
	return uuidRegex.MatchString(strings.ToLower(uuid))
}

// ============== 方案2：手动检查（字节级） ==============
func validateUUIDManual(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查固定位置的分隔符
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 检查每段
	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			continue // 跳过分隔符
		default:
			c := uuid[i]
			isDigit := c >= '0' && c <= '9'
			isLower := c >= 'a' && c <= 'f'
			isUpper := c >= 'A' && c <= 'F'
			if !isDigit && !isLower && !isUpper {
				return false
			}
		}
	}

	return true
}

// ============== 方案3：分段验证 ==============
func validateUUIDSegmented(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 验证第一段：8个字符
	if !isHexSegment(uuid[0:8]) {
		return false
	}

	// 验证第二段：4个字符
	if !isHexSegment(uuid[9:13]) {
		return false
	}

	// 验证第三段：4个字符
	if !isHexSegment(uuid[14:18]) {
		return false
	}

	// 验证第四段：4个字符
	if !isHexSegment(uuid[19:23]) {
		return false
	}

	// 验证第五段：12个字符
	if !isHexSegment(uuid[24:36]) {
		return false
	}

	return true
}

func isHexSegment(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		isDigit := c >= '0' && c <= '9'
		isLower := c >= 'a' && c <= 'f'
		isUpper := c >= 'A' && c <= 'F'
		if !isDigit && !isLower && !isUpper {
			return false
		}
	}
	return true
}

// ============== 方案4：使用 strings.IndexByte ==============
func validateUUIDIndexByte(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符位置
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 移除分隔符后检查
	hexPart := uuid[:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]

	for i := 0; i < len(hexPart); i++ {
		c := hexPart[i]
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') && !(c >= 'A' && c <= 'F') {
			return false
		}
	}

	return true
}

// ============== 方案5：预计算查找表 ==============
var hexTable = [256]bool{}

func init() {
	for i := 0; i < 256; i++ {
		c := byte(i)
		hexTable[i] = (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
	}
}

func validateUUIDLookupTable(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			if uuid[i] != '-' {
				return false
			}
		default:
			if !hexTable[uuid[i]] {
				return false
			}
		}
	}

	return true
}

// ============== 方案6：混合模式（快速路径 + 分段验证） ==============
func validateUUIDHybrid(uuid string) bool {
	// 快速长度检查
	if len(uuid) != 36 {
		return false
	}

	// 快速分隔符检查
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 批量检查十六进制字符
	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			continue
		default:
			c := uuid[i]
			if !(((c >= '0') && (c <= '9')) || ((c >= 'a') && (c <= 'f')) || ((c >= 'A') && (c <= 'F'))) {
				return false
			}
		}
	}

	return true
}

// ============== 方案7：使用 unicode.Is ==============
func validateUUIDUnicode(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			if uuid[i] != '-' {
				return false
			}
		default:
			c := rune(uuid[i])
			if !unicode.IsDigit(c) && !(c >= 'a' && c <= 'f') && !(c >= 'A' && c <= 'F') {
				return false
			}
		}
	}

	return true
}

// ============== 方案8：直接字节比较（无分支） ==============
func validateUUIDByteCompare(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符
	if uuid[8]|uuid[13]|uuid[18]|uuid[23] != '-' {
		return false
	}

	// 检查所有字符
	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		c := uuid[i]
		isValid := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
		if !isValid {
			return false
		}
	}

	return true
}

// ============== 方案9：使用 ASCII 边界检查 ==============
func validateUUIDASCIICheck(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		c := uuid[i]
		// ASCII 快速检查
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}

	return true
}

// ============== 方案10：预定义有效字符集 ==============
var validHexChars = map[byte]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true,
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true,
}

func validateUUIDMapCheck(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		if !validHexChars[uuid[i]] {
			return false
		}
	}

	return true
}

// ============== 方案11：位操作优化 ==============
func validateUUIDBitOps(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		c := uuid[i]
		// 使用位操作优化范围检查
		isDigit := (c-'0')&0xFF <= 9
		isLower := ((c|0x20)-'a')&0xFF <= 5
		isUpper := (c-'A')&0xFF <= 5
		if !isDigit && !isLower && !isUpper {
			return false
		}
	}

	return true
}

// ============== 方案12：使用 strings.IndexAny 检查无效字符 ==============
func validateUUIDIndexAny(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 检查是否包含无效字符
	hexPart := uuid[:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]
	if strings.IndexAny(hexPart, "0123456789abcdefABCDEF") == -1 {
		return len(hexPart) == 0
	}

	// 逐字符验证
	for _, c := range hexPart {
		if !(((c >= '0') && (c <= '9')) || ((c >= 'a') && (c <= 'f')) || ((c >= 'A') && (c <= 'F'))) {
			return false
		}
	}

	return true
}

// ============== 基准测试 ==============
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

// ============== 验证正确性的测试 ==============
func TestValidateUUID_Correctness(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(string) bool
	}{
		{"Original", validateUUIDOriginal},
		{"Manual", validateUUIDManual},
		{"Segmented", validateUUIDSegmented},
		{"IndexByte", validateUUIDIndexByte},
		{"LookupTable", validateUUIDLookupTable},
		{"Hybrid", validateUUIDHybrid},
		{"Unicode", validateUUIDUnicode},
		{"ByteCompare", validateUUIDByteCompare},
		{"ASCIICheck", validateUUIDASCIICheck},
		{"MapCheck", validateUUIDMapCheck},
		{"BitOps", validateUUIDBitOps},
		{"IndexAny", validateUUIDIndexAny},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			// 测试有效 UUID
			for _, uuid := range validUUIDs {
				if !impl.fn(uuid) {
					t.Errorf("%s: 有效 UUID 被拒绝: %s", impl.name, uuid)
				}
			}

			// 测试无效 UUID
			for _, uuid := range invalidUUIDs {
				if impl.fn(uuid) {
					t.Errorf("%s: 无效 UUID 被接受: %s", impl.name, uuid)
				}
			}
		})
	}
}

// 性能对比辅助函数
func runComparisonBenchmark() {
	fmt.Println("运行 UUID 验证性能对比...")
	fmt.Println("请在终端执行: cd validator && go test -bench=BenchmarkValidateUUID -benchmem -benchtime=3s | tee uuid_comparison_results.txt")
}
