package candy

import (
	"testing"
)

// TestMax 测试 Max 函数
func TestMax(t *testing.T) {
	t.Run("empty slice returns zero value", func(t *testing.T) {
		result := Max([]int{}...)
		if result != 0 {
			t.Errorf("Max([]int{}...) = %d, want 0", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		result := Max([]int{42}...)
		if result != 42 {
			t.Errorf("Max([42]) = %d, want 42", result)
		}
	})

	t.Run("max of positive integers", func(t *testing.T) {
		result := Max([]int{3, 1, 4, 1, 5, 9, 2, 6}...)
		if result != 9 {
			t.Errorf("Max([3,1,4,1,5,9,2,6]) = %d, want 9", result)
		}
	})

	t.Run("max of negative integers", func(t *testing.T) {
		result := Max([]int{-3, -1, -4, -2}...)
		if result != -1 {
			t.Errorf("Max([-3,-1,-4,-2]) = %d, want -1", result)
		}
	})

	t.Run("max of mixed integers", func(t *testing.T) {
		result := Max([]int{-5, 0, 3, -2, 7}...)
		if result != 7 {
			t.Errorf("Max([-5,0,3,-2,7]) = %d, want 7", result)
		}
	})

	t.Run("max of floats", func(t *testing.T) {
		result := Max([]float64{3.14, 1.618, 2.718, 1.414}...)
		if result != 3.14 {
			t.Errorf("Max([3.14,1.618,2.718,1.414]) = %f, want 3.14", result)
		}
	})

	t.Run("max of strings", func(t *testing.T) {
		result := Max([]string{"apple", "banana", "cherry", "date"}...)
		if result != "date" {
			t.Errorf("Max([apple,banana,cherry,date]) = %s, want date", result)
		}
	})

	t.Run("max at beginning", func(t *testing.T) {
		result := Max([]int{9, 1, 2, 3}...)
		if result != 9 {
			t.Errorf("Max([9,1,2,3]) = %d, want 9", result)
		}
	})

	t.Run("max at end", func(t *testing.T) {
		result := Max([]int{1, 2, 3, 9}...)
		if result != 9 {
			t.Errorf("Max([1,2,3,9]) = %d, want 9", result)
		}
	})

	t.Run("all elements equal", func(t *testing.T) {
		result := Max([]int{5, 5, 5, 5}...)
		if result != 5 {
			t.Errorf("Max([5,5,5,5]) = %d, want 5", result)
		}
	})

	t.Run("max with uint", func(t *testing.T) {
		result := Max([]uint{1, 5, 3, 2}...)
		if result != 5 {
			t.Errorf("Max([1,5,3,2]) = %d, want 5", result)
		}
	})

	t.Run("max with int64", func(t *testing.T) {
		result := Max([]int64{100, 200, 150, 250}...)
		if result != 250 {
			t.Errorf("Max([100,200,150,250]) = %d, want 250", result)
		}
	})

	t.Run("max with float32", func(t *testing.T) {
		result := Max([]float32{1.5, 2.5, 1.2, 3.7}...)
		if result != 3.7 {
			t.Errorf("Max([1.5,2.5,1.2,3.7]) = %f, want 3.7", result)
		}
	})
}

// TestMin 测试 Min 函数
func TestMin(t *testing.T) {
	t.Run("empty slice returns zero value", func(t *testing.T) {
		result := Min([]int{}...)
		if result != 0 {
			t.Errorf("Min([]int{}...) = %d, want 0", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		result := Min([]int{42}...)
		if result != 42 {
			t.Errorf("Min([42]) = %d, want 42", result)
		}
	})

	t.Run("min of positive integers", func(t *testing.T) {
		result := Min([]int{3, 1, 4, 1, 5, 9, 2, 6}...)
		if result != 1 {
			t.Errorf("Min([3,1,4,1,5,9,2,6]) = %d, want 1", result)
		}
	})

	t.Run("min of negative integers", func(t *testing.T) {
		result := Min([]int{-3, -1, -4, -2}...)
		if result != -4 {
			t.Errorf("Min([-3,-1,-4,-2]) = %d, want -4", result)
		}
	})

	t.Run("min of mixed integers", func(t *testing.T) {
		result := Min([]int{-5, 0, 3, -2, 7}...)
		if result != -5 {
			t.Errorf("Min([-5,0,3,-2,7]) = %d, want -5", result)
		}
	})

	t.Run("min of floats", func(t *testing.T) {
		result := Min([]float64{3.14, 1.618, 2.718, 1.414}...)
		if result != 1.414 {
			t.Errorf("Min([3.14,1.618,2.718,1.414]) = %f, want 1.414", result)
		}
	})

	t.Run("min of strings", func(t *testing.T) {
		result := Min([]string{"apple", "banana", "cherry", "date"}...)
		if result != "apple" {
			t.Errorf("Min([apple,banana,cherry,date]) = %s, want apple", result)
		}
	})

	t.Run("min at beginning", func(t *testing.T) {
		result := Min([]int{1, 9, 8, 7}...)
		if result != 1 {
			t.Errorf("Min([1,9,8,7]) = %d, want 1", result)
		}
	})

	t.Run("min at end", func(t *testing.T) {
		result := Min([]int{9, 8, 7, 1}...)
		if result != 1 {
			t.Errorf("Min([9,8,7,1]) = %d, want 1", result)
		}
	})

	t.Run("all elements equal", func(t *testing.T) {
		result := Min([]int{5, 5, 5, 5}...)
		if result != 5 {
			t.Errorf("Min([5,5,5,5]) = %d, want 5", result)
		}
	})

	t.Run("min with uint", func(t *testing.T) {
		result := Min([]uint{5, 1, 3, 2}...)
		if result != 1 {
			t.Errorf("Min([5,1,3,2]) = %d, want 1", result)
		}
	})

	t.Run("min with int64", func(t *testing.T) {
		result := Min([]int64{100, 200, 50, 250}...)
		if result != 50 {
			t.Errorf("Min([100,200,50,250]) = %d, want 50", result)
		}
	})

	t.Run("min with float32", func(t *testing.T) {
		result := Min([]float32{1.5, 2.5, 0.2, 3.7}...)
		if result != 0.2 {
			t.Errorf("Min([1.5,2.5,0.2,3.7]) = %f, want 0.2", result)
		}
	})
}

// TestSum 测试 Sum 函数
func TestSum(t *testing.T) {
	t.Run("empty variadic returns zero", func(t *testing.T) {
		result := Sum[int]()
		if result != 0 {
			t.Errorf("Sum() = %d, want 0", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		result := Sum(42)
		if result != 42 {
			t.Errorf("Sum(42) = %d, want 42", result)
		}
	})

	t.Run("sum of positive integers", func(t *testing.T) {
		result := Sum(1, 2, 3, 4, 5)
		if result != 15 {
			t.Errorf("Sum(1,2,3,4,5) = %d, want 15", result)
		}
	})

	t.Run("sum of negative integers", func(t *testing.T) {
		result := Sum(-1, -2, -3)
		if result != -6 {
			t.Errorf("Sum(-1,-2,-3) = %d, want -6", result)
		}
	})

	t.Run("sum of mixed integers", func(t *testing.T) {
		result := Sum(-5, 10, -3, 8)
		if result != 10 {
			t.Errorf("Sum(-5,10,-3,8) = %d, want 10", result)
		}
	})

	t.Run("sum of floats", func(t *testing.T) {
		result := Sum(1.5, 2.5, 3.0)
		if result != 7.0 {
			t.Errorf("Sum(1.5,2.5,3.0) = %f, want 7.0", result)
		}
	})

	t.Run("sum with zero", func(t *testing.T) {
		result := Sum(0, 0, 0)
		if result != 0 {
			t.Errorf("Sum(0,0,0) = %d, want 0", result)
		}
	})

	t.Run("sum with uint", func(t *testing.T) {
		result := Sum[uint](1, 2, 3, 4)
		if result != 10 {
			t.Errorf("Sum(1,2,3,4) = %d, want 10", result)
		}
	})

	t.Run("sum with int8", func(t *testing.T) {
		result := Sum[int8](10, 20, 30)
		if result != 60 {
			t.Errorf("Sum(10,20,30) = %d, want 60", result)
		}
	})

	t.Run("sum with int16", func(t *testing.T) {
		result := Sum[int16](100, 200, 300)
		if result != 600 {
			t.Errorf("Sum(100,200,300) = %d, want 600", result)
		}
	})

	t.Run("sum with int32", func(t *testing.T) {
		result := Sum[int32](1000, 2000, 3000)
		if result != 6000 {
			t.Errorf("Sum(1000,2000,3000) = %d, want 6000", result)
		}
	})

	t.Run("sum with int64", func(t *testing.T) {
		result := Sum[int64](10000, 20000, 30000)
		if result != 60000 {
			t.Errorf("Sum(10000,20000,30000) = %d, want 60000", result)
		}
	})

	t.Run("sum with uint8", func(t *testing.T) {
		result := Sum[uint8](10, 20, 30)
		if result != 60 {
			t.Errorf("Sum(10,20,30) = %d, want 60", result)
		}
	})

	t.Run("sum with uint16", func(t *testing.T) {
		result := Sum[uint16](100, 200, 300)
		if result != 600 {
			t.Errorf("Sum(100,200,300) = %d, want 600", result)
		}
	})

	t.Run("sum with uint32", func(t *testing.T) {
		result := Sum[uint32](1000, 2000, 3000)
		if result != 6000 {
			t.Errorf("Sum(1000,2000,3000) = %d, want 6000", result)
		}
	})

	t.Run("sum with uint64", func(t *testing.T) {
		result := Sum[uint64](10000, 20000, 30000)
		if result != 60000 {
			t.Errorf("Sum(10000,20000,30000) = %d, want 60000", result)
		}
	})

	t.Run("sum with float32", func(t *testing.T) {
		result := Sum[float32](1.5, 2.5, 3.0)
		if result != 7.0 {
			t.Errorf("Sum(1.5,2.5,3.0) = %f, want 7.0", result)
		}
	})
}

// TestAverage 测试 Average 函数
func TestAverage(t *testing.T) {
	t.Run("empty variadic returns zero", func(t *testing.T) {
		result := Average[int]()
		if result != 0 {
			t.Errorf("Average() = %d, want 0", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		result := Average(42)
		if result != 42 {
			t.Errorf("Average(42) = %d, want 42", result)
		}
	})

	t.Run("average of positive integers", func(t *testing.T) {
		result := Average(1, 2, 3, 4, 5)
		if result != 3 {
			t.Errorf("Average(1,2,3,4,5) = %d, want 3", result)
		}
	})

	t.Run("average of negative integers", func(t *testing.T) {
		result := Average(-2, -4, -6)
		if result != -4 {
			t.Errorf("Average(-2,-4,-6) = %d, want -4", result)
		}
	})

	t.Run("average of mixed integers", func(t *testing.T) {
		result := Average(-5, 5, 0)
		if result != 0 {
			t.Errorf("Average(-5,5,0) = %d, want 0", result)
		}
	})

	t.Run("average of floats", func(t *testing.T) {
		result := Average(1.0, 2.0, 3.0, 4.0)
		if result != 2.5 {
			t.Errorf("Average(1.0,2.0,3.0,4.0) = %f, want 2.5", result)
		}
	})

	t.Run("average with rounding", func(t *testing.T) {
		result := Average(1, 2, 3)
		if result != 2 {
			t.Errorf("Average(1,2,3) = %d, want 2", result)
		}
	})

	t.Run("average with zeros", func(t *testing.T) {
		result := Average(0, 0, 0, 0)
		if result != 0 {
			t.Errorf("Average(0,0,0,0) = %d, want 0", result)
		}
	})

	t.Run("average with uint", func(t *testing.T) {
		result := Average[uint](2, 4, 6, 8)
		if result != 5 {
			t.Errorf("Average(2,4,6,8) = %d, want 5", result)
		}
	})

	t.Run("average with int8", func(t *testing.T) {
		result := Average[int8](10, 20, 30)
		if result != 20 {
			t.Errorf("Average(10,20,30) = %d, want 20", result)
		}
	})

	t.Run("average with int16", func(t *testing.T) {
		result := Average[int16](100, 200, 300)
		if result != 200 {
			t.Errorf("Average(100,200,300) = %d, want 200", result)
		}
	})

	t.Run("average with int32", func(t *testing.T) {
		result := Average[int32](1000, 2000, 3000)
		if result != 2000 {
			t.Errorf("Average(1000,2000,3000) = %d, want 2000", result)
		}
	})

	t.Run("average with int64", func(t *testing.T) {
		result := Average[int64](10000, 20000, 30000)
		if result != 20000 {
			t.Errorf("Average(10000,20000,30000) = %d, want 20000", result)
		}
	})

	t.Run("average with uint8", func(t *testing.T) {
		result := Average[uint8](10, 20, 30)
		if result != 20 {
			t.Errorf("Average(10,20,30) = %d, want 20", result)
		}
	})

	t.Run("average with uint16", func(t *testing.T) {
		result := Average[uint16](100, 200, 300)
		if result != 200 {
			t.Errorf("Average(100,200,300) = %d, want 200", result)
		}
	})

	t.Run("average with uint32", func(t *testing.T) {
		result := Average[uint32](1000, 2000, 3000)
		if result != 2000 {
			t.Errorf("Average(1000,2000,3000) = %d, want 2000", result)
		}
	})

	t.Run("average with uint64", func(t *testing.T) {
		result := Average[uint64](10000, 20000, 30000)
		if result != 20000 {
			t.Errorf("Average(10000,20000,30000) = %d, want 20000", result)
		}
	})

	t.Run("average with float32", func(t *testing.T) {
		result := Average[float32](1.0, 2.0, 3.0)
		if result != 2.0 {
			t.Errorf("Average(1.0,2.0,3.0) = %f, want 2.0", result)
		}
	})

	t.Run("average of two elements", func(t *testing.T) {
		result := Average(10, 20)
		if result != 15 {
			t.Errorf("Average(10,20) = %d, want 15", result)
		}
	})
}

// TestAbs 测试 Abs 函数
func TestAbs(t *testing.T) {
	t.Run("positive integer", func(t *testing.T) {
		result := Abs(42)
		if result != 42 {
			t.Errorf("Abs(42) = %d, want 42", result)
		}
	})

	t.Run("negative integer", func(t *testing.T) {
		result := Abs(-42)
		if result != 42 {
			t.Errorf("Abs(-42) = %d, want 42", result)
		}
	})

	t.Run("zero", func(t *testing.T) {
		result := Abs(0)
		if result != 0 {
			t.Errorf("Abs(0) = %d, want 0", result)
		}
	})

	t.Run("positive float", func(t *testing.T) {
		result := Abs(3.14)
		if result != 3.14 {
			t.Errorf("Abs(3.14) = %f, want 3.14", result)
		}
	})

	t.Run("negative float", func(t *testing.T) {
		result := Abs(-3.14)
		if result != 3.14 {
			t.Errorf("Abs(-3.14) = %f, want 3.14", result)
		}
	})

	t.Run("abs with int8", func(t *testing.T) {
		result := Abs[int8](-127)
		if result != 127 {
			t.Errorf("Abs(-127) = %d, want 127", result)
		}
	})

	t.Run("abs with int16", func(t *testing.T) {
		result := Abs[int16](-1000)
		if result != 1000 {
			t.Errorf("Abs(-1000) = %d, want 1000", result)
		}
	})

	t.Run("abs with int32", func(t *testing.T) {
		result := Abs[int32](-100000)
		if result != 100000 {
			t.Errorf("Abs(-100000) = %d, want 100000", result)
		}
	})

	t.Run("abs with int64", func(t *testing.T) {
		result := Abs[int64](-9223372036854775807)
		if result != 9223372036854775807 {
			t.Errorf("Abs large negative failed")
		}
	})

	t.Run("abs with uint returns same", func(t *testing.T) {
		result := Abs[uint](42)
		if result != 42 {
			t.Errorf("Abs(42) = %d, want 42", result)
		}
	})

	t.Run("abs with uint8", func(t *testing.T) {
		result := Abs[uint8](255)
		if result != 255 {
			t.Errorf("Abs(255) = %d, want 255", result)
		}
	})

	t.Run("abs with uint16", func(t *testing.T) {
		result := Abs[uint16](1000)
		if result != 1000 {
			t.Errorf("Abs(1000) = %d, want 1000", result)
		}
	})

	t.Run("abs with uint32", func(t *testing.T) {
		result := Abs[uint32](100000)
		if result != 100000 {
			t.Errorf("Abs(100000) = %d, want 100000", result)
		}
	})

	t.Run("abs with uint64", func(t *testing.T) {
		result := Abs[uint64](18446744073709551615)
		if result != 18446744073709551615 {
			t.Errorf("Abs large uint64 failed")
		}
	})

	t.Run("abs with float32", func(t *testing.T) {
		result := Abs[float32](-1.5)
		if result != 1.5 {
			t.Errorf("Abs(-1.5) = %f, want 1.5", result)
		}
	})

	t.Run("abs with positive float64", func(t *testing.T) {
		result := Abs(2.718)
		if result != 2.718 {
			t.Errorf("Abs(2.718) = %f, want 2.718", result)
		}
	})

	t.Run("abs with negative float64", func(t *testing.T) {
		result := Abs(-2.718)
		if result != 2.718 {
			t.Errorf("Abs(-2.718) = %f, want 2.718", result)
		}
	})

	t.Run("abs of one", func(t *testing.T) {
		result := Abs(1)
		if result != 1 {
			t.Errorf("Abs(1) = %d, want 1", result)
		}
	})

	t.Run("abs of negative one", func(t *testing.T) {
		result := Abs(-1)
		if result != 1 {
			t.Errorf("Abs(-1) = %d, want 1", result)
		}
	})
}
