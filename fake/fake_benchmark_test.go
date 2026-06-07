package fake_test

import (
	"testing"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func BenchmarkName(b *testing.B) {
	f := fake.New(country.UnitedStates)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Name()
	}
}

func BenchmarkNameCN(b *testing.B) {
	f := fake.New(country.China)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Name()
	}
}

func BenchmarkUUIDv4(b *testing.B) {
	f := fake.New(country.UnitedStates)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.UUIDv4()
	}
}

func BenchmarkUUIDv7(b *testing.B) {
	f := fake.New(country.UnitedStates)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.UUIDv7()
	}
}

func BenchmarkPhoneCN(b *testing.B) {
	f := fake.New(country.China)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Phone()
	}
}

func BenchmarkPhoneUS(b *testing.B) {
	f := fake.New(country.UnitedStates)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Phone()
	}
}

func BenchmarkIdCardCN(b *testing.B) {
	f := fake.New(country.China)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.IdCard()
	}
}

func BenchmarkIdCardUS(b *testing.B) {
	f := fake.New(country.UnitedStates)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.IdCard()
	}
}

func BenchmarkEmail(b *testing.B) {
	f := fake.New(country.UnitedStates)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Email()
	}
}

func BenchmarkGlobalName(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fake.Name()
	}
}
