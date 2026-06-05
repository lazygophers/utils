# CLAUDE.md — LazyGophers Utils

Go 工具库，模块路径 `github.com/lazygophers/utils`，Go 1.26.2。

## 项目结构

- 根包 `utils`：`Must()`, `MustSuccess()`, `MustOk()`, `Value()`, `Scan()`
- 每个子目录 = 独立 Go package（如 `candy/`, `xtime/`, `cache/lru/`）
- `cache/` 下有 10 个子算法包（alfu, arc, fbr, lfu, lru, lruk, mru, slru, tinylfu, wtinylfu）
- `docs/` 是 Rspress 多语言文档站（zh-CN 默认 + zh-TW + en）

## 常用命令

```bash
make test          # 跑测试
make test-coverage # 覆盖率
make lint          # golangci-lint
make fmt           # gofmt
make all           # clean + fmt + lint + test + build
```

## 编码约定

- fail-fast 错误处理：内部用 `Must*` 系列断言，不堆 `if err != nil`
- 公共 API 必须有 godoc 注释
- 泛型优先：`candy/`, `anyx/` 等包广泛使用 Go generics
- 保持零分配路径和高性能基准
- 不引入重型依赖
- **测试文件命名**：`xxx.go` 的测试在 `xxx_test.go`，性能测试在 `xxx_benchmark_test.go`（禁止 `_coverage_test.go`、`_perf_test.go` 等变体后缀）

## 文档维护

- 三语言文档必须结构对称（zh-CN / zh-TW / en）
- 新增模块 → 三语言各补对应 .md
- 文档优先说明"适合做什么、适用场景、约束"，而非堆 API 列表
- `docs/` 下 `npm run dev` 预览，`npm run build` 构建
