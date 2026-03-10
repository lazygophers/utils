---
title: app
---

# app

`app` 包关注的是**应用自身的元信息与初始化上下文**，例如名称、版本、构建信息和与应用身份有关的公共数据。

## 适合什么场景

- 想在程序内统一读取版本、构建时间、提交信息等元数据。
- 需要为日志、命令行输出或诊断接口提供统一的应用身份信息。
- 想把应用层的公共上下文集中管理，而不是到处散落常量。

## 使用建议

- 构建信息最好由 CI 或发布流程注入，避免手工维护。
- `app` 更适合描述“程序是谁”，配置、路径和退出逻辑仍应交给对应主题包处理。
- 如果你要接入性能采样或诊断输出，可联动看 [pyroscope](/modules/dev/pyroscope)。

## 相关文档

- [runtime](/modules/system/runtime)
- [config](/modules/system/config)
- [pyroscope](/modules/dev/pyroscope)
