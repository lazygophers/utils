package currency_test

import (
	"testing"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
	"github.com/lazygophers/utils/language"
)

func BenchmarkGet(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = currency.Get("CNY")
	}
}

func BenchmarkGetByNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = currency.GetByNumeric(156)
	}
}

func BenchmarkCnyCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = currency.CNY.Code()
	}
}

func BenchmarkNameIn_Hit(b *testing.B) {
	b.ReportAllocs()
	cny := currency.CNY
	for i := 0; i < b.N; i++ {
		_ = cny.NameIn(xlanguage.English)
	}
}

func BenchmarkName_GoroutineLocal(b *testing.B) {
	b.ReportAllocs()
	language.Set(language.Make("en"))
	defer language.Del()
	cny := currency.CNY
	for i := 0; i < b.N; i++ {
		_ = cny.Name()
	}
}
