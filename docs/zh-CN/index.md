---
pageType: home
hero:
  name: LazyGophers Utils
  text: 面向 Go 工程实践的实用工具库
  tagline: 用更稳定的基础能力，减少重复造轮子与样板代码
  actions:
    - theme: brand
      text: 快速开始
      link: /guide/getting-started
    - theme: alt
      text: 浏览模块
      link: /modules/overview
features:
  - title: 明确分组
    details: 以核心工具、数据处理、缓存、时间、系统、网络、并发、开发测试八个主题组织能力，查找入口更直接。
    icon: 🧭
  - title: 紧贴源码
    details: 中文文档以仓库中的真实导出能力、包说明和局部约束为依据，不再依赖未经证实的营销描述。
    icon: 🔎
  - title: 从场景出发
    details: 每个页面优先回答“什么时候该用它、先看什么、有哪些边界”，减少上手成本。
    icon: 🚀
  - title: 多语言文档站
    details: 同一套站点同时承载简体中文、繁體中文与 English，适合团队协作与对外发布。
    icon: 🌍
---

## 这套文档解决什么问题

LazyGophers Utils 是一个覆盖面很广的 Go 工具库。真正的难点不在于“有没有功能”，而在于你是否能快速定位到合适的包、理解它的适用边界，并在工程里安全落地。

这份中文文档因此围绕三个问题组织：

1. **先用哪个包**：按场景而不是按目录名称理解模块。
2. **能直接拿来做什么**：从真实入口函数出发，而不是抽象宣传语。
3. **有哪些边界**：哪些能力适合初始化阶段，哪些适合业务路径，哪些要注意平台或并发限制。

## 推荐阅读路径

### 第一次接触仓库

- 从 [快速开始](/guide/getting-started) 建立导入与调用方式的整体印象。
- 从 [模块总览](/modules/overview) 判断你要落到哪个主题分组。
- 进入对应分类页，再看具体模块页与相关约束。

### 想解决具体问题

| 目标 | 建议先看 |
| --- | --- |
| 初始化阶段快速失败 | [must](/modules/core/must) |
| JSON 与数据库字段映射 | [orm](/modules/core/orm) |
| 配置加载与格式兼容 | [config](/modules/system/config) |
| 农历、节气与排班计算 | [xtime](/modules/time/xtime) |
| 选择缓存策略 | [缓存策略](/modules/cache/) |
| 默认值、随机数据、测试支撑 | [开发与测试](/modules/dev/) |
| 按国家维度造假数据（姓名 / 证件 / 电话 / 地址） | [fake](/modules/dev/fake) |

## 文档使用方式

- 想看 API 面：转到 [API 概览](/api/overview)，再跳转到 pkg.go.dev。
- 想找模块入口：从 [模块总览](/modules/overview) 或左侧边栏进入。
- 想判断是否适合生产使用：优先看每个页面的“适用场景”和“使用建议”。

## 项目事实速览

- 仓库以根级包组织能力，没有传统的 `pkg/` 或 `internal/` 分层文档入口。
- 根入口适合放少量通用辅助，更多能力位于子包中。
- 缓存、时间与退出处理等主题存在专门约束，文档会直接说明这些局部规则。
