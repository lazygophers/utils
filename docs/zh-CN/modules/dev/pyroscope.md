---
title: pyroscope
---

# pyroscope

`pyroscope` 用于把应用接入 **Pyroscope 性能采样**。如果你要把运行中的程序接入持续剖析，而不是只做本地一次性 benchmark，这个包更有意义。

## 适合什么场景

- 想在服务运行期持续采集性能画像。
- 需要把应用的 profiling 接入流程收口到统一入口。
- 想结合版本、实例、环境标签进行观测。

## 常用入口

- `pyroscope.Load`

## 使用建议

- 先确认采样平台、标签策略和上报地址，再做接入代码。
- profiling 是观测能力，不应和业务逻辑耦合得过深。
- 应用元信息通常要配合 [app](/modules/system/app) 一起整理。

## 相关文档

- [runtime](/modules/system/runtime)
- [app](/modules/system/app)
- [开发与测试](/modules/dev/)
