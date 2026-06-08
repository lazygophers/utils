# CLAUDE.md — LazyGophers Utils

Go 工具库，模块路径 `github.com/lazygophers/utils`，Go 1.26.2。

## 项目结构

- 根包 `utils`：`Must()`, `MustSuccess()`, `MustOk()`, `Value()`, `Scan()`
- 每个子目录 = 独立 Go package（如 `candy/`, `xtime/`, `cache/lru/`）
- `cache/` 下有 10 个子算法包（alfu, arc, fbr, lfu, lru, lruk, mru, slru, tinylfu, wtinylfu）
- `language/` 提供语言标签解析 + goroutine-local 语言存储（`SetDefault`/`Default`/`Set`/`Get`）
- `i18n/` 多语言翻译：`Pack`/`I18n` + `Localize`/`LocalizeWithLang` + `LoadFile`/`LoadFs`/`LoadDir`/`LoadFsDir`/`LoadLocalizes`，内置 json/yaml/toml 解析，复用 `language` 包做 goroutine-local 语言
- `xerror/` 错误码 + 堆栈 + 多 error 聚合 + panic/recover 辅助（`New`/`Wrap`/`Join`/`Try`/`Recover`）
- `country/`：ISO 3166-1 国家/地区数据（249 区，1 区 1 文件）+ 多语言名/官方名/首都 + 时区/区号/TLD/官方语言（双形态 API：`Get(code)` 或常量 `country.China`）
- `currency/`：ISO 4217 货币数据（154 种，1 币 1 文件）+ 多语言名（en/zh 默认 + 7 扩展 build tag），双形态 API：`Get(code)` 或常量 `currency.Cny`
- `fake/`：假数据生成器（faker 风格）。`New(country.Code, ...Option)` 实例 + 全局函数（`fake.Name()` 等）；249 国全骨架 + CN/US/JP 真数据（身份证含 GB 11643 校验、My Number NTA 算法、SSN 保留段排除）；`math/rand/v2` 可 seed 复现；强耦合 `country.Code`，goroutine-local 语言走 `language` 包
- **`fake/<code>.go` 必须镜像 `country/<code>.go` 的 build tag**：12 常驻国（cn/de/fr/gb/hk/in/jp/kr/ru/sg/tw/us）无 tag，其余 237 国走 `//go:build country_<xx> || country_all || country_<region>`；否则默认 build 下 `country.Get(code)` 返回 nil 触发 register panic
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
- **公共 API 语言 tag 用 stdlib**：跨包暴露的 tag 参数 / 返回 / 字段必须用 `golang.org/x/text/language.Tag`（值类型），不得用 `utils/language.*Tag`。`utils/language` 仅内部用于 FallbackChain / intern / goroutine-local。详见 [`.trellis/spec/guides/stdlib-language-tag.md`](.trellis/spec/guides/stdlib-language-tag.md)
- **内置多语言按 `locale_<lang>.go` 拆文件**：en/zh 默认不加 tag，其他语言走 `//go:build lang_xx || lang_all`；文件名严格 `locale_<lang>.go`。详见 [`.trellis/spec/guides/locale-file-convention.md`](.trellis/spec/guides/locale-file-convention.md)
- **大规模静态数据按 `<id>.go` + `<id>_<lang>.go` 拆**：249 国 / 154 币这类查表数据，1 entity 1 数据文件 + 1 entity × 1 lang 1 locale 文件；含「官方语言豁免」build tag 规则。详见 [`.trellis/spec/guides/per-entity-locale-files.md`](.trellis/spec/guides/per-entity-locale-files.md)

## 文档维护

- **每个含 `.go` 的目录必须有 `llms.txt`**（给 LLM 的包级说明）：H1 + 导入路径 + 功能 + 快速开始 + 核心 API + 文件结构表，内容须与代码一致（增/改 API 同步更新）；目录含 `.go` 子目录时须加「子目录索引」段列出各子目录相对路径链接。质量样板见 [`validator/llms.txt`](validator/llms.txt)
- 三语言文档必须结构对称（zh-CN / zh-TW / en）
- 新增模块 → 三语言各补对应 .md
- 文档优先说明"适合做什么、适用场景、约束"，而非堆 API 列表
- `docs/` 下 `npm run dev` 预览，`npm run build` 构建
- `docs/` 基于 Rspress 框架，修改时参考 [Rspress 文档](https://rspress.rs/llms.txt)
