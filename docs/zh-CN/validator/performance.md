---
title: 性能与最佳实践 - Validator
---

# 性能与最佳实践

## 性能

Validator 内置多种算法变体，通过基准测试选择最优实现：

| 验证项 | 优化方向 |
|--------|---------|
| 银行卡号 | 12 种 Luhn 算法变体（字节级、查表、分支消除等） |
| 身份证号 | 9 种校验算法变体 |
| 消息格式化 | 5 种字符串模板实现（原始/Builder/编译/字节切片/无 fmt） |
| 邮箱 | 多种正则/非正则方案对比 |

运行基准测试：

```bash
go test -bench=. ./validator/
```

## 最佳实践

**标签设计**：每个字段用 2-3 个标签组合，而非一个万能标签。

```go
// ✅ 推荐
Email string `validate:"required,email"`
Name  string `validate:"required,min=2,max=50"`

// ❌ 太粗
Email string `validate:"required"`
Name  string `validate:"required"`
```

**验证时机**：在数据进入系统的第一站验证（配置加载、请求解析），不要等到业务深处。

**错误处理**：用 `ValidationErrors` 类型断言，逐字段生成用户友好的消息。

**复用实例**：用 `Default()` 或缓存 `New()` 实例，避免每次创建。
