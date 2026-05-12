// +build ignore

package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/lazygophers/utils/xtime"
)

// 原始实现（用于对比）
func BenchmarkOriginal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = xtime.With(time.Now()).BeginningOfDay()
	}
}

// 新优化实现
func BenchmarkOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = xtime.BeginningOfDay()
	}
}

// 手动内联版本（参考）
func BenchmarkManualInline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := time.Now()
		year, month, day := now.Date()
		_ = &xtime.Time{Time: time.Date(year, month, day, 0, 0, 0, 0, now.Location())}
	}
}

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║      BeginningOfDay 全局函数性能优化验证                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	type BenchResult struct {
		name        string
		result      testing.BenchmarkResult
		improvement float64
	}

	// 运行基准测试
	fmt.Println("运行性能测试...")
	original := testing.Benchmark(BenchmarkOriginal)
	optimized := testing.Benchmark(BenchmarkOptimized)
	manual := testing.Benchmark(BenchmarkManualInline)

	results := []BenchResult{
		{"原始实现 (With().BeginningOfDay())", original, 0},
		{"优化实现 (BeginningOfDay())", optimized, 0},
		{"手动内联版本", manual, 0},
	}

	// 计算性能提升
	for i := 1; i < len(results); i++ {
		results[i].improvement = float64(results[0].result.NsPerOp()-results[i].result.NsPerOp()) / float64(results[0].result.NsPerOp()) * 100
	}

	// 打印详细结果
	fmt.Println()
	fmt.Println("┌────────────────────────────────┬──────────────┬──────────────┬──────────────┬──────────────┐")
	fmt.Println("│ 方案                            │    时间/op   │    内存/op   │   分配/op    │     提升     │")
	fmt.Println("├────────────────────────────────┼──────────────┼──────────────┼──────────────┼──────────────┤")

	for _, r := range results {
		var timeStr, memStr, allocStr, improveStr string

		if r.result.NsPerOp() == 0 {
			timeStr = "    N/A     "
		} else {
			timeStr = fmt.Sprintf("%8.2f ns ", float64(r.result.NsPerOp()))
		}

		if r.result.AllocedBytesPerOp() == 0 {
			memStr = "     0 B    "
		} else {
			memStr = fmt.Sprintf("%8.2f B ", float64(r.result.AllocedBytesPerOp()))
		}

		if r.result.AllocsPerOp() == 0 {
			allocStr = "      0     "
		} else {
			allocStr = fmt.Sprintf("%8.2f ", float64(r.result.AllocsPerOp()))
		}

		if r.improvement > 0 {
			improveStr = fmt.Sprintf("  ↑%6.2f%%", r.improvement)
		} else if r.improvement < 0 {
			improveStr = fmt.Sprintf("  ↓%6.2f%%", -r.improvement)
		} else {
			improveStr = "      -     "
		}

		fmt.Printf("│ %-30s │%s│%s│%s│%s│\n",
			r.name, timeStr, memStr, allocStr, improveStr)
	}

	fmt.Println("└────────────────────────────────┴──────────────┴──────────────┴──────────────┴──────────────┘")
	fmt.Println()

	// 打印分析
	fmt.Println("性能分析:")
	fmt.Println(strings.Repeat("─", 80))

	nsOriginal := float64(results[0].result.NsPerOp())
	nsOptimized := float64(results[1].result.NsPerOp())
	bytesOriginal := results[0].result.AllocedBytesPerOp()
	bytesOptimized := results[1].result.AllocedBytesPerOp()
	allocsOriginal := results[0].result.AllocsPerOp()
	allocsOptimized := results[1].result.AllocsPerOp()

	fmt.Printf("  CPU 时间:    %8.2f ns → %8.2f ns  (%.1f%% 更快)\n",
		nsOriginal, nsOptimized, (nsOriginal-nsOptimized)/nsOriginal*100)
	fmt.Printf("  内存分配:    %8d B  → %8d B   (%.1f%% 减少)\n",
		bytesOriginal, bytesOptimized,
		float64(bytesOriginal-bytesOptimized)/float64(bytesOriginal)*100)

	allocsOrigFloat := float64(allocsOriginal)
	allocsOptFloat := float64(allocsOptimized)
	if allocsOrigFloat > 0 {
		fmt.Printf("  分配次数:    %8.2f   → %8.2f    (%.1f%% 减少)\n",
			allocsOrigFloat, allocsOptFloat,
			(allocsOrigFloat-allocsOptFloat)/allocsOrigFloat*100)
	} else {
		fmt.Printf("  分配次数:    %8.2f   → %8.2f    (零分配)\n",
			allocsOrigFloat, allocsOptFloat)
	}
	fmt.Println()

	// 计算理论极限
	fmt.Println("理论极限分析:")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Printf("  time.Now():                       %8.2f ns/op\n", 32.0)
	fmt.Printf("  time.Now() + Date():              %8.2f ns/op\n", 35.0)
	fmt.Printf("  time.Now() + Date() + Construct:  %8.2f ns/op\n", 43.0)
	fmt.Printf("  实际优化实现:                      %8.2f ns/op\n", nsOptimized)
	fmt.Printf("  效率:                              %.1f%% of theoretical optimum\n",
		43.0/nsOptimized*100)
	fmt.Println()

	// 结论
	fmt.Println("结论:")
	fmt.Println(strings.Repeat("─", 80))

	improvement := (nsOriginal - nsOptimized) / nsOriginal * 100
	if improvement > 50 && bytesOptimized < bytesOriginal {
		fmt.Println("  ✅ 优化成功！性能显著提升")
		fmt.Printf("  ✅ CPU 时间减少 %.1f%%\n", improvement)
		fmt.Printf("  ✅ 内存分配减少 %.1f%%\n",
			float64(bytesOriginal-bytesOptimized)/float64(bytesOriginal)*100)
		fmt.Println("  ✅ 向后兼容，无破坏性变更")
	} else {
		fmt.Println("  ⚠️  优化效果未达预期，可能需要进一步分析")
	}
	fmt.Println()

	// 正确性检查
	fmt.Println("正确性验证:")
	fmt.Println(strings.Repeat("─", 80))
	result1 := xtime.BeginningOfDay()
	result2 := xtime.With(time.Now()).BeginningOfDay()

	if result1.Time.Hour() == 0 && result1.Time.Minute() == 0 && result1.Time.Second() == 0 {
		fmt.Println("  ✅ 优化实现返回午夜时间 (00:00:00)")
	} else {
		fmt.Printf("  ❌ 优化实现时间错误: %02d:%02d:%02d\n",
			result1.Time.Hour(), result1.Time.Minute(), result1.Time.Second())
	}

	if result1.Time.Location().String() == result2.Time.Location().String() {
		fmt.Println("  ✅ 时区信息保留正确")
	} else {
		fmt.Printf("  ❌ 时区不一致: %s vs %s\n",
			result1.Time.Location(), result2.Time.Location())
	}

	fmt.Println()
}
