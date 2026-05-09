package validator

import (
	"fmt"
	"testing"
)

// BenchmarkValidate_SingleField 单字段验证性能测试
func BenchmarkValidate_SingleField(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name string `validate:"required"`
	}
	user := User{Name: "test"}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(user)
	}
}

// BenchmarkValidate_SingleField_Invalid 单字段验证失败性能测试
func BenchmarkValidate_SingleField_Invalid(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name string `validate:"required"`
	}
	user := User{Name: ""}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(user)
	}
}

// BenchmarkValidate_StringValidation email/url 验证性能测试
func BenchmarkValidate_StringValidation(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("Email", func(b *testing.B) {
		type Form struct {
			Email string `validate:"email"`
		}
		form := Form{Email: "user@example.com"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("URL", func(b *testing.B) {
		type Form struct {
			URL string `validate:"url"`
		}
		form := Form{URL: "https://example.com/path?query=value"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("EmailAndURL", func(b *testing.B) {
		type Form struct {
			Email string `validate:"required,email"`
			URL   string `validate:"required,url"`
		}
		form := Form{
			Email: "user@example.com",
			URL:   "https://example.com",
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})
}

// BenchmarkValidate_MultiField 多字段结构体验证性能测试
func BenchmarkValidate_MultiField(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("5Fields", func(b *testing.B) {
		type Form struct {
			Name     string `json:"name" validate:"required,min=2"`
			Email    string `json:"email" validate:"required,email"`
			Age      int    `json:"age" validate:"required,min=1,max=150"`
			Phone    string `json:"phone" validate:"required"`
			Password string `json:"password" validate:"required,min=6"`
		}
		form := Form{
			Name:     "Zhang San",
			Email:    "test@example.com",
			Age:      25,
			Phone:    "13812345678",
			Password: "MyPass123!",
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("10Fields", func(b *testing.B) {
		type Form struct {
			Name     string `json:"name" validate:"required,min=2"`
			Email    string `json:"email" validate:"required,email"`
			Age      int    `json:"age" validate:"required,min=1,max=150"`
			Phone    string `json:"phone" validate:"required"`
			Password string `json:"password" validate:"required,min=6"`
			Address  string `json:"address" validate:"required,min=5"`
			City     string `json:"city" validate:"required"`
			Country  string `json:"country" validate:"required,min=2"`
			ZipCode  string `json:"zip_code" validate:"required,len=6"`
			Website  string `json:"website" validate:"url"`
		}
		form := Form{
			Name:     "Zhang San",
			Email:    "test@example.com",
			Age:      25,
			Phone:    "13812345678",
			Password: "MyPass123!",
			Address:  "No.1 Main Street",
			City:     "Beijing",
			Country:  "China",
			ZipCode:  "100000",
			Website:  "https://example.com",
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})
}

// BenchmarkValidate_CustomValidator 自定义验证器性能测试
func BenchmarkValidate_CustomValidator(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("Mobile", func(b *testing.B) {
		type Form struct {
			Phone string `validate:"mobile"`
		}
		form := Form{Phone: "13812345678"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("IDCard", func(b *testing.B) {
		type Form struct {
			IDCard string `validate:"idcard"`
		}
		form := Form{IDCard: "11010119800101123X"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("ChineseName", func(b *testing.B) {
		type Form struct {
			Name string `validate:"chinese_name"`
		}
		form := Form{Name: "张三"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})

	b.Run("StrongPassword", func(b *testing.B) {
		type Form struct {
			Password string `validate:"strong_password"`
		}
		form := Form{Password: "MyPass123!"}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Struct(form)
		}
	})
}

// BenchmarkValidate_NestedStruct 嵌套结构体验证性能测试
func BenchmarkValidate_NestedStruct(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type Address struct {
		Street  string `json:"street" validate:"required,min=5"`
		City    string `json:"city" validate:"required"`
		ZipCode string `json:"zip_code" validate:"required,len=6"`
	}

	type Company struct {
		Name    string  `json:"name" validate:"required,min=2"`
		Address Address `json:"address"`
	}

	type Employee struct {
		Name    string  `json:"name" validate:"required,min=2"`
		Email   string  `json:"email" validate:"required,email"`
		Age     int     `json:"age" validate:"required,min=18,max=65"`
		Company Company `json:"company"`
		Home    Address `json:"home"`
	}

	employee := Employee{
		Name:  "Zhang San",
		Email: "zhangsan@example.com",
		Age:   30,
		Company: Company{
			Name: "Tech Corp",
			Address: Address{
				Street:  "No.100 Innovation Road",
				City:    "Beijing",
				ZipCode: "100000",
			},
		},
		Home: Address{
			Street:  "No.1 Residential Area",
			City:    "Beijing",
			ZipCode: "100001",
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(employee)
	}
}

// BenchmarkValidate_Batch 批量验证性能测试
func BenchmarkValidate_Batch(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name  string `json:"name" validate:"required,min=2"`
		Email string `json:"email" validate:"required,email"`
		Age   int    `json:"age" validate:"required,min=1"`
	}

	// 预构建 100 个对象
	users := make([]User, 100)
	for i := range users {
		users[i] = User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + i%50,
		}
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := range users {
			_ = v.Struct(users[j])
		}
	}
}

// BenchmarkValidate_ChineseMessage 中文错误消息性能测试
func BenchmarkValidate_ChineseMessage(b *testing.B) {
	v, err := New(WithLocale("zh"))
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name  string `json:"name" validate:"required,min=2"`
		Email string `json:"email" validate:"required,email"`
		Phone string `json:"phone" validate:"mobile"`
	}

	// 使用无效数据触发中文错误消息
	user := User{
		Name:  "",
		Email: "invalid",
		Phone: "123",
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(user)
	}
}

// BenchmarkValidate_Var 单变量验证性能测试
func BenchmarkValidate_Var(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	b.Run("Required", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Var("hello", "required")
		}
	})

	b.Run("Email", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Var("user@example.com", "email")
		}
	})

	b.Run("MinMax", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = v.Var(25, "min=1,max=100")
		}
	})
}

// BenchmarkValidate_Parallel 并发验证性能测试
func BenchmarkValidate_Parallel(b *testing.B) {
	v, err := New()
	if err != nil {
		b.Fatalf("Failed to create validator: %v", err)
	}

	type User struct {
		Name  string `json:"name" validate:"required,min=2"`
		Email string `json:"email" validate:"required,email"`
	}

	user := User{
		Name:  "Zhang San",
		Email: "test@example.com",
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = v.Struct(user)
		}
	})
}
