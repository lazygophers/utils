//go:build country_all || country_europe || country_me || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Montenegro — Montenegro.
var dataMontenegro = &Country{
	alpha2:       "ME",
	alpha3:       "MNE",
	numeric:      499,
	callingCodes: []string{"+382"},
	timezones:    []string{"Europe/Podgorica"},
	tlds:         []string{".me"},
	languages:    []xlanguage.Tag{xlanguage.Serbian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇲🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMontenegro) }

var Montenegro = dataMontenegro
