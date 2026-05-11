package validator

import "regexp"

// ========== 当前实现（基线） ==========

// validateIDCard18_Current 当前实现：使用正则表达式
func validateIDCard18_Current(idcard string) bool {
	// 本地正则，避免依赖外部变量
	idcard18Regex := regexp.MustCompile(`^\d{17}[\dXx]$`)
	if !idcard18Regex.MatchString(idcard) {
		return false
	}
	// 当前只做格式验证，不做校验码验证
	return true
}

// ========== 优化方案1：纯字节检查 ==========

// validateIDCard18_Opt1 纯字节检查：零内存分配，最快路径
func validateIDCard18_Opt1(idcard string) bool {
	// 快速失败：长度检查
	if len(idcard) != 18 {
		return false
	}

	// 前17位必须是数字
	for i := 0; i < 17; i++ {
		c := idcard[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 最后一位：数字或X/x
	last := idcard[17]
	isDigit := last >= '0' && last <= '9'
	isX := last == 'X' || last == 'x'
	if !isDigit && !isX {
		return false
	}

	return true
}

// ========== 优化方案2：ASCII 快速路径 ==========

// validateIDCard18_Opt2 ASCII快速路径：利用ASCII连续性
func validateIDCard18_Opt2(idcard string) bool {
	if len(idcard) != 18 {
		return false
	}

	// 使用范围检查，利用CPU分支预测
	for i := 0; i < 17; i++ {
		c := idcard[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 最后一位特殊处理
	c := idcard[17]
	if (c >= '0' && c <= '9') || c == 'X' || c == 'x' {
		return true
	}
	return false
}

// ========== 优化方案3：提前返回优化 ==========

// validateIDCard18_Opt3 提前返回：尽早失败
func validateIDCard18_Opt3(idcard string) bool {
	// 多级快速失败
	l := len(idcard)
	if l != 18 {
		return false
	}

	// 检查第一位（应该是1-9，不能是0）
	if idcard[0] < '1' || idcard[0] > '9' {
		return false
	}

	// 批量检查中间16位（第2-17位）
	for i := 1; i < 17; i++ {
		c := idcard[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 最后一位
	last := idcard[17]
	if (last >= '0' && last <= '9') || last == 'X' || last == 'x' {
		return true
	}
	return false
}

// ========== 优化方案4：查表法 ==========

// validateIDCard18_Opt4 查表法：预计算数字查找表
func validateIDCard18_Opt4(idcard string) bool {
	if len(idcard) != 18 {
		return false
	}

	// 数字查找表：true表示是数字
	var digitTable = [256]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true, '8': true, '9': true,
	}

	// 前17位
	for i := 0; i < 17; i++ {
		if !digitTable[idcard[i]] {
			return false
		}
	}

	// 最后一位
	last := idcard[17]
	if digitTable[last] || last == 'X' || last == 'x' {
		return true
	}
	return false
}

// ========== 优化方案5：混合策略 ==========

// validateIDCard18_Opt5 混合策略：结合快速路径和详细检查
func validateIDCard18_Opt5(idcard string) bool {
	l := len(idcard)

	// 快速路径：常见长度错误
	if l != 18 {
		return false
	}

	// 快速路径：首字符检查（身份证不能以0开头）
	if idcard[0] == '0' {
		return false
	}

	// 批量数字检查
	for i := 0; i < 17; i++ {
		c := idcard[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 校验位检查
	last := idcard[17]
	if last >= '0' && last <= '9' {
		return true
	}
	if last == 'X' || last == 'x' {
		return true
	}
	return false
}

// ========== 优化方案6：完全展开循环 ==========

// validateIDCard18_Opt6 完全展开：手动展开循环减少迭代开销
func validateIDCard18_Opt6(idcard string) bool {
	if len(idcard) != 18 {
		return false
	}

	// 完全展开17位检查
	if idcard[0] < '0' || idcard[0] > '9' {
		return false
	}
	if idcard[1] < '0' || idcard[1] > '9' {
		return false
	}
	if idcard[2] < '0' || idcard[2] > '9' {
		return false
	}
	if idcard[3] < '0' || idcard[3] > '9' {
		return false
	}
	if idcard[4] < '0' || idcard[4] > '9' {
		return false
	}
	if idcard[5] < '0' || idcard[5] > '9' {
		return false
	}
	if idcard[6] < '0' || idcard[6] > '9' {
		return false
	}
	if idcard[7] < '0' || idcard[7] > '9' {
		return false
	}
	if idcard[8] < '0' || idcard[8] > '9' {
		return false
	}
	if idcard[9] < '0' || idcard[9] > '9' {
		return false
	}
	if idcard[10] < '0' || idcard[10] > '9' {
		return false
	}
	if idcard[11] < '0' || idcard[11] > '9' {
		return false
	}
	if idcard[12] < '0' || idcard[12] > '9' {
		return false
	}
	if idcard[13] < '0' || idcard[13] > '9' {
		return false
	}
	if idcard[14] < '0' || idcard[14] > '9' {
		return false
	}
	if idcard[15] < '0' || idcard[15] > '9' {
		return false
	}
	if idcard[16] < '0' || idcard[16] > '9' {
		return false
	}

	// 最后一位
	last := idcard[17]
	if (last >= '0' && last <= '9') || last == 'X' || last == 'x' {
		return true
	}
	return false
}

// ========== 优化方案7：SIMD 风格批量检查 ==========

// validateIDCard18_Opt7 SIMD风格：批量处理8字节
func validateIDCard18_Opt7(idcard string) bool {
	if len(idcard) != 18 {
		return false
	}

	// 前16位分两批检查（每批8位）
	for batch := 0; batch < 2; batch++ {
		base := batch * 8
		for i := 0; i < 8; i++ {
			c := idcard[base+i]
			if c < '0' || c > '9' {
				return false
			}
		}
	}

	// 第17位
	if idcard[16] < '0' || idcard[16] > '9' {
		return false
	}

	// 最后一位
	last := idcard[17]
	if (last >= '0' && last <= '9') || last == 'X' || last == 'x' {
		return true
	}
	return false
}

// ========== 优化方案8：双重检查锁定模式 ==========

// validateIDCard18_Opt8 双重检查：先粗后细
func validateIDCard18_Opt8(idcard string) bool {
	l := len(idcard)
	if l != 18 {
		return false
	}

	// 第一轮：粗略检查（ASCII范围）
	for i := 0; i < 18; i++ {
		c := idcard[i]
		// 所有字符都在ASCII可打印范围内
		if c < 0x20 || c > 0x7E {
			return false
		}
	}

	// 第二轮：精确检查
	for i := 0; i < 17; i++ {
		if idcard[i] < '0' || idcard[i] > '9' {
			return false
		}
	}

	last := idcard[17]
	if (last >= '0' && last <= '9') || last == 'X' || last == 'x' {
		return true
	}
	return false
}

// ========== 优化方案9：边界内联 ==========

// validateIDCard18_Opt9 边界内联：内联边界检查
func validateIDCard18_Opt9(idcard string) bool {
	// 内联长度检查
	if idcard == "" || len(idcard) != 18 {
		return false
	}

	// 使用连续内存访问模式
	c0 := idcard[0]
	if c0 < '0' || c0 > '9' {
		return false
	}

	c1 := idcard[1]
	if c1 < '0' || c1 > '9' {
		return false
	}

	// 批量检查中间
	for i := 2; i < 17; i++ {
		c := idcard[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 最后位
	last := idcard[17]
	isDigit := last >= '0' && last <= '9'
	isX := last == 'X' || last == 'x'
	return isDigit || isX
}

// ========== 优化方案10：最小分支 ==========

// validateIDCard18_Opt10 最小分支：减少条件分支
func validateIDCard18_Opt10(idcard string) bool {
	if len(idcard) != 18 {
		return false
	}

	// 使用位运算减少分支
	valid := true

	// 前17位
	for i := 0; i < 17; i++ {
		c := idcard[i]
		valid = valid && (c >= '0' && c <= '9')
	}

	// 最后位
	last := idcard[17]
	validLast := (last >= '0' && last <= '9') || last == 'X' || last == 'x'

	return valid && validLast
}

// ========== 方案11：包含完整校验码验证 ==========

// validateIDCard18_Opt11_WithChecksum 包含完整校验码验证
// 性能权衡：准确性与速度
func validateIDCard18_Opt11_WithChecksum(idcard string) bool {
	// 先做格式验证（使用最优方案）
	if len(idcard) != 18 {
		return false
	}

	// 前17位必须是数字
	for i := 0; i < 17; i++ {
		c := idcard[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// 最后一位：数字或X/x
	last := idcard[17]
	isDigit := last >= '0' && last <= '9'
	isX := last == 'X' || last == 'x'
	if !isDigit && !isX {
		return false
	}

	// 校验码验证（优化版本）
	// 权重因子
	weights := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	// 计算加权和
	sum := 0
	for i := 0; i < 17; i++ {
		// 直接使用ASCII码计算，避免 strconv.Atoi
		digit := int(idcard[i] - '0')
		sum += digit * weights[i]
	}

	// 计算校验码索引
	checkIndex := sum % 11

	// 校验码对照表
	checkCodes := "10X98765432"

	// 转换最后一位为大写
	var expectedCheck byte
	if last == 'x' {
		last = 'X'
	}

	// 比对校验码
	if checkIndex < len(checkCodes) {
		expectedCheck = checkCodes[checkIndex]
		return expectedCheck == last
	}

	return false
}
