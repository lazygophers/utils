---
title: config
---

# config

`config` 包负责**配置加载、格式解析与初始化阶段的配置装配**。它适合放在程序启动时，把文件内容读成结构体，再决定是否进入后续流程。

## 适合什么场景

- 项目需要从文件加载配置，并兼容多种格式。
- 想统一配置搜索顺序和解析入口。
- 希望把“加载”和“校验”两个阶段拆开控制。

## 常用入口

- `config.LoadConfig`：加载并执行校验
- `config.LoadConfigSkipValidate`：只加载，不做校验
- `config.LoadConfigWithInheritance`：多文件配置继承（后加载覆盖先加载）
- `config.LoadConfigByEnvironment`：根据 ENV 自动加载配置文件
- `config.SetConfig`：设置全局配置对象
- `config.RegisterParser`：注册额外解析器

## 特性

- **配置继承**：支持多个配置文件的优先级覆盖
  - 调用顺序：`LoadConfigWithInheritance(cfg, "base.json", "env.json", "local.json")`
  - 后加载的配置完全覆盖先加载的（而非只覆盖零值）
- **HCL 完整支持**：支持 slice、array、map 类型的读写

## 使用建议

- 对外可变配置与启动必需配置，最好分层管理，不要都塞进一个超大结构体。
- 如果配置必须合法才能启动，优先使用 `LoadConfig` 并结合 [validator](/modules/core/validator)。
- 文档只保留“支持多格式”这一事实，不再宣传无法在当前仓库中验证的额外能力。

## 相关文档

- [validator](/modules/core/validator)
- [json](/modules/data/json)
- [app](/modules/system/app)
