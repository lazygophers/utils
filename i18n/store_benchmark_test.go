package i18n

import (
	"strconv"
	"testing"
)

func benchData(n int) []struct {
	locale, key, text string
} {
	out := make([]struct{ locale, key, text string }, n)
	for i := 0; i < n; i++ {
		out[i] = struct{ locale, key, text string }{
			locale: "en",
			key:    "msg_" + strconv.Itoa(i),
			text:   "Translated message number " + strconv.Itoa(i),
		}
	}
	return out
}

func BenchmarkMapStoreGet(b *testing.B) {
	s := newMapStore()
	data := benchData(1000)
	for _, d := range data {
		s.Set(d.locale, d.key, d.text)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = s.Get("en", data[i%len(data)].key)
	}
}

func BenchmarkChunkStoreGet(b *testing.B) {
	s := newChunkStore(1024 * 1024)
	data := benchData(1000)
	for _, d := range data {
		s.Set(d.locale, d.key, d.text)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = s.Get("en", data[i%len(data)].key)
	}
}

func BenchmarkMapStoreSet(b *testing.B) {
	s := newMapStore()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Set("en", "k"+strconv.Itoa(i&1023), "v")
	}
}

func BenchmarkChunkStoreSet(b *testing.B) {
	s := newChunkStore(1024 * 1024)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Set("en", "k"+strconv.Itoa(i&1023), "v")
	}
}
