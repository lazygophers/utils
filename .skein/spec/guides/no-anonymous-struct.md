# No Anonymous Struct in Variable Declarations

> **强制规范**：任何变量声明都禁止直接使用 anonymous struct 类型；必须先命名 type，再以该类型声明变量。

---

## 禁止写法

```go
// ❌ 包级 var
var registry = struct {
    sync.RWMutex
    m map[string]string
}{m: make(map[string]string)}

// ❌ 局部 := 单值
fl := struct {
    Field string
}{Field: "x"}

// ❌ table-driven 测试（最常见）
tests := []struct {
    name string
    in   int
    want int
}{
    {"a", 1, 2},
}

// ❌ 函数返回 / 参数中的内联 struct 字面量
return struct{ Code int; Msg string }{500, "err"}
```

## 必须写法

```go
// ✅ 包级：先命名 type，再声明 var
type messageStore struct {
    sync.RWMutex
    m map[string]string
}
var registry = messageStore{m: make(map[string]string)}

// ✅ 局部：函数内 type X struct{...} 完全合法
type fieldHolder struct {
    Field string
}
fl := fieldHolder{Field: "x"}

// ✅ table-driven：先类型，再切片
type testCase struct {
    name string
    in   int
    want int
}
tests := []testCase{
    {"a", 1, 2},
}
```

---

## 为什么

1. **可读性**：类型名传达意图（`testCase` / `messageStore` / `fieldHolder`），匿名 struct 强迫读者每次现解构字段语义。
2. **可测性**：命名类型可在测试中独立构造、对比、传参；匿名类型不能跨函数共享。
3. **可演进**：往字段加新成员时，命名类型只改一处声明；匿名 struct 散落各处时每个字面量都要改。
4. **诊断友好**：Go 编译错误对 `testCase` 的提示远比 `struct { name string; in int; ... }` 易读，IDE 类型跳转也工作。
5. **零运行时成本**：命名 type 不增加任何运行时开销，仅是源码层的命名。
6. **统一风格**：消除"有时匿名、有时命名"的判断负担，所有 struct 出现位置形态一致。

---

## 适用范围

- **所有 Go 代码**强制生效：产品代码、测试、benchmark、_test.go。
- **所有 struct 形态**：`var X = struct{}{}` / `:= struct{}{}` / `:= []struct{}{}` / `return struct{}{...}` / 函数参数内联 struct 字面量 / map value 类型为 anonymous struct（`map[K]struct{...}`）等等。
- **唯一例外**：`struct{}` 空 struct（如 `chan struct{}` / `map[K]struct{}` 集合语义），零字段无字段语义负担，允许保留。

---

## 命名建议

- **函数内单用**：直接 `testCase` / `entry` / `fixture` / `fieldHolder`。
- **多函数共用**：提到包级，按文件主题命名（`hashCase` / `parseCase` / `localeCase`）。
- **包级状态**：按角色命名（`messageStore` / `pluckIndexCache`），避免 `T` / `Type` 等无信息后缀。
- 类型名首字母小写（除非真要导出），避免污染包公共 API。

---

## 适配既有代码

扫描违规：

```bash
grep -rnE ":= \[\]struct \{|:= struct \{|^[[:space:]]*var [a-zA-Z_]+ = struct \{|^[[:space:]]*var [a-zA-Z_]+ struct \{" --include="*.go" .
```

批量改：脚本逐个识别 `<var> := []struct{...}{` 块，提取字段集合为函数局部 `type <Name> struct{...}`，再把字面量改成 `[]<Name>{...}`。注意保留原数据条目不变，只改类型声明形态。

---

## 关联约定

- [Error Handling Style](./error-handling-style.md) — error 接收与判断必须分两行
- 项目根 `CLAUDE.md` 中"禁止防御性编程"、"命名 Id 非 ID"等编码约定与本规则配合执行
