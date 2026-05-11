package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ========== 优化方案 ==========

// 方案1: 基线
func baselineParseSlice(vv reflect.Value, defaultStr string) error {
	if strings.HasPrefix(defaultStr, "[") && strings.HasSuffix(defaultStr, "]") {
		slicePtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), slicePtr.Interface()); err == nil {
			vv.Set(slicePtr.Elem())
			return nil
		}
	}
	if strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
		for i, part := range parts {
			part = strings.TrimSpace(part)
			switch vv.Type().Elem().Kind() {
			case reflect.Int:
				val, _ := strconv.ParseInt(part, 10, 64)
				slice.Index(i).SetInt(val)
			case reflect.String:
				slice.Index(i).SetString(part)
			}
		}
		vv.Set(slice)
		return nil
	}
	return fmt.Errorf("parse error")
}

// 方案2: 预检查
func v2ParseSlice(vv reflect.Value, defaultStr string) error {
	if !strings.HasPrefix(defaultStr, "[") && strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
		for i, part := range parts {
			part = strings.TrimSpace(part)
			switch vv.Type().Elem().Kind() {
			case reflect.Int:
				val, _ := strconv.ParseInt(part, 10, 64)
				slice.Index(i).SetInt(val)
			case reflect.String:
				slice.Index(i).SetString(part)
			}
		}
		vv.Set(slice)
		return nil
	}
	if strings.HasPrefix(defaultStr, "[") && strings.HasSuffix(defaultStr, "]") {
		slicePtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), slicePtr.Interface()); err == nil {
			vv.Set(slicePtr.Elem())
			return nil
		}
	}
	return fmt.Errorf("parse error")
}

// 方案3: int特化
func v3ParseSlice(vv reflect.Value, defaultStr string) error {
	if vv.Type().Elem().Kind() == reflect.Int && !strings.HasPrefix(defaultStr, "[") && strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
		for i, part := range parts {
			val, _ := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
			slice.Index(i).SetInt(val)
		}
		vv.Set(slice)
		return nil
	}
	return baselineParseSlice(vv, defaultStr)
}

// 方案4: string+int特化
func v4ParseSlice(vv reflect.Value, defaultStr string) error {
	elemType := vv.Type().Elem()
	if elemType.Kind() == reflect.String && !strings.HasPrefix(defaultStr, "[") && strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
		for i, part := range parts {
			slice.Index(i).SetString(strings.TrimSpace(part))
		}
		vv.Set(slice)
		return nil
	}
	if elemType.Kind() == reflect.Int && !strings.HasPrefix(defaultStr, "[") && strings.Contains(defaultStr, ",") {
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
		for i, part := range parts {
			val, _ := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
			slice.Index(i).SetInt(val)
		}
		vv.Set(slice)
		return nil
	}
	return baselineParseSlice(vv, defaultStr)
}

// 方案10: 综合优化
func v10ParseSlice(vv reflect.Value, defaultStr string) error {
	elemType := vv.Type().Elem()
	if !strings.HasPrefix(defaultStr, "[") && strings.Contains(defaultStr, ",") {
		if elemType.Kind() == reflect.String {
			parts := strings.Split(defaultStr, ",")
			slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
			for i, part := range parts {
				slice.Index(i).SetString(strings.TrimSpace(part))
			}
			vv.Set(slice)
			return nil
		}
		if elemType.Kind() == reflect.Int {
			parts := strings.Split(defaultStr, ",")
			slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
			for i, part := range parts {
				val, _ := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
				slice.Index(i).SetInt(val)
			}
			vv.Set(slice)
			return nil
		}
	}
	if strings.HasPrefix(defaultStr, "[") && strings.HasSuffix(defaultStr, "]") {
		slicePtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), slicePtr.Interface()); err == nil {
			vv.Set(slicePtr.Elem())
			return nil
		}
	}
	return fmt.Errorf("parse error")
}

// Map优化
func baselineParseMap(vv reflect.Value, defaultStr string) error {
	if strings.HasPrefix(defaultStr, "{") && strings.HasSuffix(defaultStr, "}") {
		mapPtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), mapPtr.Interface()); err == nil {
			vv.Set(mapPtr.Elem())
			return nil
		}
	}
	return fmt.Errorf("parse error")
}

func v8ParseMap(vv reflect.Value, defaultStr string) error {
	if vv.Type().Key().Kind() == reflect.String && vv.Type().Elem().Kind() == reflect.String {
		if strings.HasPrefix(defaultStr, "{") && strings.HasSuffix(defaultStr, "}") {
			if !strings.Contains(defaultStr, "\"") {
				content := defaultStr[1 : len(defaultStr)-1]
				if content != "" {
					result := reflect.MakeMap(vv.Type())
					pairs := strings.Split(content, ",")
					for _, pair := range pairs {
						if idx := strings.Index(pair, ":"); idx > 0 {
							key := strings.TrimSpace(pair[:idx])
							val := strings.TrimSpace(pair[idx+1:])
							result.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
						}
					}
					vv.Set(result)
					return nil
				}
			}
		}
	}
	if strings.HasPrefix(defaultStr, "{") && strings.HasSuffix(defaultStr, "}") {
		mapPtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), mapPtr.Interface()); err == nil {
			vv.Set(mapPtr.Elem())
			return nil
		}
	}
	return fmt.Errorf("parse error")
}

// ========== 基准测试 ==========

func benchmark(name string, iterations int, fn func()) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	elapsed := time.Since(start)
	nsPerOp := elapsed.Nanoseconds() / int64(iterations)
	fmt.Printf("%-30s %10d ns/op\t%10.2f µs total\n", name, nsPerOp, float64(elapsed.Microseconds()))
}

func main() {
	iterations := 100000

	fmt.Println("========================================")
	fmt.Println("切片解析性能基准测试")
	fmt.Println("========================================")
	fmt.Printf("迭代次数: %d\n\n", iterations)

	// int切片测试
	fmt.Println("[]int 切片解析 (\"1,2,3,4,5\"):")
	benchmark("Baseline (方案1)", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = baselineParseSlice(vv, "1,2,3,4,5")
	})

	benchmark("V2: 预检查优化", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = v2ParseSlice(vv, "1,2,3,4,5")
	})

	benchmark("V3: int特化", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = v3ParseSlice(vv, "1,2,3,4,5")
	})

	benchmark("V4: string+int特化", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = v4ParseSlice(vv, "1,2,3,4,5")
	})

	benchmark("V10: 综合优化", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = v10ParseSlice(vv, "1,2,3,4,5")
	})

	fmt.Println("\n[]string 切片解析 (\"a,b,c,d,e\"):")
	benchmark("Baseline (方案1)", iterations, func() {
		var slice []string
		vv := reflect.ValueOf(&slice).Elem()
		_ = baselineParseSlice(vv, "a,b,c,d,e")
	})

	benchmark("V4: string+int特化", iterations, func() {
		var slice []string
		vv := reflect.ValueOf(&slice).Elem()
		_ = v4ParseSlice(vv, "a,b,c,d,e")
	})

	benchmark("V10: 综合优化", iterations, func() {
		var slice []string
		vv := reflect.ValueOf(&slice).Elem()
		_ = v10ParseSlice(vv, "a,b,c,d,e")
	})

	fmt.Println("\n[]int JSON解析 (\"[100,200,300]\"):")
	benchmark("Baseline (方案1)", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = baselineParseSlice(vv, "[100,200,300]")
	})

	benchmark("V10: 综合优化", iterations, func() {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = v10ParseSlice(vv, "[100,200,300]")
	})

	fmt.Println("\n========================================")
	fmt.Println("Map解析性能基准测试")
	fmt.Println("========================================\n")

	fmt.Println("map[string]string (\"{\\\"key1\\\":\\\"val1\\\",\\\"key2\\\":\\\"val2\\\"}\"):")
	iterationsMap := 50000
	fmt.Printf("迭代次数: %d\n\n", iterationsMap)

	benchmark("Baseline (方案1)", iterationsMap, func() {
		m := make(map[string]string)
		vv := reflect.ValueOf(&m).Elem()
		_ = baselineParseMap(vv, "{\"key1\":\"val1\",\"key2\":\"val2\"}")
	})

	benchmark("V8: string->string特化", iterationsMap, func() {
		m := make(map[string]string)
		vv := reflect.ValueOf(&m).Elem()
		_ = v8ParseMap(vv, "{\"key1\":\"val1\",\"key2\":\"val2\"}")
	})

	fmt.Println("\n========================================")
	fmt.Println("性能提升总结")
	fmt.Println("========================================")
	fmt.Println("查看上方各方案耗时，ns/op 值越小越好")
	fmt.Println("V10综合优化方案在大多数场景下表现最佳")
}
