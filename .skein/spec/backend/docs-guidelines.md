# Docs Guidelines

> Rspress v2 文档站维护规范和已知坑。

---

## 位置

`docs/` 目录。三语言对称：zh-CN（默认）、zh-TW、en。

## 命令

```bash
cd docs && npm run dev     # 预览
cd docs && npm run build   # 构建（必须从 docs/ 目录运行）
cd docs && npm run preview # 预览构建产物
```

## Known Pitfalls

### 1. 死链检查器 bug

`markdown.link.checkDeadLinks: false` **无效**。`routeService` 在 MDX 编译阶段不含路由，所有内部链接都报死链。

**修复**：`pnpm patch @rspress/core`，在 `link.js` 中将 `if (shouldCheckDeadLinks && routeService)` 改为 `if (false)`。

patch 文件：`docs/patches/@rspress__core@2.0.13.patch`（注册在 `docs/pnpm-workspace.yaml`）。升级 `@rspress/core` 后需确认上游是否修复。

### 2. Shiki code block 无 CSS class

Rspress/Shiki 渲染 ` ```mermaid ` 为 `<pre data-lang="mermaid"><code>`，**不给 `<code>` 加 `language-mermaid` 类**。任何自定义代码块渲染组件必须用 `pre[data-lang="<lang>"]` 选择器。

### 3. SSG 从根目录构建失败

从项目根运行 `rspress build` 会把 `rspress.config.ts` 当页面路由。**必须 `cd docs && npx rspress build`**。

### 4. Node.js built-in 不能进客户端 bundle

`rspress.config.ts` 会被打入客户端 bundle。不能直接 `import 'path'` 或 `import 'url'`。需要 Node.js API 的插件定义放到独立文件（参考 `docs/plugins/mermaid.ts`）。

## 文件结构

```
docs/
├── rspress.config.ts        # 主配置（sidebar、nav、i18n）
├── plugins/mermaid.ts       # Mermaid 插件（隔离 Node.js API）
├── components/Mermaid.tsx   # Mermaid 运行时渲染组件
├── patches/                 # pnpm patch 文件
├── pnpm-workspace.yaml      # patchedDependencies 注册
├── zh-CN/                   # 简体中文（默认语言）
├── zh-TW/                   # 繁体中文
├── en/                      # 英文
└── doc_build/               # 构建产物（gitignore）
```

## 新增模块 Checklist

1. 三语言各补 `.md`，路径结构一致
2. `rspress.config.ts` 的 `moduleGroups` 加侧边栏入口
3. 构建验证：`cd docs && npx rspress build`
