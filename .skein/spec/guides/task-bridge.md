# Task Bridge: Trellis task.py ↔ Claude built-in TaskCreate

> **Purpose**: 消除系统级 `<system-reminder>`"必须用 built-in TaskCreate"的噪声，让两套 task 系统协作而非互相干扰。

---

## 两套 Task 系统

| 系统 | 存储位置 | 生命周期 | 用途 |
|------|---------|---------|------|
| **Trellis `task.py`** | `.trellis/tasks/<MM-DD-slug>/` | 持久, 跨会话, 跨开发者 | source of truth: PRD / design / implement / research / 子任务树 |
| **Claude built-in TaskCreate** | 会话内 ephemeral | 单会话, 不持久 | 进度可视化 + 静默系统级 reminder |

它们不是替代关系——Trellis 解决"跨会话 / 跨人 / 持久 / 协作"，built-in TaskCreate 解决"本会话进度展示 + reminder 消音"。

---

## Bridge 协议

主线 AI 在以下时机镜像 Trellis task 到 built-in TaskCreate：

### 进入 task（`task.py start` 之后）

`task.py start <name>` 跑完时 stdout 末尾会输出一行：

```
[trellis-bridge] Mirror to Claude TaskCreate (manual): "Active task: <task title>"
```

看到这一行后，主线 AI **必须**立即用 built-in TaskCreate 在会话 todo 里建一条 `in_progress` 条目，标题与 Trellis task title 一致。

> 注：built-in TaskCreate API 不能由项目脚本调用，必须由 AI 主动执行；hook/stdout 只能 hint。

### 离开 task（`task.py archive` 之后）

任务归档完成时，主线 AI 用 built-in TaskUpdate 把对应 todo 标 `completed`。

---

## 为什么不全部走 built-in TaskCreate

- 会话结束即丢失，无法跨人协作
- 没有 PRD / design / implement 字段
- 没有子任务树、agent 绑定、acceptance criteria
- 无法被 CI / hooks / lint 工具消费

## 为什么不全部走 task.py

- 系统级 prompt 强制 reminder 检测 built-in TaskCreate 调用
- 没有进度气泡 UI
- AI 看不到自己的 todo 列表

---

## 实操清单（主线 AI）

- [ ] `task.py start` 后看到 `[trellis-bridge]` 行：调 TaskCreate 镜像
- [ ] 跨子任务切换：TaskUpdate 旧 todo → completed, TaskCreate 新 todo
- [ ] `task.py archive` 后：TaskUpdate 对应 todo → completed
- [ ] 不要反向同步——TaskCreate 不是 source of truth，不能用它替代 PRD

---

## 相关

- [.trellis/workflow.md](../../workflow.md) Phase 1.4 — Activate task
- `~/.claude/CLAUDE.md` — "必须使用 built-in Task 模块" 全局约束
