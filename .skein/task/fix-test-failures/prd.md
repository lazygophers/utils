# 修复全仓测试失败 — PRD (主入口)

## 目标
让 `go test ./...` 全绿（除 human 包）。覆盖率扫描暴露 3 个包系统性测试失败，根因各异。修后 anyx/singledo/xtime 三包测试全过，go.mod 残留单独提交。

## 边界
- 范围内：anyx 切片索引测试对齐实现、singledo Do error 缓存 bug、xtime 农历年算法 bug、go.mod log bump 单独提交
- 范围外：singledo DoCtx/panic 保护 API（已删测试，待后续单独 task 补）、新功能、性能优化
- 约束：遵循项目 fail-fast / 无 context / 无 defensive coding 约定（见 CLAUDE.md）；测试文件命名 `xxx_test.go`
- 已知：main 工作区已有预删改动（anyx YamlUnmarshalError 测试 + singledo 3 个 DoCtx 测试），属本 task 范围

## 验收标准
- [ ] `go test ./anyx/` PASS
- [ ] `go test ./singledo/` PASS
- [ ] `go test ./xtime/` PASS（农历 6 测试 + 全部）
- [ ] go.mod/go.sum log bump 单独 commit
- [ ] gofmt / go vet 各包 clean

## 索引
- 详细设计：[design.md](design.md)
- 调研收敛：[findings.md](findings.md)
- 任务/子任务/调度：task.json（`skein subtask list fix-test-failures`）
