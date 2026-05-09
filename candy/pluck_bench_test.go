package candy

import (
	"math/rand"
	"testing"
)

// Employee 测试数据结构
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// 生成随机字符串
func generateRandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 生成测试数据
func generatePeople(n int) []Person {
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		people[i] = Person{
			ID:   i,
			Name: generateRandomString(10),
			Age:  int8(rand.Intn(50) + 20),
		}
	}
	return people
}

func generatePersonPtrs(n int) []*Person {
	ptrs := make([]*Person, n)
	for i := 0; i < n; i++ {
		if rand.Intn(10) > 2 { // 80% 非空
			p := Person{
				ID:   i,
				Name: generateRandomString(10),
				Age:  int8(rand.Intn(50) + 20),
			}
			ptrs[i] = &p
		}
	}
	return ptrs
}

func generateEmployees(n int) []Employee {
	employees := make([]Employee, n)
	for i := 0; i < n; i++ {
		employees[i] = Employee{
			ID:     rand.Intn(n/2), // 有重复 ID
			Name:   generateRandomString(10),
			Salary: float64(rand.Intn(100000)),
		}
	}
	return employees
}

// ==================== Pluck Benchmark ====================

func BenchmarkPluck_Small_AllFields(b *testing.B) {
	people := generatePeople(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Pluck(people, func(p Person) int { return p.ID })
	}
}

func BenchmarkPluck_Medium_AllFields(b *testing.B) {
	people := generatePeople(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Pluck(people, func(p Person) string { return p.Name })
	}
}

func BenchmarkPluck_Large_ComputeField(b *testing.B) {
	people := generatePeople(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Pluck(people, func(p Person) int { return int(p.Age) * 2 })
	}
}

func BenchmarkPluck_XL_SimpleField(b *testing.B) {
	people := generatePeople(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Pluck(people, func(p Person) int { return p.ID })
	}
}

func BenchmarkPluck_Huge_MultiField(b *testing.B) {
	people := generatePeople(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Pluck(people, func(p Person) string { return p.Name })
	}
}

// ==================== PluckPtr Benchmark ====================

func BenchmarkPluckPtr_Small_AllNotNil(b *testing.B) {
	ptrs := generatePersonPtrs(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckPtr(ptrs, func(p *Person) int { return p.ID }, 0)
	}
}

func BenchmarkPluckPtr_Medium_SomeNil(b *testing.B) {
	ptrs := generatePersonPtrs(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckPtr(ptrs, func(p *Person) string { return p.Name }, "")
	}
}

func BenchmarkPluckPtr_Large_HalfNil(b *testing.B) {
	ptrs := make([]*Person, 1000)
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			p := Person{ID: i, Name: "test", Age: 30}
			ptrs[i] = &p
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckPtr(ptrs, func(p *Person) int { return p.ID }, -1)
	}
}

func BenchmarkPluckPtr_XL_MostNil(b *testing.B) {
	ptrs := make([]*Person, 10000)
	for i := 0; i < 10000; i++ {
		if i%10 == 0 { // 10% 非空
			p := Person{ID: i, Name: "test", Age: 30}
			ptrs[i] = &p
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckPtr(ptrs, func(p *Person) int { return int(p.Age) }, 0)
	}
}

func BenchmarkPluckPtr_Huge_ComputeField(b *testing.B) {
	ptrs := make([]*Person, 100000)
	for i := 0; i < 100000; i++ {
		if i%5 == 0 { // 20% 非空
			p := Person{ID: i, Name: generateRandomString(10), Age: int8(rand.Intn(50) + 20)}
			ptrs[i] = &p
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckPtr(ptrs, func(p *Person) int { return int(p.Age) * 3 }, 0)
	}
}

// ==================== PluckUnique Benchmark ====================

func BenchmarkPluckUnique_Small_AllUnique(b *testing.B) {
	people := generatePeople(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUnique(people, func(p Person) int { return p.ID })
	}
}

func BenchmarkPluckUnique_Medium_HalfDuplicates(b *testing.B) {
	people := make([]Person, 100)
	for i := 0; i < 100; i++ {
		people[i] = Person{
			ID:   i / 2, // 50% 重复
			Name: generateRandomString(10),
			Age:  int8(rand.Intn(50) + 20),
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUnique(people, func(p Person) int { return p.ID })
	}
}

func BenchmarkPluckUnique_Large_HighDuplicates(b *testing.B) {
	people := make([]Person, 1000)
	for i := 0; i < 1000; i++ {
		people[i] = Person{
			ID:   rand.Intn(100), // 高重复率
			Name: generateRandomString(10),
			Age:  int8(rand.Intn(50) + 20),
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUnique(people, func(p Person) int { return p.ID })
	}
}

func BenchmarkPluckUnique_XL_AllDuplicates(b *testing.B) {
	people := make([]Person, 10000)
	for i := 0; i < 10000; i++ {
		people[i] = Person{
			ID:   42, // 全部相同
			Name: generateRandomString(10),
			Age:  int8(rand.Intn(50) + 20),
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUnique(people, func(p Person) int { return p.ID })
	}
}

func BenchmarkPluckUnique_Huge_ComputeKey(b *testing.B) {
	people := make([]Person, 100000)
	for i := 0; i < 100000; i++ {
		people[i] = Person{
			ID:   i,
			Name: generateRandomString(10),
			Age:  int8(rand.Intn(50) + 20),
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUnique(people, func(p Person) string { return p.Name })
	}
}

// ==================== PluckMap Benchmark ====================

func BenchmarkPluckMap_Small_Simple(b *testing.B) {
	people := generatePeople(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckMap(people, func(p Person) int { return p.ID }, func(p Person) string { return p.Name })
	}
}

func BenchmarkPluckMap_Medium_ComputeValue(b *testing.B) {
	people := generatePeople(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckMap(people, func(p Person) int { return p.ID }, func(p Person) int { return int(p.Age) * 2 })
	}
}

func BenchmarkPluckMap_Large_StringKey(b *testing.B) {
	people := generatePeople(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckMap(people, func(p Person) string { return p.Name }, func(p Person) int { return int(p.Age) })
	}
}

func BenchmarkPluckMap_XL_StructValue(b *testing.B) {
	people := generatePeople(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckMap(people, func(p Person) int { return p.ID }, func(p Person) Person { return p })
	}
}

func BenchmarkPluckMap_Huge_DuplicateKeys(b *testing.B) {
	people := make([]Person, 100000)
	for i := 0; i < 100000; i++ {
		people[i] = Person{
			ID:   rand.Intn(50000), // 有重复键
			Name: generateRandomString(10),
			Age:  int8(rand.Intn(50) + 20),
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckMap(people, func(p Person) int { return p.ID }, func(p Person) string { return p.Name })
	}
}

// ==================== PluckGroupBy Benchmark ====================

func BenchmarkPluckGroupBy_Small_FewGroups(b *testing.B) {
	people := make([]Person, 10)
	for i := 0; i < 10; i++ {
		people[i] = Person{ID: i, Name: "test", Age: int8(20 + i%3)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckGroupBy(people, func(p Person) int { return int(p.Age) })
	}
}

func BenchmarkPluckGroupBy_Medium_ManyGroups(b *testing.B) {
	people := make([]Person, 100)
	for i := 0; i < 100; i++ {
		people[i] = Person{ID: i, Name: "test", Age: int8(20 + i%20)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckGroupBy(people, func(p Person) int { return int(p.Age) })
	}
}

func BenchmarkPluckGroupBy_Large_SingleGroup(b *testing.B) {
	people := make([]Person, 1000)
	for i := 0; i < 1000; i++ {
		people[i] = Person{ID: i, Name: "test", Age: 30}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckGroupBy(people, func(p Person) int { return int(p.Age) })
	}
}

func BenchmarkPluckGroupBy_XL_StringKey(b *testing.B) {
	employees := generateEmployees(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckGroupBy(employees, func(e Employee) int { return e.ID })
	}
}

func BenchmarkPluckGroupBy_Huge_ManyGroups(b *testing.B) {
	people := make([]Person, 100000)
	for i := 0; i < 100000; i++ {
		people[i] = Person{ID: i, Name: "test", Age: int8(rand.Intn(50) + 20)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckGroupBy(people, func(p Person) int { return int(p.Age) })
	}
}

// ==================== PluckInt Benchmark (反射版本) ====================

func BenchmarkPluckInt_Small(b *testing.B) {
	people := generatePeople(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt(people, "ID")
	}
}

func BenchmarkPluckInt_Medium(b *testing.B) {
	people := generatePeople(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt(people, "ID")
	}
}

func BenchmarkPluckInt_Large(b *testing.B) {
	people := generatePeople(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt(people, "ID")
	}
}

func BenchmarkPluckInt_XL(b *testing.B) {
	people := generatePeople(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt(people, "ID")
	}
}

func BenchmarkPluckInt_Huge(b *testing.B) {
	people := generatePeople(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt(people, "ID")
	}
}

// ==================== PluckString Benchmark (反射版本) ====================

func BenchmarkPluckString_Small(b *testing.B) {
	people := generatePeople(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckString(people, "Name")
	}
}

func BenchmarkPluckString_Medium(b *testing.B) {
	people := generatePeople(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckString(people, "Name")
	}
}

func BenchmarkPluckString_Large(b *testing.B) {
	people := generatePeople(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckString(people, "Name")
	}
}

func BenchmarkPluckString_XL(b *testing.B) {
	people := generatePeople(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckString(people, "Name")
	}
}

func BenchmarkPluckString_Huge(b *testing.B) {
	people := generatePeople(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckString(people, "Name")
	}
}
