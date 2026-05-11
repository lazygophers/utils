package defaults

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// ========== 优化方案实现 ==========

// 方案1: 基线版本
func benchBaselineParseSlice(vv reflect.Value, defaultStr string) error {
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
			case reflect.Float64:
				val, _ := strconv.ParseFloat(part, 64)
				slice.Index(i).SetFloat(val)
			}
		}
		vv.Set(slice)
		return nil
	}
	return fmt.Errorf("parse error")
}

// 方案2: 预检查优化
func benchV2ParseSlice(vv reflect.Value, defaultStr string) error {
	if !strings.HasPrefix(defaultStr, "[") {
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
				case reflect.Float64:
					val, _ := strconv.ParseFloat(part, 64)
					slice.Index(i).SetFloat(val)
				}
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

// 方案3: int特化
func benchV3ParseSlice(vv reflect.Value, defaultStr string) error {
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
	return benchBaselineParseSlice(vv, defaultStr)
}

// 方案4: string+int特化
func benchV4ParseSlice(vv reflect.Value, defaultStr string) error {
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
	return benchBaselineParseSlice(vv, defaultStr)
}

// 方案5: 预分配容量
func benchV5ParseSlice(vv reflect.Value, defaultStr string) error {
	if !strings.HasPrefix(defaultStr, "[") && strings.Contains(defaultStr, ",") {
		estimatedCap := strings.Count(defaultStr, ",") + 1
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), estimatedCap)
		for i, part := range parts {
			part = strings.TrimSpace(part)
			switch vv.Type().Elem().Kind() {
			case reflect.Int:
				val, _ := strconv.ParseInt(part, 10, 64)
				slice.Index(i).SetInt(val)
			case reflect.String:
				slice.Index(i).SetString(part)
			case reflect.Float64:
				val, _ := strconv.ParseFloat(part, 64)
				slice.Index(i).SetFloat(val)
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

// 方案10: 综合优化
func benchV10ParseSlice(vv reflect.Value, defaultStr string) error {
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
		if elemType.Kind() == reflect.Float64 {
			parts := strings.Split(defaultStr, ",")
			slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
			for i, part := range parts {
				val, _ := strconv.ParseFloat(strings.TrimSpace(part), 64)
				slice.Index(i).SetFloat(val)
			}
			vv.Set(slice)
			return nil
		}
		parts := strings.Split(defaultStr, ",")
		slice := reflect.MakeSlice(vv.Type(), len(parts), len(parts))
		for i, part := range parts {
			part = strings.TrimSpace(part)
			switch elemType.Kind() {
			case reflect.Int:
				val, _ := strconv.ParseInt(part, 10, 64)
				slice.Index(i).SetInt(val)
			case reflect.String:
				slice.Index(i).SetString(part)
			case reflect.Float64:
				val, _ := strconv.ParseFloat(part, 64)
				slice.Index(i).SetFloat(val)
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

// Map优化方案
func benchBaselineParseMap(vv reflect.Value, defaultStr string) error {
	if strings.HasPrefix(defaultStr, "{") && strings.HasSuffix(defaultStr, "}") {
		mapPtr := reflect.New(vv.Type())
		if err := json.Unmarshal([]byte(defaultStr), mapPtr.Interface()); err == nil {
			vv.Set(mapPtr.Elem())
			return nil
		}
	}
	return fmt.Errorf("parse error")
}

func benchV8ParseMap(vv reflect.Value, defaultStr string) error {
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

func BenchmarkSlice_Baseline_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchBaselineParseSlice(vv, "1,2,3,4,5")
	}
}

func BenchmarkSlice_Baseline_String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []string
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchBaselineParseSlice(vv, "a,b,c,d,e")
	}
}

func BenchmarkSlice_Baseline_Float(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []float64
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchBaselineParseSlice(vv, "1.1,2.2,3.3")
	}
}

func BenchmarkSlice_Baseline_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchBaselineParseSlice(vv, "[100,200,300]")
	}
}

func BenchmarkSlice_V2_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV2ParseSlice(vv, "1,2,3,4,5")
	}
}

func BenchmarkSlice_V3_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV3ParseSlice(vv, "1,2,3,4,5")
	}
}

func BenchmarkSlice_V4_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV4ParseSlice(vv, "1,2,3,4,5")
	}
}

func BenchmarkSlice_V4_String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []string
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV4ParseSlice(vv, "a,b,c,d,e")
	}
}

func BenchmarkSlice_V5_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV5ParseSlice(vv, "1,2,3,4,5")
	}
}

func BenchmarkSlice_V10_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV10ParseSlice(vv, "1,2,3,4,5")
	}
}

func BenchmarkSlice_V10_String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []string
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV10ParseSlice(vv, "a,b,c,d,e")
	}
}

func BenchmarkSlice_V10_Float(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []float64
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV10ParseSlice(vv, "1.1,2.2,3.3")
	}
}

func BenchmarkSlice_V10_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var slice []int
		vv := reflect.ValueOf(&slice).Elem()
		_ = benchV10ParseSlice(vv, "[100,200,300]")
	}
}

func BenchmarkMap_Baseline_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		vv := reflect.ValueOf(&m).Elem()
		_ = benchBaselineParseMap(vv, "{\"key1\":\"val1\",\"key2\":\"val2\"}")
	}
}

func BenchmarkMap_V8_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		vv := reflect.ValueOf(&m).Elem()
		_ = benchV8ParseMap(vv, "{\"key1\":\"val1\",\"key2\":\"val2\"}")
	}
}
