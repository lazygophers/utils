#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils

echo "开始运行 Struct 基准测试..."
echo "=========================================="

go test -bench=Benchmark_Struct_Original -benchmem -run=^$ -benchtime=2s ./validator > validator/STRUCT_BENCH_RESULTS.txt 2>&1 &
PID=$!

echo "测试进程 PID: $PID"
echo "等待测试完成..."

wait $PID

echo "=========================================="
echo "基准测试完成！"

# 显示结果
echo ""
echo "========== 原始实现结果 =========="
cat validator/STRUCT_BENCH_RESULTS.txt | grep "Benchmark_Struct_Original"

echo ""
echo "========== 对象池优化结果 =========="
cat validator/STRUCT_BENCH_RESULTS.txt | grep "Benchmark_Struct_Pool"

echo ""
echo "========== 减少反射优化结果 =========="
cat validator/STRUCT_BENCH_RESULTS.txt | grep "Benchmark_Struct_LessReflect"

echo ""
echo "========== 组合优化结果 =========="
cat validator/STRUCT_BENCH_RESULTS.txt | grep "Benchmark_Struct_Combined"

echo ""
echo "========== 完整结果保存在: validator/STRUCT_BENCH_RESULTS.txt =========="
