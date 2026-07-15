# 修复全仓测试失败 — 调研收敛

## anyx（7 测试函数 ~25 子用例）— 测试断言落后于实现扩展

### 根因
主路径 `navigateToValue`（map_any.go:1064-1144）**已全面支持所有数值切片类型**（[]int/[]int8..64/[]uint../[]float32/64/[]string/[]bool/[]byte/[]interface{}）。
fallback 函数 `accessGenericSlice`（map_any.go:2503）也 fast-path 支持 []uint/[]float32/[]int32 + reflect 兜底全切片。

但失败测试（`TestAccessGenericSlice_*` / `TestNavigateToValue_*` / `TestMapGetWithSeparator_Coverage`）仍断言这些类型应返回 "invalid slice type" 错误——**旧契约**，实现已扩展为 feature。

### 修法
更新测试断言对齐实现：原 "应报错" 的用例改为 "应成功返回值"。失败用例类型：[]uint/[]float32/[]int32/[]int16/[]uint8/16/64/[]interface{}显式/自定义切片/空切片/单元素/负索引/超大索引/特殊字符键/空索引/负数索引。
注意区分：真正的越界（[]uint 超大索引）仍应报 ErrOutOfRange；仅"类型不支持"断言需翻转。

### 引用
- map_any.go:1064-1144（主路径全切片支持）
- map_any.go:2503-2531（accessGenericSlice fast+reflect）
- map_any_test.go:7460-7464（失败断言示例：期望 error 实得 nil）

## singledo（error_not_cached）— Do error 分支未清 call 缓存

### 根因
`Single.Do`（singledo.go:44-50）error 分支：
```go
if callM.err == nil {  // error 时整块跳过
    s.last = now
    s.result = callM.val
    s.call = nil         // ← error 时不清 s.call
}
```
error 后 `s.call` 保持非 nil。第 2 次 Do 走 L30 `callM := s.call`（非 nil）→ L32-33 返回缓存的旧 callM.err。测试 `error_not_cached` 期望重试，实际拿到缓存 error（calls=1, v=""）。

### 修法
error 分支也清 `s.call = nil`（不更新 last/result，保持 error 不缓存语义）：
```go
s.mux.Lock()
s.call = nil  // 无条件清，error 也不缓存 call
if callM.err == nil {
    s.last = now
    s.result = callM.val
}
s.mux.Unlock()
```

### 引用
- singledo.go:44-50（error 分支）
- group_test.go:75-97（TestSingle_Do/error_not_cached）

## xtime（6 农历测试）— FromSolarTimestamp 农历年 off-by-one

### 根因（初步定位，需数值调试确认）
`FromSolarTimestamp`（lunar.go:260-328）农历年循环：
```go
for i = 1900; i < 2101 && offset > 0; i++ {
    daysOfYear = daysOfLunarYear(i)
    offset -= daysOfYear
}
if offset < 0 { offset += daysOfYear; i-- }
lunarYear = i
```
系统性偏差（2023-01-22 正月初一返回 2022 而非 2023），疑似 offset 边界（offset==0 时 i 已自增）或 lunars 数据表基准偏移。6 测试全挂 → 系统性非个别边界。

### 修法
exec 阶段数值调试 FromSolarTimestamp：打印 2023-01-22 / 2024-02-10 / 1900-01-31 等已知日期的 offset 轨迹，定位是循环边界还是数据表问题。修正后 6 测试应全过。风险：农历算法核心，改动需谨慎验证多组已知日期。

### 引用
- lunar.go:260-328（FromSolarTimestamp）
- lunar.go:346-376（daysOfLunarYear / leapMonth / leapDays，数据表 lunars）
- lunar_test.go:865-880（TestLunarYearAliasFormat）/ 920-936（TestLunarDateFormat）

## go.mod — lazygophers/log 版本 bump 残留

非本 task 引入（go mod tidy 副作用）。单独 commit。
