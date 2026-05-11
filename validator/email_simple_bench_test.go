package validator

import (
	"reflect"
	"strings"
	"testing"
)

var testEmails = []string{
	"test@example.com",
	"user.name@example.com",
	"user+tag@example.com",
	"invalid",
	"@example.com",
	"user@",
	"a@b.c",
	"very.long.email.address@very.long.domain.name.com",
	"",  // 空字符串
	"admin@mail.net",
	"hello@world.io",
	"info@company.co",
	"user123@test-domain.com",
	"first.last@sub.domain.org",
}

// 测试原始正则表达式方案（保留用于对比）
func BenchmarkEmail_OriginalRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, email := range testEmails {
			emailRegex.MatchString(email)
		}
	}
}

// 测试新的优化实现
func BenchmarkEmail_Optimized(b *testing.B) {
	validator := Email()
	fl := &fieldLevel{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, email := range testEmails {
			fl.field = reflect.ValueOf(email)
			validator(fl)
		}
	}
}

// 简单的内联版本用于基准测试
func BenchmarkEmail_InlineOptimized(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, email := range testEmails {
			if email == "" {
				continue
			}
			n := len(email)
			if n < 6 || n > 254 {
				continue
			}

			atIndex := strings.IndexByte(email, '@')
			if atIndex == -1 || atIndex == 0 || atIndex == n-1 {
				continue
			}

			localPart := email[:atIndex]
			domainPart := email[atIndex+1:]

			if len(localPart) == 0 || len(localPart) > 64 {
				continue
			}

			if len(domainPart) == 0 {
				continue
			}

			lastDot := strings.LastIndexByte(domainPart, '.')
			if lastDot == -1 || lastDot == 0 {
				continue
			}

			if len(domainPart)-lastDot-1 < 2 {
				continue
			}
		}
	}
}
