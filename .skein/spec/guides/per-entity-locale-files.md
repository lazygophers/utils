# Per-Entity Per-Language Locale Files

> **适用规范**：大规模静态数据包（249 国 / 154 币 / N item 等）按 **1 entity 1 数据文件 + 1 entity × 1 lang 1 locale 文件** 拆分，文件名直接编码 entity 标识。

参考实现：`country/<alpha2>.go` + `country/<alpha2>_<lang>.go`，`currency/<code>.go` + `currency/<code>_<lang>.go`。

## 适用场景

- 静态查表数据包，entity 数 ≥ 50
- 每个 entity 有 ≥ 2 个多语言文本字段（名称 / 官方名 / 首都 等）
- 需要按语言按需编译（移动端 / 嵌入场景）

## 文件布局

```
<pkg>/
├── <pkg>.go          # 类型 + 方法
├── registry.go       # 索引 + Get/List
├── name.go           # 多语言名 + 回退链
├── constants.go      # 强类型常量 alias
├── cn.go             # 1 entity data + register（仅 var + init）
├── cn_en.go          # 英文 locale（无 build tag）
├── cn_zh.go          # 简中 locale（无 build tag）
├── cn_ja.go          # 日语 locale（build tag）
├── ...
└── us.go / us_en.go / us_zh.go / ...
```

## 命名规则

| 类型 | 模板 | 例 |
|---|---|---|
| Data 文件 | `<id>.go` | `cn.go`（ISO 3166 alpha-2）/ `cny.go`（ISO 4217） |
| Locale 文件 | `<id>_<lang>.go` | `cn_en.go` / `cn_zh.go` / `cn_zh_hant.go` |

`<id>` 用业界标准短码（小写，如 alpha-2 / ISO 4217）。不用国名 slug（如 `china.go`），保持文件名简短一致。

## build tag 规则

每 locale 文件，**默认带 build tag** `//go:build lang_<xx> || lang_all`，**例外两类**：

1. **项目默认语言**（en + zh）：始终无 build tag
2. **该 entity 的官方语言豁免**（限 country 这类有"官方语言"概念的实体）：若该 entity 的 `languages` 字段含某 lang，则其 `<id>_<lang>.go` 也无 build tag

例：
- `cn_en.go`：无 tag（项目默认 en）
- `cn_zh.go`：无 tag（项目默认 zh + China 官方语言 Chinese）
- `cn_ja.go`：`//go:build lang_ja || lang_all`（China 无 Japanese 官方）
- `jp_ja.go`：无 tag（Japan 官方 Japanese）
- `hk_zh_hant.go`：无 tag（HK 官方 zh-Hant）
- `cny_zh.go`：无 tag（项目默认 zh）
- `cny_ja.go`：`//go:build lang_ja || lang_all`（currency 无官方语言豁免）

## 文件内容模板

### Data 文件（无 locale 调用）

```go
package country

import xlanguage "golang.org/x/text/language"

// China — People's Republic of China.
var dataChina = &Country{
    alpha2:    "CN",
    alpha3:    "CHN",
    numeric:   156,
    languages: []xlanguage.Tag{xlanguage.Chinese},
    // ... 完整字段，禁缺省
    names:    make(map[xlanguage.Tag]string),
    official: make(map[xlanguage.Tag]string),
    capital:  make(map[xlanguage.Tag]string),
}

func init() { register(dataChina) }
```

### Locale 文件（仅 Register* 调用）

```go
// cn_zh.go：无 build tag（项目默认 + China 官方语言）
package country

import xlanguage "golang.org/x/text/language"

func init() {
    dataChina.RegisterName(xlanguage.Chinese, "中国")
    dataChina.RegisterOfficialName(xlanguage.Chinese, "中华人民共和国")
    dataChina.RegisterCapital(xlanguage.Chinese, "北京")
}
```

```go
// cn_ja.go：扩展语言 build tag
//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
    dataChina.RegisterName(xlanguage.Japanese, "中国")
    dataChina.RegisterOfficialName(xlanguage.Japanese, "中華人民共和国")
    dataChina.RegisterCapital(xlanguage.Japanese, "北京")
}
```

## init 顺序

Go 同包按文件名字母序 init。`<id>.go`（如 `cn.go`）字母序 < `<id>_<lang>.go`（如 `cn_en.go`），确保 data var 先注册再被 locale 引用。无需额外 init 协调。

## 与 locale_<lang>.go 单文件模式的区别

| 维度 | per-entity (本规范) | locale_<lang>.go (单文件) |
|---|---|---|
| 适用 | 大规模静态数据（≥ 50 entity） | 中小规模翻译（< 50 key） |
| 文件数 | entity × lang | 仅 lang 数 |
| 加 entity 影响 | 新 entity = 1 data + N locale 文件 | 编辑 N 个 locale 文件加 key |
| 加 lang 影响 | entity × 1 文件 | +1 文件 |
| 参考 | country / currency | xerror / validator |

实体数小、翻译 key 少 → 用 locale_<lang>.go（详见 [Locale File Convention](./locale-file-convention.md)）。
实体数大 → 用 per-entity。

## 数据源权威

ISO 标准（3166-1 / 4217）+ 联合国 / 维基百科。准确性优先于完整性；不确定的翻译走英文 fallback（RegisterName 仍调用，保证回退路径无意外）。

## 性能特性

- 包 init 完成后索引只读，`Get`/`Get*` 无锁
- locale RegisterName 写 map 用 RWMutex，运行期仅 RLock 读
- 切片字段返回副本（`CallingCodes()` / `Tlds()` 等）防外部 mutate
- hot path（`Get` / `常量直访` / `NameIn` 命中）0 alloc
- `Name()` goroutine-local 1 alloc / 48B（来自 `language.Get()` 内部，不可优化）

## 下游消费方约定（fake 等包参考）

下游消费 entity registry 的包（如 `fake/` 消费 `country/`），文件布局复用 per-entity 模式时必须满足两条**消费方契约**，否则默认 build 下 init panic 或 lookup 静默失败。

### 契约 1：build tag 镜像源 entity

**Wrong**：
```go
// fake/ad.go (Andorra)
package fake
// 无 build tag
func init() {
    c := country.Get("AD")
    register(&Locale{Country: c, ...}) // c == nil → register panic
}
```
- 默认 build 下 `country/ad.go` 有 `//go:build country_ad || country_all || country_europe || country_southern_europe`
- 该 tag 未启用 → `byAlpha2` 未注册 AD → `country.Get("AD")` 返回 nil
- `register(&Locale{Country: nil})` panic（registry.go 显式 nil 检查）

**Correct**：
```go
//go:build country_ad || country_all || country_europe || country_southern_europe

package fake

func init() {
    c := country.Get("AD")
    register(&Locale{Country: c, ...})
}
```

**规则**：
- 12 个常驻国（cn/de/fr/gb/hk/in/jp/kr/ru/sg/tw/us）源端无 tag → 消费端也无 tag
- 其余 237 国源端有 tag → 消费端 head -1 抓 `//go:build` 行原样回写
- 脚本批量生成时：`tag=$(head -1 country/${code}.go | grep "^//go:build" || echo "")`

### 契约 2：locale map key 用源 entity 返回的 tag

**Wrong**：
```go
// fake/cn_zh.go
var langZh = xlanguage.SimplifiedChinese   // ❌
localeCN.LastNames[langZh] = chineseLastNames
// Faker.lang 来自 country.China.OfficialLanguage() == xlanguage.Chinese
// → 查 localeCN.LastNames[Chinese] 返回 nil
```

**Correct**：
```go
var langZh = xlanguage.Chinese              // ✓ 与 country.China.OfficialLanguage() 一致
```

**规则**：locale data map key 必须 `==` 源 entity 暴露的 tag 值（`country.X.OfficialLanguage()` / `currency.Y.SpokenLangs()`）。stdlib `xlanguage.Chinese` 与 `xlanguage.SimplifiedChinese` 是不同 tag 实例，不能混用。

验证方法：写测试 `for c := range registry { assert.NotNil(t, c.locale.FirstNames[c.OfficialLanguage()]) }`。

## 关联约定

- [Locale File Convention](./locale-file-convention.md) — 中小翻译数据
- [Stdlib Language Tag](./stdlib-language-tag.md) — 公共 API tag stdlib 强制
- [No Anonymous Struct](./no-anonymous-struct.md) — Data var 用 `var dataXxx = &Entity{...}`
- [Error Handling Style](./error-handling-style.md)
