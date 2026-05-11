package validator

import (
	"testing"
)

// TestValidateStrongPassword_Correctness 验证正确性
func TestValidateStrongPassword_Correctness(t *testing.T) {
	v, _ := New()

	type Form struct {
		Password string `validate:"strong_password"`
	}

	validPasswords := []string{
		"Abc123!@",           // 最小长度，包含所有类型
		"Password123!",       // 常见强密码
		"SecurePass#2024",    // 包含数字
		"MyP@ssw0rd",         // 复杂密码
		"Test@1234",          // 简单但符合
		"ADMIN@123",          // 全大写+数字+特殊
		"student#123",        // 全小写+数字+特殊
		"User2024$Pass",      // 混合
		"1A2b3C4d!",          // 交替字符
		"P@ssw0rd123456",     // 较长密码
	}

	invalidPasswords := []struct {
		password string
		reason   string
	}{
		{"", "空密码"},
		{"short1A", "太短（7字符）"},
		{"lowercas", "只1种类型（小写）"},
		{"UPPERCAS", "只1种类型（大写）"},
		{"12345678", "只1种类型（数字）"},
		{"!@#$%^&*", "只1种类型（特殊）"},
		{"lower12", "只2种类型（小写+数字）"},
		{"LOWER12", "只2种类型（大写+数字）"},
		{"lower!!", "只2种类型（小写+特殊）"},
		{"LOWER!!", "只2种类型（大写+特殊）"},
		{"1234!!", "只2种类型（数字+特殊）"},
		{"lowerUP", "只2种类型（小写+大写）"},
	}

	t.Run("ValidPasswords", func(t *testing.T) {
		for _, pwd := range validPasswords {
			t.Run(pwd, func(t *testing.T) {
				form := Form{Password: pwd}
				err := v.Struct(form)
				if err != nil {
					t.Errorf("有效密码被拒绝: %s, error: %v", pwd, err)
				}
			})
		}
	})

	t.Run("InvalidPasswords", func(t *testing.T) {
		for _, tc := range invalidPasswords {
			t.Run(tc.reason, func(t *testing.T) {
				form := Form{Password: tc.password}
				err := v.Struct(form)
				if err == nil {
					t.Errorf("无效密码被接受: %s (%s)", tc.password, tc.reason)
				}
			})
		}
	})
}

// BenchmarkStrongPassword_Validator 集成基准测试
func BenchmarkStrongPassword_Validator(b *testing.B) {
	v, _ := New()

	type Form struct {
		Password string `validate:"strong_password"`
	}

	validPasswords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}

	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validPasswords {
				form := Form{Password: pwd}
				_ = v.Struct(form)
			}
		}
	})
}
