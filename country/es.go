//go:build country_all || country_es || country_europe || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Spain — Kingdom of Spain.
var dataSpain = &Country{
	alpha2:       "ES",
	alpha3:       "ESP",
	numeric:      724,
	callingCodes: []string{"+34"},
	timezones:    []string{
		"Europe/Madrid",
		"Atlantic/Canary",
	},
	tlds:         []string{".es"},
	languages:    []xlanguage.Tag{xlanguage.Spanish, xlanguage.Catalan},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇪🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSpain) }

var Spain = dataSpain
