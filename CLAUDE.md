# CLAUDE.md — LazyGophers Utils

Go 工具库，模块路径 `github.com/lazygophers/utils`，Go 1.26.2。

## 项目结构

- 根包 `utils`：`Must()`, `MustSuccess()`, `MustOk()`, `Value()`, `Scan()`
- 每个子目录 = 独立 Go package（如 `candy/`, `xtime/`, `cache/lru/`）
- `cache/` 下有 10 个子算法包（alfu, arc, fbr, lfu, lru, lruk, mru, slru, tinylfu, wtinylfu）
- `language/` 提供语言标签解析 + goroutine-local 语言存储（`SetDefault`/`Default`/`Set`/`Get`）
- `i18n/` 多语言翻译：`Pack`/`I18n` + `Localize`/`LocalizeWithLang` + `LoadFile`/`LoadFs`/`LoadDir`/`LoadFsDir`/`LoadLocalizes`，内置 json/yaml/toml 解析，复用 `language` 包做 goroutine-local 语言
- `xerror/` 错误码 + 堆栈 + 多 error 聚合 + panic/recover 辅助（`New`/`Wrap`/`Join`/`Try`/`Recover`）
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
- **禁止 `cmd/` 目录和 `package main`**：这是纯库项目，不包含可执行文件
- **build tag 策略**：`en`（英文）和 `zh`（中文）的 locale 文件不加 `//go:build` tag，始终注册；其他语言（ja/ko/ar/es/fr/ru/zh-tw）保留 `//go:build lang_xx || lang_all`
- **命名约定**：缩写词用首字母大写其余小写（`Id` 而非 `ID`，`Http` 而非 `HTTP`）
- **禁止 `context.Context`**：不在本库中引入 context 依赖，协程级存储用 goroutine-local 方案
- **禁止防御性编程**：内部代码不做 nil-check / try-catch，只在系统边界校验
- **Error 接收与判断必须分两行**：禁止 `if err := X(); err != nil {}` 内联形态；先 `err := X()` 再 `if err != nil {}`。详见 [`.trellis/spec/guides/error-handling-style.md`](.trellis/spec/guides/error-handling-style.md)
- **变量声明禁用 anonymous struct**：包括 `var X = struct{}{}` / `:= struct{}{}` / `tests := []struct{}{}` 等所有形态；必须先 `type Name struct{...}` 再使用。`struct{}` 空结构（chan/map 集合语义）例外。详见 [`.trellis/spec/guides/no-anonymous-struct.md`](.trellis/spec/guides/no-anonymous-struct.md)

## 文档维护

- 三语言文档必须结构对称（zh-CN / zh-TW / en）
- 新增模块 → 三语言各补对应 .md
- 文档优先说明"适合做什么、适用场景、约束"，而非堆 API 列表
- `docs/` 下 `npm run dev` 预览，`npm run build` 构建
- `docs/` 基于 Rspress 框架，修改时参考 [Rspress 文档](https://rspress.rs/llms.txt)
