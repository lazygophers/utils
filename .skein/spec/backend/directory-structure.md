# Directory Structure

> Package organization and file naming conventions.

---

## Layout

```
.
├── must.go / orm.go          # 根包 utils: Must, MustSuccess, MustOk, Value, Scan
├── candy/                    # 类型转换糖（泛型集合操作）
├── anyx/                     # map[string]any 深度导航
├── stringx/                  # Unicode-aware 字符串工具
├── json/                     # 基于 bytedance/sonic 的 JSON 增强
├── bufiox/                   # bufio 扫描扩展
├── xtime/                    # 高级时间（农历、节气、生肖）
├── xtime996/ xtime955/ xtime007/  # 工时调度
├── config/                   # 多格式配置加载（JSON/YAML/TOML/INI/HCL）
├── runtime/                  # 运行时信息
├── osx/                      # OS 操作
├── app/                      # 应用生命周期、版本解析
├── atexit/                   # 优雅退出（平台特定文件用 _darwin.go/_linux.go 后缀）
├── network/                  # HTTP 连接池、重试
├── cryptox/                  # 加密、哈希、ULID
├── pgp/                      # PGP 签名/加密
├── urlx/                     # URL 解析/构建
├── routine/                  # goroutine 工作池
├── wait/                     # 同步/异步等待、限流
├── hystrix/                  # 熔断器
├── singledo/                 # 单次执行去重
├── event/                    # 发布/订阅
├── cache/                    # 缓存算法（子目录各一个算法）
│   ├── lru/ arc/ lfu/ ...    # 每个 = 独立 package，实现 cache.Cache 接口
│   └── cache.go              # 公共接口定义
├── fake/                     # 测试数据生成
├── randx/                    # 安全随机
├── defaults/                 # struct 默认值填充
├── pyroscope/                # 性能 profiling
├── human/                    # 人性化格式
├── language/                 # BCP 47 语言标签 + goroutine-local 语言存储
└── validator/                # 数据验证（邮箱、手机号）
```

## Naming Rules

- **包名** = 目录名，全小写，无下划线（`cryptox` 不写 `crypt_x`）
- **文件名** = `snake_case.go`，同一包内的功能按文件拆分（如 `candy/map.go`, `candy/slice_chunk.go`）
- **平台特定** = `_darwin.go`, `_linux.go`, `_windows.go`, `_js.go`, `_aix.go` 等（参考 `atexit/`）
- **测试** = `xxx.go` 的单元测试在 `xxx_test.go`，性能测试在 `xxx_benchmark_test.go`（严格 1:1 映射，不允许 `xxx_coverage_test.go`、`xxx_perf_test.go` 等变体后缀）
- **接口文件** = 包内只有一个 `interface.go` 或内嵌在主文件中（如 `cache/cache.go`）
- **禁止 cmd/ 和 package main** = 纯库项目，不允许可执行文件
- **Build tags** = `en`（英文）和 `zh`（中文）locale 文件不加 `//go:build`，始终注册；其他语言保留 `//go:build lang_xx || lang_all`

## New Package Checklist

1. 目录名 = 包名，小写无下划线
2. 所有导出函数有 godoc 注释
3. `*_test.go` 覆盖核心路径
4. 三语言文档各补 `.md`（`docs/zh-CN/modules/<category>/<name>.md` + zh-TW + en）
5. 如有复杂子结构（如 cache），每个子算法独立子包

## Gotcha: 根包函数 vs 子包函数

> **Warning**: 根包 `utils` 只导出 `Must`, `MustSuccess`, `MustOk`, `Value`, `Scan`。`Validate` 等功能在独立子包（如 `validator/`）中。
>
> 新增导出函数后，必须同步验证 README 模块表、CLAUDE.md 函数列表、三语言文档是否引用了正确的包路径。历史上出现过 README 列了不存在的 `validate.go` 根包文件的 bug。
