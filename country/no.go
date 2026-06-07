//go:build country_all || country_europe || country_no || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Norway — Kingdom of Norway.
var dataNorway = &Country{
	alpha2:       "NO",
	alpha3:       "NOR",
	numeric:      578,
	callingCodes: []string{"+47"},
	timezones:    []string{"Europe/Oslo"},
	tlds:         []string{".no"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("nb"), xlanguage.MustParse("nn")},
	currency:     currency.Nok,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇳🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNorway) }

var Norway = dataNorway
