package candy

import (
	"testing"
)

// ============== 测试数据 ==============
var (
	sliceSmall  = make([]int, 10)
	sliceMedium = make([]int, 100)
	sliceLarge  = make([]int, 1000)
)

func init() {
	for i := range sliceSmall {
		sliceSmall[i] = i
	}
	for i := range sliceMedium {
		sliceMedium[i] = i
	}
	for i := range sliceLarge {
		sliceLarge[i] = i
	}
}

// ============== First 函数的优化方案 ==============
func FirstOpt1[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}
	return ss[0]
}

func FirstOpt2[T any](ss []T) (ret T) {
	if len(ss) > 0 {
		return ss[0]
	}
	return
}

func FirstOpt3[T any](ss []T) T {
	var zero T
	if len(ss) == 0 {
		return zero
	}
	return ss[0]
}

func FirstOpt4[T any](ss []T) T {
	if len(ss) != 0 {
		return ss[0]
	}
	var zero T
	return zero
}

func FirstOpt5[T any](ss []T) T {
	var zero T
	l := len(ss)
	if l == 0 {
		return zero
	}
	return ss[0]
}

// ============== Last 函数的优化方案 ==============
func LastOpt1[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}
	return ss[len(ss)-1]
}

func LastOpt2[T any](ss []T) (ret T) {
	l := len(ss)
	if l == 0 {
		return
	}
	return ss[l-1]
}

func LastOpt3[T any](ss []T) T {
	var zero T
	l := len(ss)
	if l == 0 {
		return zero
	}
	return ss[l-1]
}

func LastOpt4[T any](ss []T) T {
	if len(ss) != 0 {
		return ss[len(ss)-1]
	}
	var zero T
	return zero
}

// ============== FirstOr 函数的优化方案 ==============
func FirstOrOpt1[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}
	return ss[0]
}

func FirstOrOpt2[T any](ss []T, or T) T {
	l := len(ss)
	if l == 0 {
		return or
	}
	return ss[0]
}

func FirstOrOpt3[T any](ss []T, or T) T {
	if len(ss) != 0 {
		return ss[0]
	}
	return or
}

func FirstOrOpt4[T any](ss []T, or T) (ret T) {
	ret = or
	if len(ss) > 0 {
		ret = ss[0]
	}
	return
}

// ============== LastOr 函数的优化方案 ==============
func LastOrOpt1[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}
	return ss[len(ss)-1]
}

func LastOrOpt2[T any](ss []T, or T) T {
	l := len(ss)
	if l == 0 {
		return or
	}
	return ss[l-1]
}

func LastOrOpt3[T any](ss []T, or T) T {
	if len(ss) != 0 {
		return ss[len(ss)-1]
	}
	return or
}

// ============== Top 函数的优化方案 ==============
func TopOpt1[T any](ss []T, n int) (ret []T) {
	if n <= 0 {
		return []T{}
	}
	if n > len(ss) {
		n = len(ss)
	}
	ret = make([]T, n)
	copy(ret, ss[:n])
	return ret
}

func TopOpt2[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	ret := make([]T, n)
	copy(ret, ss[:n])
	return ret
}

func TopOpt3[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	ret := make([]T, 0, n)
	ret = append(ret, ss[:n]...)
	return ret
}

func TopOpt4[T any](ss []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	l := len(ss)
	if n > l {
		n = l
	}
	return ss[:n:n]
}

// ============== Bottom 函数的优化方案 ==============
func BottomOpt1[T any](ss []T, n int) (ret []T) {
	if n <= 0 {
		return []T{}
	}
	if n > len(ss) {
		n = len(ss)
	}
	ret = make([]T, n)
	copy(ret, ss[len(ss)-n:])
	return ret
}

func BottomOpt2[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	ret := make([]T, n)
	copy(ret, ss[l-n:])
	return ret
}

func BottomOpt3[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	ret := make([]T, 0, n)
	ret = append(ret, ss[l-n:]...)
	return ret
}

func BottomOpt4[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	return ss[l-n : l : l]
}

// ============== First 基准测试 ==============
func BenchmarkFirstOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOpt1(sliceMedium)
	}
}

func BenchmarkFirstOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOpt2(sliceMedium)
	}
}

func BenchmarkFirstOpt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOpt3(sliceMedium)
	}
}

func BenchmarkFirstOpt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOpt4(sliceMedium)
	}
}

func BenchmarkFirstOpt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOpt5(sliceMedium)
	}
}

// ============== Last 基准测试 ==============
func BenchmarkLastOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOpt1(sliceMedium)
	}
}

func BenchmarkLastOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOpt2(sliceMedium)
	}
}

func BenchmarkLastOpt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOpt3(sliceMedium)
	}
}

func BenchmarkLastOpt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOpt4(sliceMedium)
	}
}

// ============== FirstOr 基准测试 ==============
func BenchmarkFirstOrOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOrOpt1(sliceMedium, 999)
	}
}

func BenchmarkFirstOrOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOrOpt2(sliceMedium, 999)
	}
}

func BenchmarkFirstOrOpt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOrOpt3(sliceMedium, 999)
	}
}

func BenchmarkFirstOrOpt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstOrOpt4(sliceMedium, 999)
	}
}

// ============== LastOr 基准测试 ==============
func BenchmarkLastOrOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOrOpt1(sliceMedium, 999)
	}
}

func BenchmarkLastOrOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOrOpt2(sliceMedium, 999)
	}
}

func BenchmarkLastOrOpt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastOrOpt3(sliceMedium, 999)
	}
}

// ============== Top 基准测试 ==============
func BenchmarkTopOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TopOpt1(sliceMedium, 10)
	}
}

func BenchmarkTopOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TopOpt2(sliceMedium, 10)
	}
}

func BenchmarkTopOpt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TopOpt3(sliceMedium, 10)
	}
}

func BenchmarkTopOpt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TopOpt4(sliceMedium, 10)
	}
}

// ============== Bottom 基准测试 ==============
func BenchmarkBottomOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BottomOpt1(sliceMedium, 10)
	}
}

func BenchmarkBottomOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BottomOpt2(sliceMedium, 10)
	}
}

func BenchmarkBottomOpt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BottomOpt3(sliceMedium, 10)
	}
}

func BenchmarkBottomOpt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BottomOpt4(sliceMedium, 10)
	}
}

// ==================== First 基准测试 ====================

func BenchmarkFirst_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

// ==================== FirstOr 基准测试 ====================

func BenchmarkFirstOr_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, "default")
	}
}

// ==================== Last 基准测试 ====================

func BenchmarkLast_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

// ==================== LastOr 基准测试 ====================

func BenchmarkLastOr_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, "default")
	}
}

// ==================== Bottom 基准测试 ====================

func BenchmarkBottom_Int_Small_N1_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 1)
	}
}

func BenchmarkBottom_Int_Small_N5_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 5)
	}
}

func BenchmarkBottom_Int_Medium_N10_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 10)
	}
}

func BenchmarkBottom_Int_Large_N100_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 100)
	}
}

func BenchmarkBottom_Int_Large_N1000_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 1000)
	}
}

func BenchmarkBottom_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 10)
	}
}
