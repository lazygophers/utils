package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// AlandIslands — Åland Islands — autonomous region of Finland.
var dataAlandIslands = &Country{
	alpha2:       "AX",
	alpha3:       "ALA",
	numeric:      248,
	callingCodes: []string{"+358-18"},
	timezones:    []string{"Europe/Mariehamn"},
	tlds:         []string{".ax"},
	languages:    []xlanguage.Tag{xlanguage.Swedish},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇦🇽",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAlandIslands) }
