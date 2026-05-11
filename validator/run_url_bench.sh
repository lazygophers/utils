#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils/validator

echo "运行 URL 验证基准测试..."
echo "=========================================="

# 运行基准测试
go test -bench="BenchmarkURL" -benchmem -benchtime=500ms -run="^$" 2>&1 | tee url_benchmark_raw.txt

# 提取结果
echo ""
echo "=========================================="
echo "基准测试结果摘要："
echo "=========================================="

grep -E "BenchmarkURL.*ns/op" url_benchmark_raw.txt | sort -t: -k2 -n

echo ""
echo "完整的基准测试结果已保存到: url_benchmark_raw.txt"
