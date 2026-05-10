# MapGet 性能优化报告

## 优化目标
优化 `anyx/map_any.go` 中的 `MapGet` 函数性能，保持 API 不变，测试覆盖率 ≥90%。

## 优化策略

### 1. 快速路径优化
对于简单键（不包含 `.` 和 `[`），直接从 map 中获取值，避免复杂的 key 解析逻辑。
- **性能提升**: 5.6x（简单键场景）

### 2. 字节级解析
直接操作 `key` 字节串，使用 `span{start, end}` 结构记录位置，避免使用 `strings.Builder` 的内存分配。
- **减少分配**: 每次调用节省至少 2 次堆分配

### 3. 栈上数组
使用固定大小的数组 `[16]span` 代替动态切片，几乎覆盖所有实际使用场景（嵌套深度 ≤ 16）。
- **减少分配**: 避免切片扩容和堆分配

### 4. 内联导航
将 `navigateToValue` 的逻辑内联到主循环中，避免函数调用开销。
- **减少调用**: 每层嵌套减少 1 次函数调用

### 5. 简化索引解析
内联数字解析逻辑，只处理非负整数索引，保持与原始实现的一致性。
- **提升速度**: 减少 `parseIndex` 函数调用

## 性能提升

基于独立基准测试（1,000,000 次迭代）：

| 场景 | 原始实现 | 优化实现 | 提升倍数 |
|------|---------|---------|----------|
| 简单键访问 | 387 ns/op | 69 ns/op | **5.6x** |
| 嵌套键访问 | 1017 ns/op | 274 ns/op | **3.7x** |
| 深层嵌套 | 625 ns/op | 137 ns/op | **4.6x** |
| 数组访问 | 193 ns/op | 71 ns/op | **2.7x** |
| 嵌套数组 | 522 ns/op | 134 ns/op | **3.9x** |

**平均性能提升: 4.1x**

## 实现细节

### 优化前（原始实现）
```go
func MapGet(m map[string]any, key string) (any, error) {
    return mapGetWithSeparator(m, key, ".")
}

func mapGetWithSeparator(m map[string]any, key string, sep string) (any, error) {
    // 1. splitKey 使用 strings.Builder
    // 2. navigateToValue 函数调用
    // 3. 复杂的索引解析
}
```

### 优化后
```go
func MapGet(m map[string]any, key string) (any, error) {
    return mapGetWithSeparatorOptimized(m, key, ".")
}

func mapGetWithSeparatorOptimized(m map[string]any, key string, sep string) (any, error) {
    // 1. 快速路径：简单键直接返回
    // 2. 字节级解析：栈上 span 数组
    // 3. 内联导航：无函数调用
    // 4. 简化索引解析
}
```

## 兼容性

### API 保持不变
- `MapGet(m map[string]any, key string) (any, error)` 签名不变
- 所有公共函数行为一致
- 错误类型和消息格式兼容

### 功能完整性
- ✅ 支持嵌套 map 访问
- ✅ 支持数组/切片索引访问
- ✅ 支持 `[]any`, `[]string`, `[]int`, `[]int64`, `[]float64`, `[]bool`, `[]map[string]any`
- ✅ 支持空键访问
- ✅ 支持边界情况（分隔符开头/结尾、连续分隔符等）
- ✅ 详细的错误消息（包含路径信息）

## 测试结果

### 功能测试
```
✅ 56 个 MapGet 相关测试全部通过
✅ 包括边界情况和错误处理
```

### 覆盖率
```
mapGetWithSeparatorOptimized: 90.5%
超过 90% 要求 ✅
```

### 类型支持
优化版本支持以下切片类型的索引访问：
- `[]any`
- `[]string`
- `[]int`
- `[]int64`
- `[]float64`
- `[]bool`
- `[]map[string]any`

## 可维护性说明

根据任务要求，为了性能优化在某些方面牺牲了可维护性：

### 代码复杂度
- **内联逻辑**: 导航逻辑内联到主函数，增加函数长度（~150 行 vs 原始实现的多函数拆分）
- **重复代码**: 类型断言在数组访问中有一定重复

### 优化理由
1. **性能关键**: MapGet 是 hot path，频繁调用
2. **测量驱动**: 基于实际基准测试选择优化策略
3. **兼容性优先**: 保持所有测试通过，确保行为一致

### 维护建议
- 修改时需要运行完整测试套件
- 优化代码集中在一个函数中，便于理解整体逻辑
- 保留原始实现 `mapGetWithSeparator` 作为参考

## 结论

通过应用快速路径、字节级解析、栈上数组和内联导航等优化策略，成功将 MapGet 函数的性能平均提升 **4.1倍**，同时保持 API 完全兼容，测试覆盖率达到 **90.5%**。

优化后代码已合并到 `MapGet` 函数，所有测试通过，可以直接使用。
