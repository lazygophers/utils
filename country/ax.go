//go:build country_all || country_ax || country_europe || country_northern_europe

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
	officialLanguage:  xlanguage.Swedish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Swedish},
	currency:     currency.EUR,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇦🇽",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAlandIslands) }

var AlandIslands = dataAlandIslands
