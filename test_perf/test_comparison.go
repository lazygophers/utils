package main

import (
	"fmt"
	"time"
	"github.com/lazygophers/utils/candy"
)

type benchStruct struct {
	ID   int
	Name string
}

func main() {
	dataSize := 1000
	data := make([]benchStruct, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}

	iterations := 10000

	fmt.Printf("优化后性能测试：\n")
	fmt.Printf("测试 %d 次操作，每次处理 %d 个元素\n\n", iterations, dataSize)

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = candy.KeyByInt(data, "ID")
	}
	elapsed := time.Since(start)
	avgNs := elapsed.Nanoseconds() / int64(iterations)
	fmt.Printf("KeyByInt (优化后): %12v total, %8d ns/op\n", elapsed, avgNs)
}
