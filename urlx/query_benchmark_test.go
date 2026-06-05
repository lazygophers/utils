package urlx

import (
	"net/url"
	"testing"
)

func BenchmarkSortQuery_Small(b *testing.B) {
	query := url.Values{
		"name": []string{"john"},
		"age":  []string{"30"},
		"city": []string{"beijing"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}

func BenchmarkSortQuery_Medium(b *testing.B) {
	query := url.Values{
		"param1":  []string{"value1"},
		"param2":  []string{"value2"},
		"param3":  []string{"value3"},
		"param4":  []string{"value4"},
		"param5":  []string{"value5"},
		"param6":  []string{"value6"},
		"param7":  []string{"value7"},
		"param8":  []string{"value8"},
		"param9":  []string{"value9"},
		"param10": []string{"value10"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}

func BenchmarkSortQuery_Large(b *testing.B) {
	query := make(url.Values)
	for i := 0; i < 100; i++ {
		query.Add("param"+string(rune(i)), "value"+string(rune(i)))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}

func BenchmarkSortQuery_Empty(b *testing.B) {
	query := url.Values{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortQuery(query)
	}
}
