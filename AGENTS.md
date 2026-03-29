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

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **utils** (7430 symbols, 17178 relationships, 275 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## When Debugging

1. `gitnexus_query({query: "<error or symptom>"})` — find execution flows related to the issue
2. `gitnexus_context({name: "<suspect function>"})` — see all callers, callees, and process participation
3. `READ gitnexus://repo/utils/process/{processName}` — trace the full execution flow step by step
4. For regressions: `gitnexus_detect_changes({scope: "compare", base_ref: "main"})` — see what your branch changed

## When Refactoring

- **Renaming**: MUST use `gitnexus_rename({symbol_name: "old", new_name: "new", dry_run: true})` first. Review the preview — graph edits are safe, text_search edits need manual review. Then run with `dry_run: false`.
- **Extracting/Splitting**: MUST run `gitnexus_context({name: "target"})` to see all incoming/outgoing refs, then `gitnexus_impact({target: "target", direction: "upstream"})` to find all external callers before moving code.
- After any refactor: run `gitnexus_detect_changes({scope: "all"})` to verify only expected files changed.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Tools Quick Reference

| Tool | When to use | Command |
|------|-------------|---------|
| `query` | Find code by concept | `gitnexus_query({query: "auth validation"})` |
| `context` | 360-degree view of one symbol | `gitnexus_context({name: "validateUser"})` |
| `impact` | Blast radius before editing | `gitnexus_impact({target: "X", direction: "upstream"})` |
| `detect_changes` | Pre-commit scope check | `gitnexus_detect_changes({scope: "staged"})` |
| `rename` | Safe multi-file rename | `gitnexus_rename({symbol_name: "old", new_name: "new", dry_run: true})` |
| `cypher` | Custom graph queries | `gitnexus_cypher({query: "MATCH ..."})` |

## Impact Risk Levels

| Depth | Meaning | Action |
|-------|---------|--------|
| d=1 | WILL BREAK — direct callers/importers | MUST update these |
| d=2 | LIKELY AFFECTED — indirect deps | Should test |
| d=3 | MAY NEED TESTING — transitive | Test if critical path |

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/utils/context` | Codebase overview, check index freshness |
| `gitnexus://repo/utils/clusters` | All functional areas |
| `gitnexus://repo/utils/processes` | All execution flows |
| `gitnexus://repo/utils/process/{name}` | Step-by-step execution trace |

## Self-Check Before Finishing

Before completing any code modification task, verify:
1. `gitnexus_impact` was run for all modified symbols
2. No HIGH/CRITICAL risk warnings were ignored
3. `gitnexus_detect_changes()` confirms changes match expected scope
4. All d=1 (WILL BREAK) dependents were updated

## Keeping the Index Fresh

After committing code changes, the GitNexus index becomes stale. Re-run analyze to update it:

```bash
npx gitnexus analyze
```

If the index previously included embeddings, preserve them by adding `--embeddings`:

```bash
npx gitnexus analyze --embeddings
```

To check whether embeddings exist, inspect `.gitnexus/meta.json` — the `stats.embeddings` field shows the count (0 means no embeddings). **Running analyze without `--embeddings` will delete any previously generated embeddings.**

> Claude Code users: A PostToolUse hook handles this automatically after `git commit` and `git merge`.

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
