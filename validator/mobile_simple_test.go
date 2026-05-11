package validator

import (
	"reflect"
	"testing"
)

func TestSimpleMobile(t *testing.T) {
	type Test struct {
		mobile string
		valid  bool
	}

	tests := []Test{
		{"13812345678", true},
		{"12812345678", false},
		{"138123456", false},
		{"", false},
	}

	for _, tt := range tests {
		result := validateMobile02(&testFieldLevel{value: reflect.ValueOf(tt.mobile)})
		if result != tt.valid {
			t.Errorf("validateMobile02(%s) = %v, want %v", tt.mobile, result, tt.valid)
		}
	}
}
