---
title: 內建規則 - Validator
---

# 內建驗證標籤

Validator 內建豐富的驗證標籤，覆蓋常見校驗場景。

## 基礎規則

| 標籤 | 說明 | 範例 |
|------|------|------|
| `required` | 非零值 | `validate:"required"` |
| `email` | 郵箱格式 | `validate:"email"` |
| `url` | URL 格式 | `validate:"url"` |
| `alpha` | 純字母 | `validate:"alpha"` |
| `alphanum` | 字母 + 數字 | `validate:"alphanum"` |
| `json` | JSON 格式 | `validate:"json"` |
| `uuid` | UUID 格式 | `validate:"uuid"` |

## 數值/長度

| 標籤 | 說明 | 範例 |
|------|------|------|
| `min=N` | 最小值/長度 | `validate:"min=3"` |
| `max=N` | 最大值/長度 | `validate:"max=100"` |
| `len=N` | 精確長度 | `validate:"len=11"` |
| `minlen=N` | 最小長度 | `validate:"minlen=6"` |
| `maxlen=N` | 最大長度 | `validate:"maxlen=20"` |
| `range=min,max` | 數值範圍 | `validate:"range=0.0,5.0"` |

## 集合/模式

| 標籤 | 說明 | 範例 |
|------|------|------|
| `in=v1,v2,...` | 列舉值 | `validate:"in=male,female"` |
| `notin=v1,v2,...` | 排除列舉 | `validate:"notin=admin,root"` |
| `pattern=regex` | 正則匹配 | `validate:"pattern=^[A-Z]"` |
| `containspecial` | 含特殊字元 | `validate:"containspecial"` |
| `strong_password` | 強密碼 | `validate:"strong_password"` |

## 邏輯組合

```go
// And — 全部滿足
combined := validator.And(
    validator.Required(),
    validator.MinLength(5),
)

// Or — 滿足其一
either := validator.Or(
    validator.Email(),
    validator.Pattern(`^\d+$`),
)

// Not — 取反
negated := validator.Not(validator.In("a", "b"))
```
