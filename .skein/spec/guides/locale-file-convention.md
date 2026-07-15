# Locale File Convention

> **强制规范**：包内置多语言翻译按语言拆 `locale_<lang>.go` 文件，各自 `init()` 注册；en/zh 默认编译，其他语言走 build tag 选择启用。

---

## 适用场景

任何 utils 子包包含**内置多语言翻译数据**（错误码消息 / 校验提示 / 默认词条）时强制采用。

参考实现：`validator/locale_*.go`、`xerror/locale_*.go`。

## 文件布局

```
<pkg>/
├── codes.go           # 公共常量 + 构造器（无翻译数据）
├── locale_en.go       # 英文翻译 init，无 build tag
├── locale_zh.go       # 简体中文 init，无 build tag
├── locale_zh_tw.go    # //go:build lang_zh_tw || lang_all
├── locale_ja.go       # //go:build lang_ja || lang_all
├── locale_ko.go       # //go:build lang_ko || lang_all
├── locale_es.go       # //go:build lang_es || lang_all
├── locale_fr.go       # //go:build lang_fr || lang_all
├── locale_ru.go       # //go:build lang_ru || lang_all
└── locale_ar.go       # //go:build lang_ar || lang_all
```

## 命名规则

- 文件名：`locale_<lang>.go`（小写 + 下划线分隔，匹配 BCP47 但用 `_` 替 `-`）
  - `zh-CN` → `locale_zh.go`（base 即可）
  - `zh-Hant` → `locale_zh_tw.go`（约定俗成的繁体标记）
- **不允许**：`codes_xx.go` / `messages_xx.go` / `i18n_xx.go` 等变体
- `locale.go` （无后缀）保留给 `LocaleConfig` 类型定义 + 全局 registry / init helper

## build tag 策略

- `en` + `zh` 是默认必装语言，**不加 build tag**
- 其他语言加 `//go:build lang_<lang> || lang_all`
  - 单语：`go build -tags lang_ja`
  - 全开：`go build -tags lang_all`
  - 默认：仅 en + zh

## 文件内容模板

```go
//go:build lang_ja || lang_all

package xerror

func init() {
    registerBuiltinLocale("ja", map[int]string{
        CodeSystem:       "システムエラー",
        CodeInvalidParam: "パラメータが無効です",
        // ...
    })
}
```

或验证器模式：

```go
// 无 build tag，always
package validator

import "golang.org/x/text/language"

func init() {
    RegisterLocaleConfig("en", &LocaleConfig{
        Language: language.Make("en"),
        Region:   "US",
        Messages: map[string]string{
            "required": "{field} is required",
            // ...
        },
    })
}
```

## init 注册去向

各 `locale_<lang>.go` init() 把翻译注册到**当前包的全局表**或**i18n.Default**：

- **自有 registry 类型**（如 validator.RegisterLocaleConfig）→ 包内 sync.Map
- **i18n 二级 catalog**（如 xerror）→ `i18n.Default.Register(tag, key, value)`，按 errorKey/messageKey 拼接

## 为什么

1. **按需编译**：避免默认 binary 包含 8 种语言数据（每种约 1-5 KB，累计可观）；移动端 / 嵌入场景受益
2. **添加语言无副作用**：新语言 = 新文件 + build tag，不动 codes.go 主表
3. **跨包一致**：utils/validator + utils/xerror 同模式，未来新内置多语言包延续
4. **lang_all 全量**：CI / 通用 binary 一行 tag 全开，开发期友好
5. **avoid init order trap**：每个 locale_xx.go 独立 init，无序依赖；翻译数据隔离

## 适配既有代码

新增包带内置翻译时：
1. 抽 `LocaleConfig`（或调用 i18n.Default.Register）注册入口
2. en/zh 文件无 build tag
3. 其他语言文件按 `//go:build lang_xx || lang_all`
4. 文件名严格 `locale_<lang>.go`

## 关联约定

- [Stdlib Language Tag](./stdlib-language-tag.md) — 公共 API 必须用 xlanguage.Tag
- [Error Handling Style](./error-handling-style.md)
- [No Anonymous Struct](./no-anonymous-struct.md)

## 参考实现

- `utils/validator/locale_*.go`（13 语言，stdlib 公共 LocaleConfig.Language）
- `utils/xerror/locale_*.go`（9 语言，注册到 i18n.Default）
