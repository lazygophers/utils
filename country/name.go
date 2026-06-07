package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/language"
)

// currentTag returns the stdlib language.Tag for the current goroutine.
// Resolution order: goroutine-local override > global default.
func currentTag() xlanguage.Tag {
	t := language.Get()
	if t == nil {
		return xlanguage.English
	}
	return t.Tag()
}

// RegisterName registers a localized common name for the given language tag.
// Intended to be called from locale_<lang>.go init() functions.
func (c *Country) RegisterName(tag xlanguage.Tag, name string) {
	c.namesMu.Lock()
	if c.names == nil {
		c.names = make(map[xlanguage.Tag]string)
	}
	c.names[tag] = name
	c.namesMu.Unlock()
}

// RegisterOfficialName registers a localized official name for the given
// language tag.
func (c *Country) RegisterOfficialName(tag xlanguage.Tag, name string) {
	c.namesMu.Lock()
	if c.official == nil {
		c.official = make(map[xlanguage.Tag]string)
	}
	c.official[tag] = name
	c.namesMu.Unlock()
}

// RegisterCapital registers a localized capital city name for the given
// language tag.
func (c *Country) RegisterCapital(tag xlanguage.Tag, name string) {
	c.namesMu.Lock()
	if c.capital == nil {
		c.capital = make(map[xlanguage.Tag]string)
	}
	c.capital[tag] = name
	c.namesMu.Unlock()
}

// Capital returns the capital city in the current goroutine's language.
func (c *Country) Capital() string {
	return c.CapitalIn(currentTag())
}

// CapitalIn returns the capital city in the given language. Falls back to
// language base, then English. Returns "" if no capital is registered (e.g.
// uninhabited territories).
func (c *Country) CapitalIn(tag xlanguage.Tag) string {
	c.namesMu.RLock()
	defer c.namesMu.RUnlock()
	if v, ok := c.capital[tag]; ok {
		return v
	}
	base, _ := tag.Base()
	baseTag := xlanguage.Make(base.String())
	if v, ok := c.capital[baseTag]; ok {
		return v
	}
	if v, ok := c.capital[xlanguage.English]; ok {
		return v
	}
	return ""
}

// Name returns the country's common name in the current goroutine's language.
// Falls back to language base, then English, then the alpha-2 code.
func (c *Country) Name() string {
	return c.NameIn(currentTag())
}

// NameIn returns the country's common name in the given language, with the
// same fallback chain as [Country.Name].
func (c *Country) NameIn(tag xlanguage.Tag) string {
	return lookupName(c, tag, false)
}

// OfficialName returns the country's official name in the current goroutine's
// language.
func (c *Country) OfficialName() string {
	return c.OfficialNameIn(currentTag())
}

// OfficialNameIn returns the country's official name in the given language.
// Falls back to language base, then English official name, then [Country.NameIn].
func (c *Country) OfficialNameIn(tag xlanguage.Tag) string {
	return lookupName(c, tag, true)
}

func lookupName(c *Country, tag xlanguage.Tag, official bool) string {
	c.namesMu.RLock()
	defer c.namesMu.RUnlock()
	m := c.names
	if official {
		m = c.official
	}
	if v, ok := m[tag]; ok {
		return v
	}
	base, _ := tag.Base()
	baseTag := xlanguage.Make(base.String())
	if v, ok := m[baseTag]; ok {
		return v
	}
	if v, ok := m[xlanguage.English]; ok {
		return v
	}
	if official {
		if v, ok := c.names[xlanguage.English]; ok {
			return v
		}
	}
	return c.alpha2
}
