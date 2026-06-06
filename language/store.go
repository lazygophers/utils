package language

import (
	"bytes"
	"runtime"
	"strconv"
	"sync"
)

// defaultLang is the global default language. English by default.
var defaultLang = Make("en")

// localLangs stores per-goroutine language overrides.
var localLangs sync.Map // map[uint64]*Tag

// SetDefault sets the global default language.
func SetDefault(tag *Tag) {
	defaultLang = tag
}

// Default returns the global default language.
func Default() *Tag {
	return defaultLang
}

// Set stores the language tag for the current goroutine.
func Set(tag *Tag) {
	localLangs.Store(goroutineId(), tag)
}

// Del removes the language override for the current goroutine.
// After Del, Get() will return the global default.
func Del() {
	localLangs.Delete(goroutineId())
}

// Get returns the effective language for the current goroutine.
// Priority: goroutine-local override > global default.
func Get() *Tag {
	if v, ok := localLangs.Load(goroutineId()); ok {
		return v.(*Tag)
	}
	return defaultLang
}

// goroutineId returns the current goroutine's ID by parsing runtime.Stack output.
func goroutineId() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// buf[:n] = "goroutine 123 [running]:\n..."
	s := bytes.TrimPrefix(buf[:n], []byte("goroutine "))
	// find space before " [running]"
	i := bytes.IndexByte(s, ' ')
	if i <= 0 {
		return 0
	}
	id, _ := strconv.ParseUint(string(s[:i]), 10, 64)
	return id
}
