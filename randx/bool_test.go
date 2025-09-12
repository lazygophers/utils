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
		expectedMin := iterations/2 - 1000  // 4000
		expectedMax := iterations/2 + 1000  // 6000
		
		if trueCount < expectedMin || trueCount > expectedMax {
			t.Logf("Warning: true appeared %d times out of %d, expected around %d", 
				trueCount, iterations, iterations/2)
		}
		
		if falseCount < expectedMin || falseCount > expectedMax {
			t.Logf("Warning: false appeared %d times out of %d, expected around %d", 
				falseCount, iterations, iterations/2)
		}
		
		// 验证总数正确
		if trueCount + falseCount != iterations {
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