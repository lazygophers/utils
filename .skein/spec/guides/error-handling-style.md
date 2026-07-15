# Error Handling Style

> **强制规范**：错误的接收与判断必须分两行，禁止内联到 `if` 语句中。

---

## 禁止写法

```go
// ❌ 内联 := + if 判断（最常见的违规）
if err := doSomething(); err != nil {
    return err
}

// ❌ 同样禁止 = 形式的内联
if err = doSomething(); err != nil {
    return err
}

// ❌ 复合条件也不行
if err := h.Unmarshal(buf, &v); err == nil || err.Error() != "boom" {
    ...
}
```

## 必须写法

```go
// ✅ 拆两行：接住 → 判断
err := doSomething()
if err != nil {
    return err
}

// ✅ 同一作用域已有 err，复用赋值
err = doAnother()
if err != nil {
    return err
}

// ✅ 同一函数多次 error 处理时，前述变量已存在
result, err := doSomething()
if err != nil {
    return err
}
result2, err := doAnother()  // err 是新增就用 :=
if err != nil {
    return err
}
```

---

## 为什么

1. **可调试性**：错误变量在作用域可见，IDE 断点 / `fmt.Printf("%+v", err)` 直接可用，无需重排。
2. **可读性**：业务动作（产生 error 的调用）与控制流（判错）分离，扫读时眼睛只在 `if err != nil` 上停留。
3. **一致性**：所有 error 处理形态一致，没有"有些内联、有些分行"的混乱。
4. **可演进**：后续要在 error 处理逻辑里加 log / wrap / metadata，单独行有空间扩展，内联形态要先拆。
5. **避免 shadow**：内联 `if err := ...` 的 `err` 只活在 if 块内，外层若要传递必须重新声明，容易引入 shadow bug；分两行后作用域统一。

---

## 适用范围

- **本规则对所有 Go 模块强制生效**，包括产品代码与测试代码。
- 不限于 `error`：任何在 if 中接住返回值并立即判断的形态都应避免，但 error 是最高优先级。
- 例外：仅当返回值**纯布尔类型**且语义就是"判断"（如 `if ok := m[k]; ok`）时可酌情保留；但若 `ok` 后续还要使用，仍建议拆开。

---

## 适配既有代码

发现存量违规：用 `grep -rn "if err :=" --include="*.go"` 扫描。

批量改写：

```bash
perl -i -pe 's/^(\s*)if err := (.+?); err (==|!=) nil \{$/$1err := $2\n$1if err $3 nil {/g' file.go
```

注意：批量改写后会出现 `no new variables on left side of :=` 报错（同作用域已有 `err`），逐个改为 `=` 即可。

---

## 关联约定

- [Code Reuse Thinking Guide](./code-reuse-thinking-guide.md)
- 项目根 `CLAUDE.md` 中的 "禁止防御性编程" 与本规则配合：error 只在系统边界判断，但**判断时必须分两行**。
