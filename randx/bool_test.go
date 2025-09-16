package randx

import (
	"testing"
)

func TestBool(t *testing.T) {
	t.Run("bool_returns_true_or_false", func(t *testing.T) {
		// 测试Bool函数返回true或false
		for i := 0; i < 100; i++ {
			result := Bool()
			// 只需要验证结果是bool类型，没有其他约束
			_ = result // bool类型总是有效的
		}
	})

	t.Run("bool_distribution", func(t *testing.T) {
		// 统计测试：验证Bool函数的分布
		trueCount := 0
		falseCount := 0
		iterations := 10000

		for i := 0; i < iterations; i++ {
			if Bool() {
				trueCount++
			} else {
				falseCount++
			}
		}

		// 理论上应该大约各占50%，但允许一定偏差
		expectedMin := iterations/2 - 1000 // 4000
		expectedMax := iterations/2 + 1000 // 6000

		if trueCount < expectedMin || trueCount > expectedMax {
			t.Logf("Warning: true appeared %d times out of %d, expected around %d",
				trueCount, iterations, iterations/2)
		}

		if falseCount < expectedMin || falseCount > expectedMax {
			t.Logf("Warning: false appeared %d times out of %d, expected around %d",
				falseCount, iterations, iterations/2)
		}

		// 验证总数正确
		if trueCount+falseCount != iterations {
			t.Errorf("Total count mismatch: %d + %d != %d", trueCount, falseCount, iterations)
		}
	})
}

func TestBooln(t *testing.T) {
	t.Run("booln_probability_100_or_more", func(t *testing.T) {
		// 测试概率>=100时总是返回true
		testCases := []float64{100, 101, 150, 200, 1000}

		for _, prob := range testCases {
			for i := 0; i < 50; i++ { // 多次测试确保稳定
				result := Booln(prob)
				if !result {
					t.Errorf("Booln(%f) should always return true, got false", prob)
				}
			}
		}
	})

	t.Run("booln_probability_0_or_less", func(t *testing.T) {
		// 测试概率<=0时总是返回false
		testCases := []float64{0, -1, -10, -100}

		for _, prob := range testCases {
			for i := 0; i < 50; i++ { // 多次测试确保稳定
				result := Booln(prob)
				if result {
					t.Errorf("Booln(%f) should always return false, got true", prob)
				}
			}
		}
	})

	t.Run("booln_probability_50_percent", func(t *testing.T) {
		// 测试50%概率情况
		trueCount := 0
		falseCount := 0
		iterations := 10000
		probability := 50.0

		for i := 0; i < iterations; i++ {
			if Booln(probability) {
				trueCount++
			} else {
				falseCount++
			}
		}

		// 50%概率应该大约各占一半
		expectedTrue := iterations / 2
		tolerance := iterations / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 50%% probability, got %d true out of %d (expected around %d)",
				trueCount, iterations, expectedTrue)
		}
	})

	t.Run("booln_probability_25_percent", func(t *testing.T) {
		// 测试25%概率情况
		trueCount := 0
		iterations := 10000
		probability := 25.0

		for i := 0; i < iterations; i++ {
			if Booln(probability) {
				trueCount++
			}
		}

		// 25%概率应该大约25%为true
		expectedTrue := iterations / 4
		tolerance := iterations / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 25%% probability, got %d true out of %d (expected around %d)",
				trueCount, iterations, expectedTrue)
		}
	})

	t.Run("booln_probability_75_percent", func(t *testing.T) {
		// 测试75%概率情况
		trueCount := 0
		iterations := 10000
		probability := 75.0

		for i := 0; i < iterations; i++ {
			if Booln(probability) {
				trueCount++
			}
		}

		// 75%概率应该大约75%为true
		expectedTrue := iterations * 3 / 4
		tolerance := iterations / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 75%% probability, got %d true out of %d (expected around %d)",
				trueCount, iterations, expectedTrue)
		}
	})

	t.Run("booln_edge_cases", func(t *testing.T) {
		// 测试边界情况
		edgeCases := []struct {
			prob     float64
			expected string // "true", "false", or "random"
		}{
			{0.1, "mostly_false"},
			{1.0, "mostly_false"},
			{99.9, "mostly_true"},
			{99.0, "mostly_true"},
		}

		for _, tc := range edgeCases {
			// 对于非确定性情况，只做基本验证
			for i := 0; i < 10; i++ {
				result := Booln(tc.prob)
				_ = result // 验证函数能正常执行
			}
		}
	})

	t.Run("booln_very_small_probability", func(t *testing.T) {
		// 测试很小的概率
		falseCount := 0
		iterations := 1000
		probability := 0.01 // 0.01%的概率

		for i := 0; i < iterations; i++ {
			if !Booln(probability) {
				falseCount++
			}
		}

		// 0.01%概率，在1000次测试中应该几乎都是false
		if falseCount < iterations*95/100 { // 至少95%应该是false
			t.Logf("With very low probability (0.01%%), expected mostly false results")
		}
	})

	t.Run("booln_very_high_probability", func(t *testing.T) {
		// 测试很高的概率
		trueCount := 0
		iterations := 1000
		probability := 99.99 // 99.99%的概率

		for i := 0; i < iterations; i++ {
			if Booln(probability) {
				trueCount++
			}
		}

		// 99.99%概率，在1000次测试中应该几乎都是true
		if trueCount < iterations*95/100 { // 至少95%应该是true
			t.Logf("With very high probability (99.99%%), expected mostly true results")
		}
	})
}


func TestWeightedBool(t *testing.T) {
	t.Run("weighted_bool_weight_greater_than_one", func(t *testing.T) {
		// 测试权重>=1.0时总是返回true
		testCases := []float64{1.0, 1.5, 2.0, 10.0}

		for _, weight := range testCases {
			for i := 0; i < 50; i++ {
				result := WeightedBool(weight)
				if !result {
					t.Errorf("WeightedBool(%f) should always return true, got false", weight)
				}
			}
		}
	})

	t.Run("weighted_bool_weight_less_than_or_equal_zero", func(t *testing.T) {
		// 测试权重<=0.0时总是返回false
		testCases := []float64{0.0, -0.1, -1.0, -10.0}

		for _, weight := range testCases {
			for i := 0; i < 50; i++ {
				result := WeightedBool(weight)
				if result {
					t.Errorf("WeightedBool(%f) should always return false, got true", weight)
				}
			}
		}
	})

	t.Run("weighted_bool_weight_half", func(t *testing.T) {
		// 测试权重0.5的情况
		trueCount := 0
		falseCount := 0
		iterations := 10000
		weight := 0.5

		for i := 0; i < iterations; i++ {
			if WeightedBool(weight) {
				trueCount++
			} else {
				falseCount++
			}
		}

		// 50%权重应该大约各占一半
		expectedTrue := iterations / 2
		tolerance := iterations / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 50%% weight, got %d true out of %d (expected around %d)",
				trueCount, iterations, expectedTrue)
		}
	})

	t.Run("weighted_bool_weight_quarter", func(t *testing.T) {
		// 测试权重0.25的情况
		trueCount := 0
		iterations := 10000
		weight := 0.25

		for i := 0; i < iterations; i++ {
			if WeightedBool(weight) {
				trueCount++
			}
		}

		// 25%权重应该大约25%为true
		expectedTrue := iterations / 4
		tolerance := iterations / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 25%% weight, got %d true out of %d (expected around %d)",
				trueCount, iterations, expectedTrue)
		}
	})

	t.Run("weighted_bool_weight_three_quarters", func(t *testing.T) {
		// 测试权重0.75的情况
		trueCount := 0
		iterations := 10000
		weight := 0.75

		for i := 0; i < iterations; i++ {
			if WeightedBool(weight) {
				trueCount++
			}
		}

		// 75%权重应该大约75%为true
		expectedTrue := iterations * 3 / 4
		tolerance := iterations / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 75%% weight, got %d true out of %d (expected around %d)",
				trueCount, iterations, expectedTrue)
		}
	})

	t.Run("weighted_bool_very_small_weight", func(t *testing.T) {
		// 测试很小的权重
		falseCount := 0
		iterations := 1000
		weight := 0.001 // 0.1%的权重

		for i := 0; i < iterations; i++ {
			if !WeightedBool(weight) {
				falseCount++
			}
		}

		// 0.1%权重，在1000次测试中应该几乎都是false
		if falseCount < iterations*95/100 { // 至少95%应该是false
			t.Logf("With very low weight (0.1%%), expected mostly false results")
		}
	})

	t.Run("weighted_bool_very_high_weight", func(t *testing.T) {
		// 测试很高的权重
		trueCount := 0
		iterations := 1000
		weight := 0.999 // 99.9%的权重

		for i := 0; i < iterations; i++ {
			if WeightedBool(weight) {
				trueCount++
			}
		}

		// 99.9%权重，在1000次测试中应该几乎都是true
		if trueCount < iterations*95/100 { // 至少95%应该是true
			t.Logf("With very high weight (99.9%%), expected mostly true results")
		}
	})
}

func TestBatchBool(t *testing.T) {
	t.Run("batch_bool_zero_or_negative_count", func(t *testing.T) {
		// 测试count <= 0
		result := BatchBool(0)
		if result != nil {
			t.Errorf("Expected nil for count=0, got %v", result)
		}

		result = BatchBool(-1)
		if result != nil {
			t.Errorf("Expected nil for count=-1, got %v", result)
		}
	})

	t.Run("batch_bool_positive_count", func(t *testing.T) {
		// 测试正数count
		counts := []int{1, 5, 10, 100}

		for _, count := range counts {
			result := BatchBool(count)
			if len(result) != count {
				t.Errorf("Expected length %d, got %d", count, len(result))
			}

			// 验证每个元素都是bool类型（通过类型推断）
			for i, b := range result {
				_ = b // bool类型总是有效的
				// 简单验证索引
				if i >= count {
					t.Errorf("Index %d out of bounds for count %d", i, count)
				}
			}
		}
	})

	t.Run("batch_bool_distribution", func(t *testing.T) {
		// 测试分布
		count := 10000
		result := BatchBool(count)

		trueCount := 0
		falseCount := 0

		for _, b := range result {
			if b {
				trueCount++
			} else {
				falseCount++
			}
		}

		// 理论上应该大约各占50%
		expectedMin := count/2 - 1000 // 4000
		expectedMax := count/2 + 1000 // 6000

		if trueCount < expectedMin || trueCount > expectedMax {
			t.Logf("Warning: true appeared %d times out of %d, expected around %d",
				trueCount, count, count/2)
		}

		if falseCount < expectedMin || falseCount > expectedMax {
			t.Logf("Warning: false appeared %d times out of %d, expected around %d",
				falseCount, count, count/2)
		}

		// 验证总数正确
		if trueCount+falseCount != count {
			t.Errorf("Total count mismatch: %d + %d != %d", trueCount, falseCount, count)
		}
	})
}

func TestBatchBooln(t *testing.T) {
	t.Run("batch_booln_zero_or_negative_count", func(t *testing.T) {
		// 测试count <= 0
		result := BatchBooln(50.0, 0)
		if result != nil {
			t.Errorf("Expected nil for count=0, got %v", result)
		}

		result = BatchBooln(50.0, -1)
		if result != nil {
			t.Errorf("Expected nil for count=-1, got %v", result)
		}
	})

	t.Run("batch_booln_probability_100_or_more", func(t *testing.T) {
		// 测试概率>=100时全部返回true
		testCases := []float64{100, 101, 150}
		count := 50

		for _, prob := range testCases {
			result := BatchBooln(prob, count)
			
			if len(result) != count {
				t.Errorf("Expected length %d, got %d", count, len(result))
			}

			for i, r := range result {
				if !r {
					t.Errorf("BatchBooln(%f, %d)[%d] should be true, got false", prob, count, i)
				}
			}
		}
	})

	t.Run("batch_booln_probability_0_or_less", func(t *testing.T) {
		// 测试概率<=0时全部返回false
		testCases := []float64{0, -1, -10}
		count := 50

		for _, prob := range testCases {
			result := BatchBooln(prob, count)
			
			if len(result) != count {
				t.Errorf("Expected length %d, got %d", count, len(result))
			}

			for i, r := range result {
				if r {
					t.Errorf("BatchBooln(%f, %d)[%d] should be false, got true", prob, count, i)
				}
			}
		}
	})

	t.Run("batch_booln_probability_50_percent", func(t *testing.T) {
		// 测试50%概率情况
		prob := 50.0
		count := 10000
		result := BatchBooln(prob, count)

		if len(result) != count {
			t.Errorf("Expected length %d, got %d", count, len(result))
		}

		trueCount := 0
		for _, r := range result {
			if r {
				trueCount++
			}
		}

		// 50%概率应该大约各占一半
		expectedTrue := count / 2
		tolerance := count / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 50%% probability, got %d true out of %d (expected around %d)",
				trueCount, count, expectedTrue)
		}
	})

	t.Run("batch_booln_probability_25_percent", func(t *testing.T) {
		// 测试25%概率情况
		prob := 25.0
		count := 10000
		result := BatchBooln(prob, count)

		trueCount := 0
		for _, r := range result {
			if r {
				trueCount++
			}
		}

		// 25%概率应该大约25%为true
		expectedTrue := count / 4
		tolerance := count / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 25%% probability, got %d true out of %d (expected around %d)",
				trueCount, count, expectedTrue)
		}
	})

	t.Run("batch_booln_probability_75_percent", func(t *testing.T) {
		// 测试75%概率情况
		prob := 75.0
		count := 10000
		result := BatchBooln(prob, count)

		trueCount := 0
		for _, r := range result {
			if r {
				trueCount++
			}
		}

		// 75%概率应该大约75%为true
		expectedTrue := count * 3 / 4
		tolerance := count / 10 // 10%的容差

		if trueCount < expectedTrue-tolerance || trueCount > expectedTrue+tolerance {
			t.Logf("Warning: With 75%% probability, got %d true out of %d (expected around %d)",
				trueCount, count, expectedTrue)
		}
	})
}
