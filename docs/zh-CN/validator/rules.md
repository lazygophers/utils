---
title: 内置规则 - Validator
---

# 内置验证标签

Validator 内置丰富的验证标签，覆盖常见校验场景。

## 基础规则

| 标签 | 说明 | 示例 |
|------|------|------|
| `required` | 非零值 | `validate:"required"` |
| `email` | 邮箱格式 | `validate:"email"` |
| `url` | URL 格式 | `validate:"url"` |
| `alpha` | 纯字母 | `validate:"alpha"` |
| `alphanum` | 字母 + 数字 | `validate:"alphanum"` |
| `json` | JSON 格式 | `validate:"json"` |
| `uuid` | UUID 格式 | `validate:"uuid"` |

## 数值/长度

| 标签 | 说明 | 示例 |
|------|------|------|
| `min=N` | 最小值/长度 | `validate:"min=3"` |
| `max=N` | 最大值/长度 | `validate:"max=100"` |
| `len=N` | 精确长度 | `validate:"len=11"` |
| `minlen=N` | 最小长度 | `validate:"minlen=6"` |
| `maxlen=N` | 最大长度 | `validate:"maxlen=20"` |
| `range=min,max` | 数值范围 | `validate:"range=0.0,5.0"` |

## 集合/模式

| 标签 | 说明 | 示例 |
|------|------|------|
| `in=v1,v2,...` | 枚举值 | `validate:"in=male,female"` |
| `notin=v1,v2,...` | 排除枚举 | `validate:"notin=admin,root"` |
| `pattern=regex` | 正则匹配 | `validate:"pattern=^[A-Z]"` |
| `containspecial` | 含特殊字符 | `validate:"containspecial"` |
| `strong_password` | 强密码 | `validate:"strong_password"` |

## 逻辑组合

```go
// And — 全部满足
combined := validator.And(
    validator.Required(),
    validator.MinLength(5),
)

// Or — 满足其一
either := validator.Or(
    validator.Email(),
    validator.Pattern(`^\d+$`),
)

// Not — 取反
negated := validator.Not(validator.In("a", "b"))
```
