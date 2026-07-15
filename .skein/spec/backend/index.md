# Backend Development Guidelines

> Conventions extracted from the LazyGophers Utils codebase.

---

## Overview

Go 工具库（`github.com/lazygophers/utils`），Go 1.26.2。每个子目录是独立 Go package。根包 `utils` 提供断言/验证/ORM 辅助，20+ 子包覆盖缓存、并发、加解密、配置、时间等场景。

## Guidelines Index

| Guide | Description |
|-------|-------------|
| [Directory Structure](./directory-structure.md) | Package organization, file naming, build tags |
| [Database Helpers](./database-guidelines.md) | Value/Scan ORM helpers (not an ORM framework) |
| [Error Handling](./error-handling.md) | Must* fail-fast, panic patterns, log integration |
| [Logging](./logging-guidelines.md) | github.com/lazygophers/log usage, levels, trace |
| [Quality Guidelines](./quality-guidelines.md) | Testing, lint config, coverage, benchmarks |
| [Docs Guidelines](./docs-guidelines.md) | Rspress 文档站维护、已知坑、新增模块流程 |
| [fake/ UA 模板真实性规则](./fake-ua-templates.md) | UA 模板冻结值、平台 token、禁组合、版本池 bump 约定 |

---

## How to Use

Sub-agents loading these specs: read all files listed above before writing code. These reflect actual project patterns, not ideals.
