package validator

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 手动性能测试 - And 函数对比
func TestManualAndPerf(t *testing.T) {
	// 创建测试验证器
	v1 := func(fl FieldLevel) bool { return false }
	v2 := func(fl FieldLevel) bool { return true }
	v3 := func(fl FieldLevel) bool { return true }
	v4 := func(fl FieldLevel) bool { return true }
	v5 := func(fl FieldLevel) bool { return true }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 测试原始版本（range）
	vOriginal := And(v1, v2, v3, v4, v5)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 测试索引循环版本
	vIndex := AndIndexLoop(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	// 测试 switch 版本
	vSwitch := AndSwitch(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vSwitch(fl)
	}
	switchTime := time.Since(start)

	// 测试 goto 版本
	vGoto := AndGoto(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vGoto(fl)
	}
	gotoTime := time.Since(start)

	// 测试结构体版本
	vStruct := AndStruct(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vStruct(fl)
	}
	structTime := time.Since(start)

	// 输出结果
	fmt.Println("\n========================================")
	fmt.Println("And 函数性能测试结果 (短路场景，5个验证器，1000万次迭代):")
	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
	fmt.Printf("Switch展开:       %v (%.1f%%)\n", switchTime, float64(switchTime)*100/float64(originalTime))
	fmt.Printf("Goto优化:         %v (%.1f%%)\n", gotoTime, float64(gotoTime)*100/float64(originalTime))
	fmt.Printf("结构体方法:       %v (%.1f%%)\n", structTime, float64(structTime)*100/float64(originalTime))
	fmt.Println("========================================")
}

// 手动性能测试 - Or 函数对比
func TestManualOrPerf(t *testing.T) {
	// 创建测试验证器
	v1 := func(fl FieldLevel) bool { return true }
	v2 := func(fl FieldLevel) bool { return false }
	v3 := func(fl FieldLevel) bool { return false }
	v4 := func(fl FieldLevel) bool { return false }
	v5 := func(fl FieldLevel) bool { return false }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 测试原始版本（range）
	vOriginal := Or(v1, v2, v3, v4, v5)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 测试索引循环版本
	vIndex := OrIndexLoop(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	// 测试 switch 版本
	vSwitch := OrSwitch(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vSwitch(fl)
	}
	switchTime := time.Since(start)

	// 测试 goto 版本
	vGoto := OrGoto(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vGoto(fl)
	}
	gotoTime := time.Since(start)

	// 测试结构体版本
	vStruct := OrStruct(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vStruct(fl)
	}
	structTime := time.Since(start)

	// 输出结果
	fmt.Println("\n========================================")
	fmt.Println("Or 函数性能测试结果 (短路场景，5个验证器，1000万次迭代):")
	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
	fmt.Printf("Switch展开:       %v (%.1f%%)\n", switchTime, float64(switchTime)*100/float64(originalTime))
	fmt.Printf("Goto优化:         %v (%.1f%%)\n", gotoTime, float64(gotoTime)*100/float64(originalTime))
	fmt.Printf("结构体方法:       %v (%.1f%%)\n", structTime, float64(structTime)*100/float64(originalTime))
	fmt.Println("========================================")
}

// 手动性能测试 - Not 函数对比
func TestManualNotPerf(t *testing.T) {
	// 创建测试验证器
	v1 := func(fl FieldLevel) bool { return true }
	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 测试原始版本
	vOriginal := Not(v1)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 输出结果
	fmt.Println("\n========================================")
	fmt.Println("Not 函数性能测试结果 (1000万次迭代):")
	fmt.Printf("原始版本:  %v\n", originalTime)
	fmt.Println("结论: Not 函数已经是最优实现（仅一个取反操作）")
	fmt.Println("========================================")
}
