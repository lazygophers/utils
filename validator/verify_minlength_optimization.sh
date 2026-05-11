#!/bin/bash

# MinLength 优化验证脚本

echo "=========================================="
echo "MinLength 性能优化验证"
echo "=========================================="

cd validator

# 运行正确性测试
echo "1. 运行正确性测试..."
go test -run=TestMinLength -v . 2>&1 | grep -E "(PASS|FAIL|TestMinLength)" | head -20

echo ""
echo "2. 编译性能测试..."
go test -c -o /tmp/test_minlength_verify . 2>&1 | head -10

if [ -f /tmp/test_minlength_verify ]; then
    echo "3. 运行基准测试..."
    /tmp/test_minlength_verify -test.bench=BenchmarkMinLength -test.benchmem -test.benchtime=2s 2>&1 | grep "BenchmarkMinLength"

    echo ""
    echo "4. 性能对比分析..."
    echo "原始版本性能（测试中的原始实现）vs 当前优化版本"
else
    echo "编译失败，无法运行性能测试"
fi

echo ""
echo "=========================================="
echo "验证完成"
echo "=========================================="
