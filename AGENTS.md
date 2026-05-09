# LazyGophers Utils - Project Knowledge Base

**Generated:** 2026-03-09
**Scope:** repo root

## OVERVIEW
- Go 1.25 utility library with 20+ public root-level packages.
- Flat layout: no `cmd/`, `pkg/`, or `internal/`; packages live directly under repo root.
- Main themes: generics, performance helpers, multilingual data/validation, concurrency, cache algorithms.

## DOC BOUNDARIES
- Read this file first for repo-wide rules.
- Read [cache/AGENTS.md](cache/AGENTS.md) for cache algorithm selection and trade-offs.
- Read [atexit/AGENTS.md](atexit/AGENTS.md) for shutdown callback behavior and platform-specific exit handling.
- Read [xtime/AGENTS.md](xtime/AGENTS.md) for lunar calendar, solar term, and work-schedule time logic.
- Packages without a local `AGENTS.md` inherit this root guidance.

## REPO MAP
- Root files: `must.go`, `orm.go`, `validate.go`.
- Data + conversion: `candy/`, `json/`, `stringx/`, `anyx/`, `defaults/`.
- Time + schedules: `xtime/`.
- System + lifecycle: `runtime/`, `atexit/`, `app/`, `osx/`, `config/`.
- Reliability + concurrency: `cache/`, `routine/`, `wait/`, `hystrix/`, `singledo/`, `event/`.
- Security + network: `cryptox/`, `pgp/`, `network/`, `urlx/`.
- Test/support data: `fake/`, `randx/`, `pyroscope/`, `validator/`.

## WHERE TO LOOK
| Task | Location | Notes |
|---|---|---|
| fail-fast helpers | `must.go` | `Must`, `MustSuccess`, `MustOk` |
| DB JSON scan/value | `orm.go` | wraps JSON encode/decode for DB fields |
| root validation entry | `validate.go`, `validator/` | root helper delegates to validator package |
| cache policy choice | `cache/` | algorithm-heavy; use child doc |
| graceful shutdown | `atexit/` | build-tagged platform implementations |
| advanced time logic | `xtime/` | lunar calendar, solar terms, 007/955/996 subpackages |
| fake data generation | `fake/` | locale-specific embedded datasets and faker defaults |
| runtime helpers | `runtime/` | panic capture, stack dumping, executable/user dir helpers |
| i18n validation engine | `validator/` | locale files + custom validation engine |

## PROJECT CONVENTIONS
- Prefer existing root-level package structure; do not introduce deep hierarchy casually.
- Code is primarily English; docs should support English and Simplified Chinese.
- Many packages keep `llms.txt`; treat them as hints, not source of truth.
- Use generics where the package already does.
- Library code should return errors explicitly; only `Must*` helpers are expected to panic.
- This repo prefers `github.com/lazygophers/log` for logging.
- Tests are typically English-named and often split into normal tests plus `*_coverage_test.go` edge coverage files.

## PACKAGE-SPECIFIC PATTERNS
- `cache/`: multiple algorithms with different eviction semantics; avoid generic advice here.
- `fake/`: heavy locale/embed surface area; verify actual dataset loaders before changing behavior.
- `validator/`: custom engine + per-locale messages; check field-name strategy before editing messages.
- `runtime/`: includes OS-specific signal/path helpers; distinguish generic helpers from platform files.
- `atexit/`: callback execution semantics vary by OS family.
- `xtime/`: domain rules are calendar-specific, not generic date helpers.

## ANTI-PATTERNS
- Do not add child `AGENTS.md` files just because a package is large; add them only when local rules diverge materially.
- Do not trust generated or stale `llms.txt` content over source.
- Do not use `panic()` in ordinary library paths outside the `Must*` model.
- Do not ignore returned errors in non-Must code.
- Do not introduce circular dependencies between root packages.

## QUALITY CHECKS
```bash
make fmt
make lint
make test
make check
```

## NOTES
- Coverage badge in `README.md` currently shows overall coverage below the aspirational threshold; preserve existing test style instead of inventing new structure.
- `human/` is excluded from some coverage flows.
- Multi-platform behavior is real in this repo: build tags and cross-platform files are common enough to inspect before editing.
- **anyx 性能优化项目**（2026-05-09 进行中）：37 个函数逐个优化，每个函数独立 task。已完成的优化成果：
  * NewMap: 2.7x 性能提升
  * NewMapWithJson: 3-15% 性能提升（大数据场景）
  * NewMapWithYaml: 71% 性能提升（大数据场景）
  * NewMapWithAny: 22-31x 性能提升
  * Set: 当前实现已是最优（11 种方案对比）
  * Get: 2.86x 性能提升（并发场景）
  * Exists: 1.13x 性能提升
  * GetBool: 3-14% 性能提升
  * GetInt: 22-52% 性能提升
  详细进度见 `docs/reports/anyx-performance-optimization.md`

<!-- TRELLIS:START -->
# Trellis Instructions

These instructions are for AI assistants working in this project.

This project is managed by Trellis. The working knowledge you need lives under `.trellis/`:

- `.trellis/workflow.md` — development phases, when to create tasks, skill routing
- `.trellis/spec/` — package- and layer-scoped coding guidelines (read before writing code in a given layer)
- `.trellis/workspace/` — per-developer journals and session traces
- `.trellis/tasks/` — active and archived tasks (PRDs, research, jsonl context)

If a Trellis command is available on your platform (e.g. `/trellis:finish-work`, `/trellis:continue`), prefer it over manual steps. Not every platform exposes every command.

If you're using Codex or another agent-capable tool, additional project-scoped helpers may live in:
- `.agents/skills/` — reusable Trellis skills
- `.codex/agents/` — optional custom subagents

## Subagents

- ALWAYS wait for every spawned subagent to reach a terminal status before yielding, acting on partial results, or spawning followups.
  - On Codex, this means calling the `wait` tool with the subagent's thread id (requires `multi_agent_v2`). Do NOT infer completion from elapsed time.
  - On Claude Code / OpenCode, this means awaiting the Task/agent tool result before continuing.
- NEVER cancel or re-spawn a subagent that hasn't finished. If a subagent appears stuck, raise the wait timeout (Codex default 30s, max 1h) before judging it broken.
- Spawn subagents automatically when:
  - Parallelizable work (e.g., install + verify, npm test + typecheck, multiple tasks from plan)
  - Long-running or blocking tasks where a worker can run independently
  - Isolation for risky changes or checks

Managed by Trellis. Edits outside this block are preserved; edits inside may be overwritten by a future `trellis update`.

<!-- TRELLIS:END -->
