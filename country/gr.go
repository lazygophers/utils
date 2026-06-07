//go:build country_all || country_europe || country_gr || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Greece — Hellenic Republic.
var dataGreece = &Country{
	alpha2:       "GR",
	alpha3:       "GRC",
	numeric:      300,
	callingCodes: []string{"+30"},
	timezones:    []string{"Europe/Athens"},
	tlds:         []string{".gr"},
	languages:    []xlanguage.Tag{xlanguage.Greek},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇬🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGreece) }

var Greece = dataGreece
