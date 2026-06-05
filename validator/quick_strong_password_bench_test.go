package validator

import (
	"testing"
)

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"email"`
	Age   int    `validate:"min=18,max=100"`
}

func TestValidateFieldOptimized(t *testing.T) {
	e := NewEngine()

	// 测试 required
	user := User{Name: "John", Email: "john@example.com", Age: 25}
	err := e.Struct(&user)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 测试 email 失败
	user2 := User{Name: "Jane", Email: "invalid-email", Age: 30}
	err = e.Struct(&user2)
	if err == nil {
		t.Error("Expected error for invalid email")
	}

	// 测试 min 失败
	user3 := User{Name: "Bob", Email: "bob@example.com", Age: 15}
	err = e.Struct(&user3)
	if err == nil {
		t.Error("Expected error for age < 18")
	}

	t.Log("All tests passed!")
}
