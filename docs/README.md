# Rspress 文档系统

本目录使用 **Rspress v2** 维护 LazyGophers Utils 的多语言文档站点。

## 支持语言

- 简体中文（`zh-CN`，默认）
- 繁體中文（`zh-TW`）
- English（`en`）

## 目录结构

```text
docs/
├── package.json
├── rspress.config.ts
├── tsconfig.json
├── README.md
├── zh-CN/
├── zh-TW/
└── en/
```

其中 `zh-CN/` 是本次重构的中文文档主目录，首页、指南、模块总览、分类索引和模块页都从这里生成。

## 本地开发

```bash
cd docs
npm install
npm run dev
```

默认开发地址通常为 `http://localhost:3000`。

## 构建与预览

```bash
cd docs
npm run build
npm run preview
```

构建产物默认输出到 Rspress 的静态目录中，可直接用于 GitHub Pages 或其他静态托管方案。

## 维护约定

1. 优先以源码与包级说明为准，不保留无法证明的性能数字或命中率。
2. 新增页面时，先补侧边栏入口，再补正文内容。
3. 中文文档优先说明“适合做什么、从哪里开始、有哪些约束”，而不是堆砌 API 列表。
4. 复杂模块应使用统一模板：模块定位、适用场景、核心入口、使用建议、相关文档。

## 相关链接

- Rspress 官方文档：https://rspress.rs/
- 项目仓库：https://github.com/lazygophers/utils
- 在线 API 文档：https://pkg.go.dev/github.com/lazygophers/utils
