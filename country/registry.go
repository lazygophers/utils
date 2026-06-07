package country

import (
	"sort"
	"strings"
	"sync"

	xlanguage "golang.org/x/text/language"
)

// Registry indexes; populated by data_*.go init() via [register]. After
// package init these maps are read-only and require no synchronisation.
var (
	byAlpha2   = make(map[string]*Country, 256)
	byAlpha3   = make(map[string]*Country, 256)
	byNumeric  = make(map[int]*Country, 256)
	all        = make([]*Country, 0, 256)
	registerMu sync.Mutex
)

// register inserts c into every index. It is safe to call concurrently from
// init() functions across multiple files; production access happens after
// init completes and is lock-free.
func register(c *Country) {
	registerMu.Lock()
	defer registerMu.Unlock()
	byAlpha2[c.alpha2] = c
	byAlpha3[c.alpha3] = c
	byNumeric[c.numeric] = c
	all = append(all, c)
}

// Get looks up a country by ISO 3166-1 alpha-2 code, case-insensitive.
// Returns nil if no match.
func Get(code string) *Country {
	if len(code) != 2 {
		return nil
	}
	return byAlpha2[strings.ToUpper(code)]
}

// GetByAlpha3 looks up a country by ISO 3166-1 alpha-3 code, case-insensitive.
// Returns nil if no match.
func GetByAlpha3(code string) *Country {
	if len(code) != 3 {
		return nil
	}
	return byAlpha3[strings.ToUpper(code)]
}

// GetByNumeric looks up a country by ISO 3166-1 numeric code.
// Returns nil if no match.
func GetByNumeric(n int) *Country {
	return byNumeric[n]
}

// GetByName looks up a country whose registered common or official name (in
// any language) matches the input, case-insensitive. Returns nil if no match.
//
// This is a linear scan over the name maps; intended for low-frequency lookups
// such as form input parsing, not hot paths.
func GetByName(name string) *Country {
	if name == "" {
		return nil
	}
	target := strings.ToLower(strings.TrimSpace(name))
	for _, c := range all {
		c.namesMu.RLock()
		hit := nameMatch(c.names, target) || nameMatch(c.official, target)
		c.namesMu.RUnlock()
		if hit {
			return c
		}
	}
	return nil
}

func nameMatch(m map[xlanguage.Tag]string, target string) bool {
	for _, v := range m {
		if strings.ToLower(v) == target {
			return true
		}
	}
	return false
}

// List returns all registered countries sorted by alpha-2 code.
// The returned slice is a copy; callers may mutate it freely.
func List() []*Country {
	out := make([]*Country, len(all))
	copy(out, all)
	sort.Slice(out, func(i, j int) bool { return out[i].alpha2 < out[j].alpha2 })
	return out
}
