---
title: runtime
---

# runtime

`runtime` 包提供**运行时辅助能力**，重点在 panic 捕获、堆栈信息以及可执行文件、工作目录、用户目录等路径定位。

## 适合什么场景

- 想在应用入口统一处理 panic 和错误现场。
- 需要获取可执行文件目录、当前工作目录、用户目录等运行时路径。
- 想把一些与进程环境相关的辅助逻辑从业务代码里抽出来。

## 使用建议

- 这类能力更适合基础设施层或入口层，不建议散落在普通业务函数里。
- 涉及平台差异时，先确认对应文件是否带有 build tag 或系统分支逻辑。
- 如果你处理的是退出阶段清理，请继续看 [atexit](/modules/system/atexit)。

## 相关文档

- [osx](/modules/system/osx)
- [app](/modules/system/app)
- [atexit](/modules/system/atexit)
