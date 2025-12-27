# Rspress 文档系统

本项目使用 [Rspress](https://rspress.rs/) 构建多语言文档系统。

## 🌍 支持的语言

- **简体中文** (zh-CN) - 默认语言
- **繁體中文** (zh-TW)
- **English** (en)

## 📁 文档结构

```
docs/
├── package.json         # NPM 配置文件
├── rspress.config.ts    # Rspress 配置文件
├── tsconfig.json        # TypeScript 配置文件
├── README.md           # 文档说明
├── zh-CN/              # 简体中文文档
│   ├── index.md        # 首页
│   ├── guide/          # 指南
│   ├── modules/        # 模块文档
│   └── api/           # API 文档
├── zh-TW/              # 繁体中文文档
│   └── ...
└── en/                # 英文文档
    └── ...
```

## 🚀 本地开发

### 安装依赖

```bash
cd docs
npm install
```

### 启动开发服务器

```bash
cd docs
npm run dev
```

文档将在 `http://localhost:3000` 启动。

### 构建文档

```bash
cd docs
npm run build
```

构建产物将输出到 `docs/doc_build` 目录。

### 预览构建结果

```bash
cd docs
npm run preview
```

## 📝 添加新文档

1. 在对应的语言目录下创建 `.md` 文件
2. 文件顶部添加 frontmatter：

```markdown
---
title: 文档标题
---
```

3. 在导航配置中添加链接（在 `rspress.config.ts` 中）

## 🌐 语言切换

文档支持语言切换功能，用户可以在页面右上角选择语言。

## 🚀 部署到 GitHub Pages

文档已配置自动部署到 GitHub Pages：

1. 推送代码到 `main` 或 `master` 分支
2. GitHub Actions 将自动构建并部署文档
3. 部署完成后，文档将可通过 GitHub Pages 访问

### 手动触发部署

在 GitHub Actions 页面，选择 "Deploy Rspress site to Pages" workflow，点击 "Run workflow" 按钮手动触发部署。

## 🔧 配置文件

### rspress.config.ts

Rspress 配置文件，包含：
- 文档根目录
- 多语言配置
- 主题配置
- 导航栏配置

### package.json

包含文档系统的脚本和依赖。

## 📖 相关资源

- [Rspress 官方文档](https://rspress.rs/)
- [Rspress GitHub](https://github.com/web-infra-dev/rspress)
- [LazyGophers Utils GitHub](https://github.com/lazygophers/utils)
