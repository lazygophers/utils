# 测试文档

LazyGophers Utils 项目的完整测试体系和质量保证文档。

## 📊 项目测试概览

| 指标 | 整体状态 | 目标值 | 说明 |
|------|----------|--------|------|
| **总体覆盖率** | 85%+ | 90%+ | 🟡 良好，持续提升中 |
| **模块数量** | 20+ | - | ✅ 全覆盖 |
| **测试用例** | 500+ | - | ✅ 充足 |
| **CI/CD** | ✅ 自动化 | ✅ 完善 | ✅ GitHub Actions |

## 🧪 模块测试状态

### 测试覆盖率排行

| 模块 | 覆盖率 | 状态 | 测试用例数 | 最后更新 |
|------|--------|------|------------|----------|
| **candy** | 95%+ | 🟢 优秀 | 80+ | 2025-09-13 |
| **cryptox** | 90%+ | 🟢 优秀 | 60+ | 2025-09-12 |
| **stringx** | 92%+ | 🟢 优秀 | 45+ | 2025-09-10 |
| **json** | 90%+ | 🟢 优秀 | 30+ | 2025-09-08 |
| **network** | 85%+ | 🟡 良好 | 25+ | 2025-09-05 |
| **config** | 88%+ | 🟡 良好 | 20+ | 2025-09-03 |
| **routine** | 80%+ | 🟡 良好 | 15+ | 2025-09-01 |
| **wait** | 85%+ | 🟡 良好 | 18+ | 2025-08-30 |
| **runtime** | 88%+ | 🟡 良好 | 22+ | 2025-08-28 |
| **osx** | 90%+ | 🟢 优秀 | 35+ | 2025-08-25 |
| **hystrix** | 85%+ | 🟡 良好 | 12+ | 2025-08-20 |
| **xtime** | 72.7% | 🟨 中等 | 50+ | 2025-09-13 |
| **randx** | 85%+ | 🟡 良好 | 25+ | 2025-08-15 |
| **其他模块** | 80%+ | 🟡 良好 | 100+ | - |

### 测试质量分级

| 等级 | 覆盖率范围 | 状态 | 模块数量 | 改进计划 |
|------|------------|------|----------|----------|
| 🟢 **优秀** | 90%+ | 生产就绪 | 6个 | 保持现状 |
| 🟡 **良好** | 80-89% | 可用状态 | 10个 | 持续改进 |
| 🟨 **中等** | 70-79% | 需要改进 | 2个 | 优先提升 |
| 🔴 **不足** | <70% | 需要重点关注 | 0个 | - |

## 📋 测试文档分类

### 📊 [整体测试报告](project-test-report.md)
项目级别的综合测试分析

**包含内容**：
- 🎯 整体测试策略和方法论
- 📊 跨模块测试数据汇总
- 🔄 持续集成和自动化测试
- 📈 测试趋势和改进计划

### 📝 [模块测试详情](module-test-details.md)  
各个模块的详细测试报告

**包含内容**：
- 🧪 模块级测试用例分析
- 📊 功能覆盖率详情
- ⚡ 性能测试结果
- 🐛 已知问题和修复计划

## 🚀 测试运行指南

### 基础测试命令

```bash
# 运行所有模块测试
go test ./...

# 详细输出
go test -v ./...

# 测试覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### 模块级测试

```bash
# 测试特定模块
go test -v ./candy
go test -v ./xtime
go test -v ./cryptox

# 模块覆盖率
go test -cover ./candy
go test -cover ./xtime

# 模块基准测试
go test -bench=. ./candy
go test -bench=. ./xtime
```

### 高级测试选项

```bash
# 并发测试（检测竞态条件）
go test -race ./...

# 短测试（跳过耗时测试）
go test -short ./...

# 内存分析
go test -memprofile=mem.prof ./...

# CPU分析
go test -cpuprofile=cpu.prof ./...

# 测试超时设置
go test -timeout=30s ./...
```

## 🎯 测试策略

### 测试层次设计

#### 1. 单元测试 (Unit Tests)
**范围**: 函数和方法级别
**覆盖率目标**: 90%+

```go
func TestToString(t *testing.T) {
    tests := []struct {
        name  string
        input interface{}
        want  string
    }{
        {"int", 123, "123"},
        {"float", 3.14, "3.14"},
        {"bool", true, "1"},
        {"nil", nil, ""},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := candy.ToString(tt.input)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

#### 2. 集成测试 (Integration Tests)
**范围**: 模块间交互
**覆盖率目标**: 80%+

```go
func TestXTimeWithCandy(t *testing.T) {
    // 测试 xtime 和 candy 模块的协作
    cal := xtime.NowCalendar()
    data := cal.ToMap()
    
    // 使用 candy 进行数据转换
    jsonStr := candy.ToString(data)
    assert.NotEmpty(t, jsonStr)
}
```

#### 3. 基准测试 (Benchmark Tests)
**目标**: 性能验证和回归检测

```go
func BenchmarkToString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = candy.ToString(12345)
    }
}
```

#### 4. 压力测试 (Stress Tests)
**目标**: 高负载场景验证

```go
func TestConcurrentAccess(t *testing.T) {
    const numGoroutines = 100
    const numOperations = 1000
    
    var wg sync.WaitGroup
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < numOperations; j++ {
                _ = candy.ToString(j)
            }
        }()
    }
    wg.Wait()
}
```

### 测试分类标准

#### ✅ 必需测试 (Required)
- 所有公共 API
- 核心业务逻辑
- 错误处理路径
- 边界条件

#### 🎯 重点测试 (Priority)
- 高频使用功能
- 性能关键路径
- 安全相关功能
- 数据转换逻辑

#### 📊 补充测试 (Additional)
- 私有函数（重要的）
- 配置解析
- 兼容性测试
- 文档示例验证

## 📈 质量指标

### 代码质量门禁

#### 🔒 必须通过
- [ ] 所有测试用例通过
- [ ] 覆盖率 ≥ 80%
- [ ] 无竞态条件
- [ ] 静态检查通过

#### 🎯 建议达到
- [ ] 覆盖率 ≥ 90%
- [ ] 性能无明显退化
- [ ] 内存泄漏检测通过
- [ ] 文档示例可运行

### 持续监控指标

| 指标类型 | 监控频率 | 告警阈值 | 负责人 |
|---------|----------|----------|--------|
| **测试覆盖率** | 每次提交 | < 80% | 开发者 |
| **测试通过率** | 每次提交 | < 100% | CI系统 |
| **性能回归** | 每日 | > 10% | 性能团队 |
| **内存泄漏** | 每周 | 检测到 | 质量团队 |

## 🔄 CI/CD 集成

### GitHub Actions 配置

```yaml
name: Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: [1.24, 1.23]
        
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: ${{ matrix.go-version }}
        
    - name: Run tests
      run: |
        go test -race -coverprofile=coverage.out ./...
        
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
        
    - name: Run benchmarks
      run: go test -bench=. -benchmem ./...
```

### 质量检查流程

```bash
#!/bin/bash
# quality-check.sh

echo "=== 代码质量检查 ==="

# 1. 格式检查
echo "检查代码格式..."
go fmt ./...
goimports -w .

# 2. 静态分析
echo "执行静态分析..."
go vet ./...
golangci-lint run

# 3. 测试执行
echo "执行测试..."
go test -race -cover ./...

# 4. 基准测试
echo "执行基准测试..."
go test -bench=. -benchmem ./... > benchmark_results.txt

echo "=== 质量检查完成 ==="
```

## 📊 测试报告和分析

### 定期报告

#### 📅 日报 (每日自动生成)
- 当日测试执行情况
- 失败测试统计和分析
- 覆盖率变化趋势
- 性能基准对比

#### 📅 周报 (每周一发布)  
- 周内测试质量总结
- 新增测试用例统计
- 问题修复情况跟踪
- 下周改进计划

#### 📅 月报 (每月总结)
- 整体质量指标趋势
- 模块测试成熟度分析
- 技术债务评估
- 测试策略调整建议

### 分析维度

#### 📊 覆盖率分析
- 行覆盖率 (Line Coverage)
- 函数覆盖率 (Function Coverage)  
- 分支覆盖率 (Branch Coverage)
- 条件覆盖率 (Condition Coverage)

#### ⚡ 性能分析
- 执行时间趋势
- 内存使用模式
- 资源消耗统计
- 性能回归检测

## 🔧 测试工具和框架

### 核心测试框架

```go
// 主要使用标准库 + 增强工具
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/mock"
)
```

### 推荐工具链

| 工具 | 用途 | 安装命令 |
|------|------|----------|
| **testify** | 断言和模拟 | `go get github.com/stretchr/testify` |
| **ginkgo** | BDD测试框架 | `go get github.com/onsi/ginkgo/v2` |
| **gomega** | 匹配器库 | `go get github.com/onsi/gomega` |
| **gotests** | 测试代码生成 | `go get github.com/cweill/gotests` |
| **gocov** | 覆盖率工具 | `go get github.com/axw/gocov` |
| **go-junit-report** | JUnit报告 | `go get github.com/jstemmer/go-junit-report` |

### 测试辅助工具

```bash
# 生成测试代码
gotests -w -all *.go

# 覆盖率可视化
gocov test ./... | gocov-html > coverage.html

# JUnit格式报告
go test -v ./... 2>&1 | go-junit-report > report.xml
```

## 💡 最佳实践

### 测试编写规范

#### ✅ 推荐做法

```go
// 1. 使用表驱动测试
func TestFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    OutputType
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 测试逻辑
        })
    }
}

// 2. 清晰的测试名称
func TestToString_WithValidInput_ReturnsExpectedString(t *testing.T) {
    // 测试逻辑
}

// 3. 适当的设置和清理
func TestWithSetup(t *testing.T) {
    // Setup
    setUp := func() { /* 初始化 */ }
    tearDown := func() { /* 清理 */ }
    
    setUp()
    defer tearDown()
    
    // 测试逻辑
}
```

#### ❌ 避免的做法

```go
// 1. 避免测试名称不清晰
func TestFunc1(t *testing.T) { /* 不好的命名 */ }

// 2. 避免测试过于复杂
func TestEverything(t *testing.T) {
    // 测试太多功能，应该拆分
}

// 3. 避免硬编码依赖
func TestWithDatabase(t *testing.T) {
    db := sql.Open("mysql", "hardcoded-connection") // 不好的做法
}
```

## 🔗 相关文档

### 内部文档
- **[性能测试](../performance/)** - 基准测试和性能分析
- **[开发指南](../development/contributing.md)** - 测试规范和要求
- **[模块文档](../modules/)** - 各模块的具体测试情况

### 外部资源
- [Go Testing 官方文档](https://golang.org/pkg/testing/)
- [Testify 框架文档](https://github.com/stretchr/testify)
- [测试最佳实践](https://golang.org/doc/code.html#Testing)

---

*测试文档最后更新: 2025年09月13日*