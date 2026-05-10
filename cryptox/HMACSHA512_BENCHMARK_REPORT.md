# HMACSHA512 性能优化测试报告

## 测试环境
- **CPU**: Apple M3
- **架构**: arm64
- **操作系统**: darwin
- **Go 版本**: go1.x
- **测试次数**: 每个方案运行 3 次，每次 2 秒

## 原始实现
```go
func HMACSHA512[M string | []byte](key, message M) string {
    h := hmac.New(sha512.New, []byte(key))
    _, _ = h.Write([]byte(message))
    return fmt.Sprintf("%x", h.Sum(nil))
}
```

**基准性能**:
- 平均耗时: 685.6 ns/op
- 内存分配: 1016 B/op
- 分配次数: 8 allocs/op

## 优化方案对比

| 方案 | 平均耗时 (ns/op) | 内存分配 (B/op) | 分配次数 | 性能提升 | 描述 |
|------|------------------|----------------|----------|----------|------|
| **原始 (fmt.Sprintf)** | 685.6 | 1016 | 8 | - | 基准 |
| Solution 1: 手动 hex + 数组 | 632.0 | 992 | 7 | +8.5% | 手动编码，固定数组 |
| Solution 2: 手动 hex + 切片预分配 | 620.2 | 992 | 7 | **+10.5%** | 手动编码，预分配切片 |
| Solution 3: 标准库 hex.EncodeToString | 627.8 | 1120 | 8 | +9.2% | 标准库，额外分配 |
| Solution 4: 标准库 hex.Encode + 预分配 | 627.7 | 992 | 7 | +9.2% | 标准库，预分配输出 |
| Solution 5: 循环展开 4x | 615.0 | 992 | 7 | **+11.5%** | 4次循环展开 |
| Solution 6: 循环展开 8x | 618.5 | 992 | 7 | +10.9% | 8次循环展开 |
| Solution 7: 查表优化 (256条目) | 622.5 | 992 | 7 | +10.2% | 预生成查找表 |
| Solution 8: Unsafe 字符串构造 | 615.9 | 864 | 6 | **+11.3%** | Unsafe 避免复制 |
| Solution 9: 混合优化 | 617.3 | 992 | 7 | +11.1% | 局部变量优化 |
| Solution 10: 完全展开 64 字节 | 616.2 | 992 | 7 | +11.3% | 完全手动展开 |
| Solution 11: 预分配 buffer | 635.9 | 992 | 7 | +7.8% | range 循环预分配 |

## 关键发现

### 1. 性能提升最大的方案
- **Solution 5: 循环展开 4x** - **11.5%** 性能提升
- **Solution 2: 手动 hex + 切片预分配** - **10.5%** 性能提升
- **Solution 8: Unsafe 字符串构造** - **11.3%** 性能提升，内存分配减少 15%

### 2. 内存效率最优
- **Solution 8 (Unsafe)**: 864 B/op, 6 allocs/op
  - 相比原始实现减少 15% 内存分配
  - 减少一次内存分配

### 3. 代码复杂度 vs 性能权衡
- **Solution 2**: 代码简洁，性能优秀 (+10.5%)，推荐作为生产代码
- **Solution 5**: 性能最佳 (+11.5%)，但代码可读性稍差
- **Solution 8**: 内存效率最优，但使用 unsafe 包，需要谨慎评估

### 4. 标准库方案
- **Solution 3** 和 **Solution 4** 使用 `encoding/hex`
- 性能接近手动优化方案
- 代码可读性最好，维护成本低

## 推荐方案

### 生产环境推荐：**Solution 2** (手动 hex + 切片预分配)
**理由**：
1. 性能提升显著 (+10.5%)
2. 代码简洁可读
3. 不使用 unsafe 包
4. 与项目现有优化风格一致

```go
func HMACSHA512[M string | []byte](key, message M) string {
	h := hmac.New(sha512.New, []byte(key))
	_, _ = h.Write([]byte(message))
	sum := h.Sum(nil)

	const hexchars = "0123456789abcdef"
	result := make([]byte, 128)
	for i := 0; i < 64; i++ {
		v := sum[i]
		result[i*2] = hexchars[v>>4]
		result[i*2+1] = hexchars[v&0x0f]
	}
	return string(result)
}
```

### 性能极致场景：**Solution 5** (循环展开 4x)
**理由**：
1. 性能最佳 (+11.5%)
2. 内存分配优化
3. 代码复杂度可接受

```go
func HMACSHA512[M string | []byte](key, message M) string {
	h := hmac.New(sha512.New, []byte(key))
	_, _ = h.Write([]byte(message))
	sum := h.Sum(nil)

	const hexchars = "0123456789abcdef"
	var result [128]byte

	for i := 0; i < 64; i += 4 {
		v0, v1, v2, v3 := sum[i], sum[i+1], sum[i+2], sum[i+3]
		result[i*2] = hexchars[v0>>4]
		result[i*2+1] = hexchars[v0&0x0f]
		result[(i+1)*2] = hexchars[v1>>4]
		result[(i+1)*2+1] = hexchars[v1&0x0f]
		result[(i+2)*2] = hexchars[v2>>4]
		result[(i+2)*2+1] = hexchars[v2&0x0f]
		result[(i+3)*2] = hexchars[v3>>4]
		result[(i+3)*2+1] = hexchars[v3&0x0f]
	}
	return string(result[:])
}
```

## 性能提升分析

### fmt.Sprintf 瓶颈
1. **反射开销**: fmt.Sprintf 使用反射解析格式参数
2. **动态内存**: 需要额外的缓冲区
3. **接口转换**: 将 []byte 转换为 any interface{}

### 优化技术
1. **手动 hex 编码**: 避免反射，直接查表
2. **预分配**: 减少 make() 调用开销
3. **固定数组**: 使用 [128]byte 替代切片，减少 GC 压力
4. **循环展开**: 减少循环控制开销
5. **Unsafe**: 避免字节数组复制

## 与项目其他函数对比

项目中已优化的函数性能提升：
- **HMACSHA1**: ~2x 提升
- **HMACSHA256**: ~1.3x 提升
- **HMACSHA384**: ~14% 提升
- **HMACSHA512** (本次): **~11.5%** 提升

**注意**: HMACSHA512 的性能提升相对较小，可能原因：
1. SHA512 计算时间占比更大 (64 字节 vs 32 字节)
2. hex 编码时间占比相对较小
3. M3 芯片对 SHA512 指令优化较好

## 结论

所有 11 种优化方案均成功实现性能提升，其中：
- **最佳性能**: Solution 5 (循环展开 4x) - +11.5%
- **最佳平衡**: Solution 2 (手动 hex + 切片预分配) - +10.5%
- **最低内存**: Solution 8 (Unsafe) - 15% 内存减少

**推荐采用 Solution 2** 作为生产实现，在性能、可读性和安全性之间取得最佳平衡。
