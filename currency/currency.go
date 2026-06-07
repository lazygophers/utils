// Package currency provides ISO 4217 currency definitions with multi-language
// names. Each currency is exposed both as a strongly-typed package-level
// constant (e.g. [Cny], [Usd]) and via lookup by alphabetic/numeric code
// ([Get], [GetByNumeric]).
//
// All public APIs that accept a language tag use the standard library type
// golang.org/x/text/language.Tag. Goroutine-local language resolution is
// provided through github.com/lazygophers/utils/language.
package currency

import (
	"strings"
	"sync"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/language"
)

// Currency is an immutable ISO 4217 currency definition.
//
// All fields are unexported; accessors are read-only. Localised names are
// mutated only during package init via [Currency.RegisterName]; readers take
// an RLock at runtime.
type Currency struct {
	code    string
	symbol  string
	numeric int

	namesMu sync.RWMutex
	names   map[xlanguage.Tag]string
}

// Code returns the ISO 4217 alphabetic code (e.g. "CNY").
func (c *Currency) Code() string { return c.code }

// Symbol returns the conventional currency symbol (e.g. "¥").
func (c *Currency) Symbol() string { return c.symbol }

// Numeric returns the ISO 4217 numeric code (e.g. 156).
func (c *Currency) Numeric() int { return c.numeric }

// New constructs a Currency and initialises its name map. Intended for use
// in per-currency data files (e.g. currency/cny.go).
func New(code, symbol string, numeric int) *Currency {
	c := &Currency{
		code:    code,
		symbol:  symbol,
		numeric: numeric,
		names:   make(map[xlanguage.Tag]string),
	}
	register(c)
	return c
}

// RegisterName registers a localized currency name for the given language tag.
// Intended to be called from <code>_<lang>.go init() functions.
func (c *Currency) RegisterName(tag xlanguage.Tag, name string) {
	c.namesMu.Lock()
	if c.names == nil {
		c.names = make(map[xlanguage.Tag]string)
	}
	c.names[tag] = name
	c.namesMu.Unlock()
}

// Name returns the currency name in the current goroutine's language.
// Falls back to language base, then English, then the ISO 4217 code.
func (c *Currency) Name() string {
	return c.NameIn(currentTag())
}

// NameIn returns the currency name in the given language with the same
// fallback chain as [Currency.Name].
func (c *Currency) NameIn(tag xlanguage.Tag) string {
	c.namesMu.RLock()
	defer c.namesMu.RUnlock()
	if v, ok := c.names[tag]; ok {
		return v
	}
	base, _ := tag.Base()
	baseTag := xlanguage.Make(base.String())
	if v, ok := c.names[baseTag]; ok {
		return v
	}
	if v, ok := c.names[xlanguage.English]; ok {
		return v
	}
	return c.code
}

// String returns the ISO 4217 alphabetic code, satisfying fmt.Stringer.
func (c *Currency) String() string { return c.code }

// currentTag returns the stdlib language.Tag for the current goroutine.
func currentTag() xlanguage.Tag {
	t := language.Get()
	if t == nil {
		return xlanguage.English
	}
	return t.Tag()
}

// Registry indexes.
var (
	byCode     = make(map[string]*Currency, 200)
	byNumeric  = make(map[int]*Currency, 200)
	all        = make([]*Currency, 0, 200)
	registerMu sync.Mutex
)

func register(c *Currency) {
	registerMu.Lock()
	defer registerMu.Unlock()
	byCode[c.code] = c
	byNumeric[c.numeric] = c
	all = append(all, c)
}

// Get looks up a currency by ISO 4217 alphabetic code, case-insensitive.
// Returns nil if no match.
func Get(code string) *Currency {
	if len(code) != 3 {
		return nil
	}
	return byCode[strings.ToUpper(code)]
}

// GetByNumeric looks up a currency by ISO 4217 numeric code. Returns nil if
// no match.
func GetByNumeric(n int) *Currency { return byNumeric[n] }

// List returns a copy of all registered currencies.
func List() []*Currency {
	out := make([]*Currency, len(all))
	copy(out, all)
	return out
}
