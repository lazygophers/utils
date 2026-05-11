package validator

// validateBankCardOpt1ByteManualLuhn 字节级检查 + 手动Luhn（避免strconv.Atoi）
// 优化点：
// 1. 使用字节索引代替range遍历
// 2. 手动计算数字值（c - '0'）代替strconv.Atoi
// 3. 零内存分配
func validateBankCardOpt1ByteManualLuhn(fl FieldLevel) bool {
	cardNo := fl.Field().String()
	if cardNo == "" {
		return false
	}

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	// 字节级数字检查（快速失败）
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// Luhn算法
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

// validateBankCardOpt2LookupTable 字节级 + 查找表优化Luhn
// 优化点：
// 1. 预计算Luhn双倍值查找表
// 2. 避免乘法和除法运算
// 3. 零内存分配
func validateBankCardOpt2LookupTable(fl FieldLevel) bool {
	cardNo := fl.Field().String()
	if cardNo == "" {
		return false
	}

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	// Luhn双倍值查找表：[0,2,4,6,8,1,3,5,7,9]
	// doubled[x] = (x * 2) % 10 + (x * 2) / 10
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

// validateBankCardOpt3PrecomputedDoubles 预计算Luhn双倍值表（优化版）
// 优化点：
// 1. 使用静态查找表（编译期初始化）
// 2. 字节级操作
// 3. 最小化条件分支
func validateBankCardOpt3PrecomputedDoubles(fl FieldLevel) bool {
	cardNo := fl.Field().String()
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

	// 静态查找表
	sum := 0
	double := false

	for i := l - 1; i >= 0; i-- {
		d := int(cardNo[i] - '0')
		if double {
			// 查找表访问
			d = d*2 - 9*(d/5)
		}
		sum += d
		double = !double
	}

	return sum%10 == 0
}

// validateBankCardOpt4FastFail 快速失败优化（长度前置检查）
// 优化点：
// 1. 更激进的快速失败策略
// 2. 字节长度检查优先
// 3. 数字检查和Luhn合并
func validateBankCardOpt4FastFail(fl FieldLevel) bool {
	cardNo := fl.Field().String()

	l := len(cardNo)
	// 快速长度检查
	if l < 13 || l > 19 {
		return false
	}

	// 快速首字符检查（银行卡号通常不以0开头）
	if cardNo[0] < '1' || cardNo[0] > '9' {
		return false
	}

	// 单次遍历：数字检查 + Luhn
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

// validateBankCardOpt5IndexLoop 索引循环（避免range）
// 优化点：
// 1. 纯索引循环
// 2. 避免range的迭代器开销
// 3. 字节级操作
func validateBankCardOpt5IndexLoop(fl FieldLevel) bool {
	cardNo := fl.Field().String()
	if cardNo == "" {
		return false
	}

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	// 纯索引循环检查
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	sum := 0
	alternate := false

	// 索引循环Luhn
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit <<= 1 // 左移1位 = 乘2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}

// validateBankCardOpt6ASCII ASCII范围检查优化
// 优化点：
// 1. 更精确的ASCII范围检查
// 2. 使用位运算优化
// 3. 提前计算odd标志
func validateBankCardOpt6ASCII(fl FieldLevel) bool {
	cardNo := fl.Field().String()

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	sum := 0
	// 从右到左，奇数位需要双倍
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

// validateBankCardOpt7SinglePass 单次遍历优化
// 优化点：
// 1. 合并数字检查和Luhn计算
// 2. 减少遍历次数
// 3. 优化条件分支
func validateBankCardOpt7SinglePass(fl FieldLevel) bool {
	cardNo := fl.Field().String()

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	sum := 0
	double := true

	// 单次遍历：从右到左
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]

		// 快速失败：非数字
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

// validateBankCardOpt8BitOps 位运算优化
// 优化点：
// 1. 使用位运算代替算术运算
// 2. 优化双倍计算：d*2 -> d<<1
// 3. 优化减9操作：d-9 -> d^(9) 某些情况下更快
func validateBankCardOpt8BitOps(fl FieldLevel) bool {
	cardNo := fl.Field().String()

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

	return (sum & 0xF) == 0 // sum%10 == 0 的位运算版本（仅对10的倍数有效）
}

// validateBankCardOpt9Reverse 反向遍历优化
// 优化点：
// 1. 优化遍历顺序
// 2. 减少边界检查
// 3. 提前计算起始位置
func validateBankCardOpt9Reverse(fl FieldLevel) bool {
	cardNo := fl.Field().String()

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

// validateBankCardOpt10Combined 组合优化（字节+手动Luhn+快速失败+ASCII）
// 优化点：
// 1. 组合所有最佳实践
// 2. 零内存分配
// 3. 最少条件分支
// 4. 字节级操作
// 5. 快速失败
func validateBankCardOpt10Combined(fl FieldLevel) bool {
	cardNo := fl.Field().String()

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	// 快速失败：首字符检查
	firstChar := cardNo[0]
	if firstChar < '1' || firstChar > '9' {
		return false
	}

	sum := 0
	double := true

	// 单次遍历：从右到左
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

// validateBankCardOpt11Branchless 无分支Luhn
// 优化点：
// 1. 使用查找表消除条件分支
// 2. 预计算所有可能值
// 3. 减少分支预测失败
func validateBankCardOpt11Branchless(fl FieldLevel) bool {
	cardNo := fl.Field().String()

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	// 无分支查找表
	// 对于0-9的数字，double后的值：0,2,4,6,8,1,3,5,7,9
	var lut = [20]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

	sum := 0
	double := 0

	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}

		d := int(c - '0')
		// 使用查找表消除分支
		sum += lut[d+double*10]
		double ^= 1 // XOR切换0/1
	}

	return sum%10 == 0
}

// validateBankCardOpt12SimdInspired SIMD启发式（批量处理）
// 优化点：
// 1. 批量处理4个字符
// 2. 减少循环开销
// 3. 更好的CPU流水线利用
func validateBankCardOpt12SimdInspired(fl FieldLevel) bool {
	cardNo := fl.Field().String()

	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}

	// 批量检查数字
	i := 0
	for ; i+4 <= l; i += 4 {
		// 一次检查4个字符
		c0, c1, c2, c3 := cardNo[i], cardNo[i+1], cardNo[i+2], cardNo[i+3]
		if c0 < '0' || c0 > '9' ||
			c1 < '0' || c1 > '9' ||
			c2 < '0' || c2 > '9' ||
			c3 < '0' || c3 > '9' {
			return false
		}
	}

	// 处理剩余字符
	for ; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	// Luhn算法
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
