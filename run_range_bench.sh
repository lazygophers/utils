#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils/validator

echo "运行 Range 基准测试..."
echo "======================="

# 只运行基准测试，不运行普通测试
go test -bench=BenchmarkRange -run=^$ -benchmem -count=5 2>&1 | tee /Users/luoxin/persons/go/lazygophers/utils/range_benchmark_final_results.txt

echo ""
echo "基准测试完成！"
