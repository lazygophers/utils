package validator

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 全面性能测试 - 测试不同场景
func TestComprehensivePerf(t *testing.T) {
	fmt.Println("\n================== 全面性能测试 ==================")

	// 场景1: 非短路 And（所有验证器都执行）
	fmt.Println("\n--- 场景1: 非短路 And（所有验证器都执行） ---")
	testAndNoShortCircuit(t)

	// 场景2: 大量验证器（20个）
	fmt.Println("\n--- 场景2: 大量验证器 And（20个） ---")
	testAndManyValidators(t)

	// 场景3: 非短路 Or
	fmt.Println("\n--- 场景3: 非短路 Or（所有验证器都执行） ---")
	testOrNoShortCircuit(t)

	// 场景4: 大量验证器 Or（20个）
	fmt.Println("\n--- 场景4: 大量验证器 Or（20个） ---")
	testOrManyValidators(t)

	fmt.Println("\n================== 测试完成 ==================")
}

func testAndNoShortCircuit(t *testing.T) {
	validators := make([]ValidatorFunc, 5)
	for i := 0; i < 4; i++ {
		validators[i] = func(fl FieldLevel) bool { return true }
	}
	validators[4] = func(fl FieldLevel) bool { return false }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := And(validators...)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := AndIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

func testAndManyValidators(t *testing.T) {
	validators := make([]ValidatorFunc, 20)
	validators[0] = func(fl FieldLevel) bool { return false }
	for i := 1; i < 20; i++ {
		validators[i] = func(fl FieldLevel) bool { return true }
	}

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := And(validators...)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := AndIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

func testOrNoShortCircuit(t *testing.T) {
	validators := make([]ValidatorFunc, 5)
	for i := 0; i < 4; i++ {
		validators[i] = func(fl FieldLevel) bool { return false }
	}
	validators[4] = func(fl FieldLevel) bool { return true }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := Or(validators...)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := OrIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

func testOrManyValidators(t *testing.T) {
	validators := make([]ValidatorFunc, 20)
	validators[0] = func(fl FieldLevel) bool { return true }
	for i := 1; i < 20; i++ {
		validators[i] = func(fl FieldLevel) bool { return false }
	}

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := Or(validators...)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := OrIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}
