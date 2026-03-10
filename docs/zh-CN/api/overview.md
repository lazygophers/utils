---
title: API 概览
---

# API 概览

这份页面不是完整的 API 列表，而是帮助你更快找到正确入口。真正的函数签名与完整导出项，请结合 `go doc` 或 pkg.go.dev 使用。

## 根包：少量通用辅助

根包 `github.com/lazygophers/utils` 适合承载少量跨主题通用能力：

- `Must(value, err)`：在初始化或必须成功的路径上快速失败。
- `MustSuccess(err)`：只关心错误是否为空。
- `MustOk(value, ok)`：处理 `(T, bool)` 风格返回值。
- `Ignore(value, _)`：显式忽略不需要的第二返回值。
- `Scan(src, dst)`：把数据库字段中的 JSON/文本扫描到结构体。
- `Value(v)`：把结构体转换为数据库可写入值。

## 常用子包入口

| 主题 | 建议入口 | 说明 |
| --- | --- | --- |
| 配置加载 | `config.LoadConfig` | 自动探测文件、支持多格式、可结合校验使用 |
| 数据验证 | `validator.Struct` | 针对结构体执行规则校验 |
| 类型与集合辅助 | `candy` | 转换、集合操作、映射与切片辅助 |
| JSON 编解码 | `json.Marshal` / `json.Unmarshal` | 对标准库做工程化封装 |
| 时间与日历 | `xtime.NowCalendar` | 日历、农历、节气与排班相关能力 |
| 默认值填充 | `defaults.SetDefaults` | 按标签或选项补齐默认值 |
| 退出回调 | `atexit.Register` | 进程退出时执行清理逻辑 |
| URL 规范化 | `urlx.SortQuery` | 统一查询参数顺序 |
| 性能采样 | `pyroscope.Load` | 接入 Pyroscope |

## 如何查找完整 API

### 方式一：按本地文档主题查找

- 想先判断是否适合你的场景：从 [模块总览](/modules/overview) 或分类页进入。
- 想看某个包的定位与边界：进入对应模块页。
- 想直接跳到源码：从模块页再跳到仓库目录或 pkg.go.dev。

### 方式二：按包名查找

- 根包：`github.com/lazygophers/utils`
- 子包示例：`github.com/lazygophers/utils/config`
- 子包示例：`github.com/lazygophers/utils/xtime`
- 缓存策略：`github.com/lazygophers/utils/cache/...`

## 阅读 API 时的注意点

1. 根包不是“所有能力的统一入口”，大多数主题能力都在子包里。
2. 某些页面会给出示例，但示例只覆盖最常见用法，不代表完整导出面。
3. 缓存、排班、退出处理这类主题存在明显约束，先看模块页再查签名会更稳。

## 下一步

- 想先理解能力范围：看 [模块总览](/modules/overview)
- 想直接开始调用：看 [快速开始](/guide/getting-started)
- 想进入包级文档：访问 https://pkg.go.dev/github.com/lazygophers/utils
