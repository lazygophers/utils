#!/bin/bash

cd "$(dirname "$0")"

echo "=== Validator 性能基准测试 ==="
echo ""

echo "运行原始实现基准测试..."
go test -bench=BenchmarkOriginal -benchmem -benchtime=1s > original_bench.txt 2>&1

echo "运行优化实现基准测试..."
go test -bench=BenchmarkOptimized -benchmem -benchtime=1s > optimized_bench.txt 2>&1

echo ""
echo "=== 原始实现结果 ==="
cat original_bench.txt | grep -E "(Benchmark|goos|goarch|cpu)"

echo ""
echo "=== 优化实现结果 ==="
cat optimized_bench.txt | grep -E "(Benchmark|goos|goarch|cpu)"

echo ""
echo "=== Tag 解析对比 ==="
echo "原始实现:"
go test -bench=BenchmarkOriginalParseTag -benchmem -benchtime=500ms 2>&1 | grep -E "(Benchmark|ns/op)"
echo ""
echo "优化实现:"
go test -bench=BenchmarkFastParseTag -benchmem -benchtime=500ms 2>&1 | grep -E "(Benchmark|ns/op)"

echo ""
echo "=== 类型缓存性能 ==="
go test -bench=BenchmarkTypeCacheHit -benchmem -benchtime=500ms 2>&1 | grep -E "(Benchmark|ns/op)"

rm -f original_bench.txt optimized_bench.txt
