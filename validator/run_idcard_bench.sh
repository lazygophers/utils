#!/bin/bash
# 在 validator 目录运行身份证基准测试

cd /Users/luoxin/persons/go/lazygophers/utils/validator

echo "运行身份证18位基准测试..."
echo "================================"

go test -bench=Benchmark_IDCard18 -benchmem -count=3 -run=^$ . 2>&1 | tee idcard18_bench_results.txt

echo ""
echo "基准测试完成！"
