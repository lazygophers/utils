#!/bin/bash

# Pattern 函数性能基准测试脚本

cd /Users/luoxin/persons/go/lazygophers/utils/validator

echo "================================"
echo "Pattern 函数性能基准测试"
echo "================================"
echo ""

# 1. 测试当前实现的性能
echo "1. 当前实现 (Pattern 函数)"
go test -bench=BenchmarkPatternBaseline -benchmem -benchtime=1s 2>&1 | grep -E "Benchmark|ns/op|B/op|allocs"
echo ""

# 2. 测试不同场景
echo "2. 有效邮箱测试"
go test -bench="Baseline_Simple_Valid" -benchmem -benchtime=1s 2>&1 | grep -E "Benchmark|ns/op|B/op|allocs"
echo ""

echo "3. 无效邮箱测试"
go test -bench="Simple_Invalid" -benchmem -benchtime=1s 2>&1 | grep -E "Benchmark|ns/op|B/op|allocs"
echo ""

echo "4. 长字符串测试"
go test -bench="LongString" -benchmem -benchtime=1s 2>&1 | grep -E "Benchmark|ns/op|B/op|allocs"
echo ""

echo "5. 固定长度模式测试"
go test -bench="FixedLength" -benchmem -benchtime=1s 2>&1 | grep -E "Benchmark|ns/op|B/op|allocs"
echo ""

echo "6. 字面量模式测试"
go test -bench="Literal" -benchmem -benchtime=1s 2>&1 | grep -E "Benchmark|ns/op|B/op|allocs"
echo ""

echo "================================"
echo "基准测试完成"
echo "================================"
