//go:build country_all || country_cz || country_eastern_europe || country_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Czechia — Czech Republic.
var dataCzechia = &Country{
	alpha2:       "CZ",
	alpha3:       "CZE",
	numeric:      203,
	callingCodes: []string{"+420"},
	timezones:    []string{"Europe/Prague"},
	tlds:         []string{".cz"},
	languages:    []xlanguage.Tag{xlanguage.Czech},
	currency:     currency.Czk,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇨🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCzechia) }

var Czechia = dataCzechia
