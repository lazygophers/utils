package language

import (
	"sync"

	"github.com/petermattis/goid"
)

// defaultLang is the global default language. English by default.
var defaultLang = Make("en")

// localLangs stores per-goroutine language overrides.
var localLangs sync.Map // map[int64]*Tag

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
	localLangs.Store(goid.Get(), tag)
}

// Del removes the language override for the current goroutine.
// After Del, Get() will return the global default.
func Del() {
	localLangs.Delete(goid.Get())
}

// Get returns the effective language for the current goroutine.
// Priority: goroutine-local override > global default.
func Get() *Tag {
	v, ok := localLangs.Load(goid.Get())
	if ok {
		return v.(*Tag)
	}
	return defaultLang
}
