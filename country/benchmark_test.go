package country_test

import (
	"testing"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/language"
)

func BenchmarkGet(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = country.Get("CN")
	}
}

func BenchmarkGetByAlpha3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = country.GetByAlpha3("CHN")
	}
}

func BenchmarkGetByNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = country.GetByNumeric(156)
	}
}

func BenchmarkChinaAlpha2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = country.China.Alpha2()
	}
}

func BenchmarkNameIn_Hit(b *testing.B) {
	b.ReportAllocs()
	cn := country.China
	for i := 0; i < b.N; i++ {
		_ = cn.NameIn(xlanguage.English)
	}
}

func BenchmarkName_GoroutineLocal(b *testing.B) {
	b.ReportAllocs()
	language.Set(language.Make("zh"))
	defer language.Del()
	cn := country.China
	for i := 0; i < b.N; i++ {
		_ = cn.Name()
	}
}
