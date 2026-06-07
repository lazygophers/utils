//go:build country_all || country_asia || country_central_asia || country_uz

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Uzbekistan — Republic of Uzbekistan.
var dataUzbekistan = &Country{
	alpha2:       "UZ",
	alpha3:       "UZB",
	numeric:      860,
	callingCodes: []string{"+998"},
	timezones:    []string{
		"Asia/Tashkent",
		"Asia/Samarkand",
	},
	tlds:         []string{".uz"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("uz"), xlanguage.Russian},
	currency:     currency.Uzs,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Central Asia",
	flagEmoji:    "🇺🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUzbekistan) }

var Uzbekistan = dataUzbekistan
