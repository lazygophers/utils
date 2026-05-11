#!/bin/bash

# UUID 验证性能对比测试脚本

cd /Users/luoxin/persons/go/lazygophers/utils/validator

echo "=========================================="
echo "UUID 验证性能对比测试"
echo "=========================================="
echo ""

# 运行所有基准测试
go test -bench=BenchmarkValidateUUID -benchmem -benchtime=3s -run=^$ | tee uuid_comparison_results.txt

echo ""
echo "=========================================="
echo "测试完成！结果保存在 uuid_comparison_results.txt"
echo "=========================================="
echo ""

# 提取关键结果
echo "性能排名（有效 UUID）："
echo "----------------------------------------"
grep "Valid$" uuid_comparison_results.txt | sort -t' ' -k3 -n | head -10

echo ""
echo "性能排名（无效 UUID）："
echo "----------------------------------------"
grep "Invalid$" uuid_comparison_results.txt | sort -t' ' -k3 -n | head -10

echo ""
echo "内存分配对比（有效 UUID）："
echo "----------------------------------------"
grep "Valid$" uuid_comparison_results.txt | awk '{print $1, $4}' | sort -t' ' -k2 -n | head -10
