# Multi-Task Worktree

> **Purpose**: 一个 session 内并行最多 2 个 task, 每个 task 拥有独立的 git worktree + branch, 主仓只读, 互不污染。

---

## Why worktree?

session 内并行 task 时, 共享 working tree 会导致:
- 两个 task 的 dirty 改动互相覆盖, commit 时分不清归属
- lint / test 在主仓跑, 看不到对方 worktree 的改动
- 用户手切 branch 时另一 task 进度全乱

worktree 把每个 task 隔离到 `.worktrees/<slug>/`, 各自有独立 working tree + branch。主仓回到只读状态 (只读 git index, 不接受 edit)。

---

## 模型

```
┌──────────────────────────────────────────────────────────────────┐
│ Claude Code session (sid: claude_<uuid>)                         │
│                                                                  │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │ .trellis/.runtime/sessions/<sid>-active.json            │    │
│  │  active: [                                              │    │
│  │    {slot: 1, task: foo, worktree: .worktrees/foo, ...} │    │
│  │    {slot: 2, task: bar, worktree: .worktrees/bar, ...} │    │
│  │  ]                                                      │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                  │
│  Main session cwd: <repo root>                                   │
│                    └─ orchestrates, dispatches agents            │
│                                                                  │
│  ┌──────────────┐  ┌──────────────┐                              │
│  │ Sub-agent 1  │  │ Sub-agent 2  │  ← cwd 注入 .worktrees/*    │
│  │ cwd: foo wt  │  │ cwd: bar wt  │     by inject-subagent-ctx  │
│  └──────────────┘  └──────────────┘                              │
└──────────────────────────────────────────────────────────────────┘

.worktrees/foo/  ← git worktree, branch task/05-21-foo
.worktrees/bar/  ← git worktree, branch task/05-21-bar
<repo root>      ← master branch (或当前 HEAD)
```

active state 文件 schema 见 `.trellis/scripts/common/active_task.py` (`ActiveState` / `ActiveSlot`)。

---

## 命令 cheatsheet

```bash
# 创建 worktree + 激活 (默认行为, 自动 branch task/<slug>, 从当前 HEAD 切)
python3 ./.trellis/scripts/task.py start <slug>

# legacy 路径: 不建 worktree, 直接在主仓 working tree 改 (单 task 场景)
python3 ./.trellis/scripts/task.py start <slug> --skip-worktree

# 列当前 session 的全部 active slot (含 worktree / branch)
python3 ./.trellis/scripts/task.py current --source

# archive: 自动清 worktree + 删 branch (未 merge 时 warn 留 branch)
python3 ./.trellis/scripts/task.py archive <slug>

# archive 但 worktree 内还有未 commit 改动 → 用 --force-dirty 跳检测 (会丢)
python3 ./.trellis/scripts/task.py archive <slug> --force-dirty

# 仅 archive metadata, 保留 worktree (调试 / 后续手动 PR 用)
python3 ./.trellis/scripts/task.py archive <slug> --keep-worktree
```

并发约束: active slot ≤ 2。第 3 个 `task.py start` 会拒绝, 提示先 archive 或 finish 一个。

---

## sub-agent 自动 cwd 注入

sub-agent dispatch protocol: dispatch prompt 以 `Active task: <task path from task.py current>` 开头 → hook 注入 `## Working directory` 段:

```
## Working directory

cd /abs/path/to/.worktrees/<slug>

All edits / make commands run inside this worktree, not repo root.
Repo root is read-only for this agent.
```

skip_worktree=true 或 slot 无 worktree (legacy) 时不注入, sub-agent 默认 cwd=主仓。

---

## 故障排查

| 现象 | 原因 | 解决 |
|------|------|------|
| `git worktree add` 失败: `.worktrees/<slug>` 已存在 | 上次 archive 没清干净 / 手动残留 | `git worktree prune` + `rm -rf .worktrees/<slug>` |
| `git worktree add` 失败: branch `task/<slug>` 已存在 | 同名 task 之前跑过没删 branch | `git branch -d task/<slug>` (未 merge 用 `-D`) |
| archive 报 dirty | worktree 内还有未 commit 改动 | 进 worktree commit, 或 `--force-dirty` 弃改动 |
| `branch -d` 失败 (unmerged) | worktree commit 还没 merge 回 base | warn 不阻断, 后续手 PR / merge, 或 `-D` 强删 |
| sub-agent 仍在主仓 cwd 编辑 | dispatch prompt 没含 `Active task: <path>` 行, 或 slot.worktree 空 | 检查 prompt 模板, 跑 `task.py current` 确认 slot 字段 |

---

## 协作模式

- 2 个 task **完全独立**: branch / working tree / commit 历史都隔离, 互不可见对方 in-flight 改动
- 跨 task 共享代码: 必须先在一个 task 内 commit + push (或 merge 回 base), 另一个 task rebase / merge base 拿到
- 主仓 (`<repo root>`) 在 session 期间是 **只读视图**: 主线 AI orchestrates / dispatches, 不直接 edit / commit
- Stop hook 跑 lint **在主仓 cwd** — 当前不感知 worktree 改动 (已知 caveat, 后续 hook 优化)

---

## Out of scope (本目录不涉及)

- worktree branch 自动 merge 回 base
- 跨 session 复用同一个 worktree
- worktree 自动 push 到 remote
- worktree 内 Stop hook lint 集成
