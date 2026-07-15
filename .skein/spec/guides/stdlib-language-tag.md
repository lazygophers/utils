# Stdlib Language Tag Convention

> **强制规范**：所有跨包暴露的语言 tag 参数 / 返回 / 字段，必须使用 `golang.org/x/text/language.Tag`（stdlib 值类型），不得用 `utils/language.*Tag`（项目包装类型）。

---

## 适用范围

- 所有 Go 模块的**公共 API**：函数参数、返回值、struct 字段、interface 方法签名
- 测试代码同理（外部测试用 `xlanguage.Make` 构造）
- 例外：`utils/language` 包**内部**实现仍可用自有 `*Tag`（其 FallbackChain / interning / goroutine-local 存储需要）

---

## 正确写法

```go
import xlanguage "golang.org/x/text/language"

// ✅ 公共 API 用 stdlib 值类型
func Register(tag xlanguage.Tag, key, value string)
func (p *I18n) LocalizeWithLang(tag xlanguage.Tag, key string, args ...any) string

// ✅ 内部需要 utils/language 的 FallbackChain 时，边界转换
func internalLookup(tag xlanguage.Tag, key string) string {
    utilsTag := language.Make(tag.String())
    chain := utilsTag.FallbackChain()
    // ...
}
```

## 禁止写法

```go
import "github.com/lazygophers/utils/language"

// ❌ 公共 API 暴露项目自有包装类型
func Register(tag *language.Tag, key, value string)
func (p *I18n) LocalizeWithLang(tag *language.Tag, key string, args ...any) string
```

---

## 为什么

1. **互操作性**：stdlib `xlanguage.Tag` 是 Go 生态多语言事实标准（`golang.org/x/text/*` 全家都用），第三方库可无缝接入。
2. **值类型零分配**：`xlanguage.Tag` 是 ~16 字节值类型，传参 / 比较无堆分配；`*utils/language.Tag` 是指针，每次构造需 intern map 查询。
3. **避免 leaky abstraction**：`utils/language.Tag` 暴露给上层会绑死实现细节（指针 / 内部缓存语义），换实现就破坏 API。
4. **接口契约对齐**：xerror.Localizer / i18n.I18n / validator.LocaleConfig 全部 stdlib，自动 satisfy。

---

## 边界转换

```go
// xlanguage.Tag → *utils/language.Tag
utilsTag := language.Make(xtag.String())

// *utils/language.Tag → xlanguage.Tag
xtag := utilsTag.Tag()
```

转换发生在**公共 API 入口** 或 **离开公共 API 出口**，不在业务逻辑中部传播。

---

## 适配既有代码

```bash
# 扫描公共 API 中残留的 *language.Tag
grep -rn "\*language\.Tag" --include="*.go" . | grep -v "_test\|/language/"
```

任何命中的公共签名必须改用 `xlanguage.Tag`。

---

## 关联约定

- [Error Handling Style](./error-handling-style.md)
- [No Anonymous Struct](./no-anonymous-struct.md)
- [Locale File Convention](./locale-file-convention.md)
