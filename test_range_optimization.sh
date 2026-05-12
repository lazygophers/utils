#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils/validator

echo "========================================="
echo "Range 函数优化验证"
echo "========================================="
echo ""

echo "1. 运行 Range 正确性测试..."
go test -run=TestRange -v
echo ""

echo "2. 运行优化后的 Range 性能基准测试..."
go test -bench=BenchmarkRange -run=^$ -benchmem -count=3 2>&1 | grep -E "^(BenchmarkRange|goos|goarch|pkg)" | tee /Users/luoxin/persons/go/lazygophers/utils/range_optimization_final_results.txt
echo ""

echo "3. 对比优化前后性能..."
echo "原始 Range Float64: ~138.8 ns/op"
echo "优化后 Range Float64: 查看上面的 BenchmarkRange_BranchPrediction_Float64 结果"
echo "预期提升: 49.5% (70.1 ns/op)"
echo ""

echo "验证完成！"
