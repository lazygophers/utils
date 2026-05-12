#!/bin/bash

# EndOfQuarter 全局函数性能测试
# 手动基准测试脚本

set -e

echo "=== EndOfQuarter Global Function Performance Test ==="
echo ""

# 测试函数列表
ITERATIONS=1000000

echo "测试次数: $ITERATIONS"
echo ""

# 原始实现
echo "Testing Original implementation..."
START=$(date +%s%N)
for i in $(seq 1 $ITERATIONS); do
    result=$(go run -c 'package main; import ("fmt"; "time"; "github.com/lazygophers/utils/xtime"); func main() { _ = xtime.With(time.Now()).EndOfQuarter() }')
done
END=$(date +%s%N)
ORIGINAL_NS=$(( (END - START) / ITERATIONS ))
echo "Original: $ORIGINAL_NS ns/op"
echo ""

# 变体1: 内联逻辑
echo "Testing Variant1 (Inlined)..."
START=$(date +%s%N)
for i in $(seq 1 $ITERATIONS); do
    result=$(go run -c 'package main; import ("fmt"; "time"; "github.com/lazygophers/utils/xtime"); func main() { t := time.Now(); year, month, _ := t.Date(); quarter := (month-1)/3 + 1; endQuarterMonth := quarter * 3; _ = &xtime.Time{Time: time.Date(year, time.Month(endQuarterMonth+1), 0, 23, 59, 59, 999999999, t.Location()), Config: &xtime.Config{WeekStartDay: time.Monday, TimeLocation: time.Local, TimeFormats: []string{}, Monotonic: time.Now()}} }')
done
END=$(date +%s%N)
VARIANT1_NS=$(( (END - START) / ITERATIONS ))
echo "Variant1: $VARIANT1_NS ns/op"
echo ""

# 变体7: 最优化版本
echo "Testing Variant7 (Optimized)..."
START=$(date +%s%N)
for i in $(seq 1 $ITERATIONS); do
    result=$(go run -c 'package main; import ("fmt"; "time"; "github.com/lazygophers/utils/xtime"); func main() { now := time.Now(); year := now.Year(); month := now.Month(); quarter := (month-1)/3 + 1; endQuarterMonth := quarter * 3; _ = &xtime.Time{Time: time.Date(year, endQuarterMonth+1, 0, 23, 59, 59, 999999999, now.Location()), Config: &xtime.Config{WeekStartDay: time.Monday, TimeLocation: now.Location()}} }')
done
END=$(date +%s%N)
VARIANT7_NS=$(( (END - START) / ITERATIONS ))
echo "Variant7: $VARIANT7_NS ns/op"
echo ""

# 变体11: 完全内联
echo "Testing Variant11 (Fully Inlined)..."
START=$(date +%s%N)
for i in $(seq 1 $ITERATIONS); do
    result=$(go run -c 'package main; import ("fmt"; "time"; "github.com/lazygophers/utils/xtime"); func main() { t := time.Now(); _ = &xtime.Time{Time: time.Date(t.Year(), ((t.Month()-1)/3+1)*3+1, 0, 23, 59, 59, 999999999, t.Location()), Config: &xtime.Config{WeekStartDay: time.Monday, TimeLocation: t.Location()}} }')
done
END=$(date +%s%N)
VARIANT11_NS=$(( (END - START) / ITERATIONS ))
echo "Variant11: $VARIANT11_NS ns/op"
echo ""

# 计算加速比
echo "=== Results ==="
echo "Original: $ORIGINAL_NS ns/op"
echo "Variant1: $VARIANT1_NS ns/op (speedup: $(echo "scale=2; $ORIGINAL_NS / $VARIANT1_NS" | bc)x)"
echo "Variant7: $VARIANT7_NS ns/op (speedup: $(echo "scale=2; $ORIGINAL_NS / $VARIANT7_NS" | bc)x)"
echo "Variant11: $VARIANT11_NS ns/op (speedup: $(echo "scale=2; $ORIGINAL_NS / $VARIANT11_NS" | bc)x)"
