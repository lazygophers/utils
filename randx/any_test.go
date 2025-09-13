package randx

import (
	"testing"
)

func TestChoose(t *testing.T) {
	t.Run("choose_from_empty_slice", func(t *testing.T) {
		// 测试空切片的情况
		var empty []int
		result := Choose(empty)
		if result != 0 { // int类型的零值
			t.Errorf("Expected zero value (0), got %v", result)
		}
	})

	t.Run("choose_from_empty_string_slice", func(t *testing.T) {
		// 测试空字符串切片
		var empty []string
		result := Choose(empty)
		if result != "" { // string类型的零值
			t.Errorf("Expected zero value (empty string), got %v", result)
		}
	})

	t.Run("choose_from_single_element", func(t *testing.T) {
		// 测试单元素切片
		single := []int{42}
		result := Choose(single)
		if result != 42 {
			t.Errorf("Expected 42, got %v", result)
		}
	})

	t.Run("choose_from_single_string", func(t *testing.T) {
		// 测试单字符串切片
		single := []string{"hello"}
		result := Choose(single)
		if result != "hello" {
			t.Errorf("Expected 'hello', got %v", result)
		}
	})

	t.Run("choose_from_multiple_integers", func(t *testing.T) {
		// 测试多元素整数切片
		numbers := []int{1, 2, 3, 4, 5}

		// 多次测试确保结果在预期范围内
		for i := 0; i < 100; i++ {
			result := Choose(numbers)
			found := false
			for _, num := range numbers {
				if result == num {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Result %v not found in source slice %v", result, numbers)
			}
		}
	})

	t.Run("choose_from_multiple_strings", func(t *testing.T) {
		// 测试多元素字符串切片
		words := []string{"apple", "banana", "cherry", "date"}

		// 多次测试确保结果在预期范围内
		for i := 0; i < 50; i++ {
			result := Choose(words)
			found := false
			for _, word := range words {
				if result == word {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Result %v not found in source slice %v", result, words)
			}
		}
	})

	t.Run("choose_from_float_slice", func(t *testing.T) {
		// 测试浮点数切片
		floats := []float64{1.1, 2.2, 3.3, 4.4}
		result := Choose(floats)

		found := false
		for _, f := range floats {
			if result == f {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Result %v not found in source slice %v", result, floats)
		}
	})

	t.Run("choose_from_struct_slice", func(t *testing.T) {
		// 测试结构体切片
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
		}

		result := Choose(people)
		found := false
		for _, person := range people {
			if result == person {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Result %v not found in source slice %v", result, people)
		}
	})

	t.Run("choose_distribution_test", func(t *testing.T) {
		// 测试随机分布（统计测试）
		numbers := []int{1, 2, 3}
		counts := make(map[int]int)

		// 运行大量测试来验证分布
		iterations := 3000
		for i := 0; i < iterations; i++ {
			result := Choose(numbers)
			counts[result]++
		}

		// 每个数字应该大概出现1000次（允许一定的偏差）
		for _, num := range numbers {
			count := counts[num]
			expectedMin := iterations/len(numbers) - 200 // 800
			expectedMax := iterations/len(numbers) + 200 // 1200

			if count < expectedMin || count > expectedMax {
				t.Logf("Warning: Number %d appeared %d times, expected around %d",
					num, count, iterations/len(numbers))
			}
		}

		// 验证所有数字都被选中过
		for _, num := range numbers {
			if counts[num] == 0 {
				t.Errorf("Number %d was never selected", num)
			}
		}
	})

	t.Run("choose_large_slice", func(t *testing.T) {
		// 测试大切片
		large := make([]int, 1000)
		for i := range large {
			large[i] = i
		}

		result := Choose(large)
		if result < 0 || result >= 1000 {
			t.Errorf("Result %v is out of expected range [0, 999]", result)
		}
	})

	t.Run("choose_with_duplicate_values", func(t *testing.T) {
		// 测试包含重复值的切片
		duplicates := []int{1, 1, 2, 2, 2, 3}

		for i := 0; i < 100; i++ {
			result := Choose(duplicates)
			if result < 1 || result > 3 {
				t.Errorf("Result %v is out of expected range [1, 3]", result)
			}
		}
	})
}
