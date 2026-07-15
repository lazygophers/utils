# Quality Guidelines

> Testing, lint, coverage, and benchmark conventions.

---

## Testing

**框架**：`testing` + `github.com/stretchr/testify/assert`

```go
// ✅ 标准测试模式
func TestFeature(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"basic", "hello", "HELLO"},
        {"empty", "", ""},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := Function(tc.input)
            assert.Equal(t, tc.expected, result)
        })
    }
}
```

**文件命名**（严格）：
- `xxx.go` 的单元测试 **只允许** 在 `xxx_test.go`
- `xxx.go` 的性能测试 **只允许** 在 `xxx_benchmark_test.go`
- 禁止变体后缀（`_coverage_test.go`、`_perf_test.go`、`_quick_test.go` 等）

**函数命名**：`Test<Function>_<Scenario>` 或 `Test<Function>_<EdgeCase>` — 见 `orm_test.go`（`TestScan_EdgeCases`, `TestScanValue_RoundTrip`, `TestScan_SpecialCases`）。

**并发测试**：关键路径必须有并发安全测试（如 `TestScanValue_ConcurrentAccess`）。

**Panic 测试**：Must* 函数用 `assert.Panics` 验证：

```go
assert.Panics(t, func() {
    utils.Must(func() error { return errors.New("boom") }())
})
```

## Benchmarks

每个性能关键函数有 `Benchmark*` 测试，放在对应的 `xxx_benchmark_test.go` 文件中：

```go
func BenchmarkFeature(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Function(input)
    }
}
```

验证包有对比基准（如 `validator/` 里的 `BenchmarkEmail_OriginalRegex` vs `BenchmarkEmail_Optimized`）。

## Lint

**工具**：`golangci-lint v2.4.0+`，配置在 `.golangci.yml`。

**启用的 linter**：bodyclose, contextcheck, durationcheck, errname, nilerr, nilnil, gosec, asciicheck, bidichk 等。

**禁用的 linter**：errcheck, govet, cyclop, dupl, funlen 等（工具库宽松风格，不过度限制函数长度/复杂度）。

**Gosec 豁免规则**：
- `G103` unsafe 使用（性能优化）
- `G401/G501/G505` cryptox 中 MD5/SHA1（向后兼容）
- `G404` fake/randx/stringx 中非加密随机
- `G304` json/osx/config 中变量路径文件操作
- `G115` 字节/rune 转换中的整数溢出
- 测试文件豁免 `G104/G302/G304/G306/G204/G301/G702/G101`

## Coverage

目标 ~98.7%。`make test-coverage` 生成覆盖率报告。

## Makefile Targets

```bash
make test           # go test ./...
make test-coverage  # 覆盖率报告
make test-race      # 竞态检测
make benchmark      # 基准测试
make lint           # golangci-lint
make lint-fix       # 自动修复
make fmt            # gofmt
make all            # clean + fmt + lint + test + build
```

## Forbidden Patterns

- ❌ 不在库公共 API 返回 error 的路径上用 log.Fatal/Fatalf
- ❌ 不在非 Must 函数中 panic（返回 error 让调用方决定）
- ❌ 不引入重型依赖（当前依赖：sonic, go-crypto, pyroscope, testify, gomonkey）
- ❌ 不在 `_test.go` 文件外使用 `gomonkey`（仅用于 monkey-patch 测试）
- ❌ 不用 `context.Context` 做协程级状态传递（用 goroutine-local 方案，如 `runtime.Stack` + `sync.Map`）
- ❌ 不做防御性编程：内部代码无 nil-check / try-catch，只在系统边界校验

## Naming Conventions

- **缩写词**：首字母大写其余小写（`Id` 不写 `ID`，`Http` 不写 `HTTP`，`Url` 不写 `URL`）
- **函数命名**：`goroutineId`（不是 `goroutineID`）

```go
// ✅ 正确
func goroutineId() uint64
type RequestId string

// ❌ 错误
func goroutineID() uint64
type RequestID string
```

## Common Mistakes: 测试文件重构

> **Warning**: 合并外部测试包（`package X_test`）到内部测试包（`package X`）时必须做三件事，否则编译失败。

### 外部 → 内部测试包转换

```go
// ❌ 错误：package X 内 import 了自身包 → import cycle
package defaults
import "github.com/lazygophers/utils/defaults"  // cycle!
result := defaults.Set(&cfg)

// ✅ 正确：package X 内直接调用
package defaults
result := Set(&cfg)
```

**转换检查清单**：
1. `package X_test` → `package X`
2. 删除 self-import（`"github.com/lazygophers/utils/X"`）
3. 全局替换 `X.Symbol` → `Symbol`（用 `\bX\.(\w)` 正则）

### 跨文件重声明

同一包内多个 `_test.go` 文件合并后，`var`/`const`/`type`/`func` 同名会 redeclared。合并时必须去重，保留第一个声明。

### `math/rand` 版本冲突

```go
// ❌ 两个都 import 为 rand → redeclared
import (
    "math/rand"
    "math/rand/v2"
)

// ✅ 给 v1 加别名
import (
    stdrand "math/rand"
    "math/rand/v2"
)
```

区分 API：v1 用 `Intn()`/`Shuffle()`/`NewSource()`，v2 用 `IntN()`/`Uint64()`。
