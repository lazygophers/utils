---
title: must
---

# must

`github.com/lazygophers/utils` 根包里的 `Must`、`MustSuccess`、`MustOk`、`Ignore` 适合放在**初始化阶段**或**必须成功才能继续**的路径里，用来减少重复的错误判断样板。

## 适合什么场景

- 启动加载配置、建立连接、解析必要资源时，失败就直接终止。
- 处理 `(T, error)`、`error`、`(T, bool)` 这三类常见返回值时，希望写法更紧凑。
- 想把“这里必须成功”表达得更明确，而不是埋在多层 `if err != nil` 里。

## 常用入口

- `utils.Must(value, err)`：返回值并在 `err != nil` 时 panic。
- `utils.MustSuccess(err)`：只检查错误是否为空。
- `utils.MustOk(value, ok)`：处理 `(T, bool)` 风格返回。
- `utils.Ignore(value, _)`：显式忽略不关心的第二返回值。

## 使用建议

- 这组函数更适合**启动、装配、一次性初始化**，不适合普通业务请求路径。
- 如果调用失败后仍希望继续处理或向上返回错误，请直接写显式错误分支。
- 团队代码规范里如果强调“库代码不 panic”，那它通常只该出现在应用入口层。

## 相关文档

- [orm](/modules/core/orm)
- [validator](/modules/core/validator)
- [快速开始](/guide/getting-started)
