# LazyGophers Utils - AI 开发指南

> 本项目编码约定和开发规范

---

## 核心约定

### 语言和沟通
- **必须使用中文交互** — 所有回复、进度同步、提问、总结都必须使用中文
- **技术术语保持原样** — API、JSON、HTTP、GORM 等不翻译

### 质量标准
- **测试覆盖率阈值**：90%（Makefile `COVERAGE_THRESHOLD=90`）
- **性能优化优先**：索引循环代替 range、预分配切片容量
- **Lint**：golangci-lint + gosec

### 代码风格（从代码库提取）
- **泛型广泛使用**：类型参数命名 `T`、`K`、`V`，约束 `any`、`comparable`
- **Must 模式**：`Must`/`MustSuccess`/`MustOk` 用于初始化和测试，库函数返回 error
- **错误处理**：使用 `github.com/pkg/errors` 包装，记录日志后传播
- **日志**：统一使用 `github.com/lazygophers/log`，`Errorf` 记录不终止、`Panicf` 致命

### 项目结构
- **扁平化多包架构**：candy、cache、human、xtime、wait 等 20+ 独立包
- **根目录工具**：must.go、orm.go、bufiox.go 等通用工具
- **ORM 模式**：实现 `driver.Valuer` 和 `sql.Scanner`，使用 `utils.Scan`/`utils.Value`

---

## 开发前必读

**Trellis 管理的项目**，开发前阅读：

1. **`.trellis/spec/backend/`** — 后端编码规范
   - `quality-guidelines.md` — 性能模式、禁止模式
   - `error-handling.md` — Must 模式、错误传播
   - `generics-guidelines.md` — 泛型使用规范
   - `coding-conventions.md` — 注释风格、函数设计
   - `testing-guidelines.md` — 测试规范、覆盖率
   - `benchmark-guidelines.md` — 基准测试规范

2. **`AGENTS.md`** — 项目知识库（包结构、反模式、质量检查）

3. **`CONTRIBUTING.md`** — 贡献指南

---

## 常用命令

```bash
# 测试
make test           # 运行所有测试
make coverage       # 生成覆盖率报告（HTML: docs/reports/coverage.html）

# Lint
make lint           # golangci-lint 检查

# 格式化
make fmt            # gofmt 格式化

# 基准测试
make bench          # 运行基准测试
```

---

## 禁止模式（从 spec 提取）

### 性能相关
- ❌ range 遍历值类型切片（使用索引循环）
- ❌ append 未预分配（使用 `make([]T, 0, len)`）
- ❌ 空接口类型断言未检查（使用 `if v, ok := anyData.(Type); ok`）

### 错误处理
- ❌ 忽略错误（特定场景除外）
- ❌ 库函数使用 Must（返回 error）
- ❌ 双 nil 返回（单 nil 或返回错误）

### 代码风格
- ❌ 过度抽象（< 3 处复用不提取）
- ❌ 过度防御（内部代码不加无谓 try-catch / nil-check）
- ❌ 写复述代码的注释（only WHY，not WHAT）

---

## 项目特定依赖

- **日志**：`github.com/lazygophers/log`（非标准 log）
- **错误包装**：`github.com/pkg/errors`
- **JSON**：`github.com/bytedance/sonic`（高性能）
- **测试断言**：`github.com/stretchr/testify`

---

## 参考资源

- **Trellis Workflow**：`.trellis/workflow.md`
- **Backend Spec 索引**：`.trellis/spec/backend/index.md`
- **项目知识库**：`AGENTS.md`
- **贡献指南**：`CONTRIBUTING.md`
