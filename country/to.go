//go:build country_all || country_oceania || country_polynesia || country_to

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tonga — Kingdom of Tonga.
var dataTonga = &Country{
	alpha2:       "TO",
	alpha3:       "TON",
	numeric:      776,
	callingCodes: []string{"+676"},
	timezones:    []string{"Pacific/Tongatapu"},
	tlds:         []string{".to"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("to")},
	currency:     currency.Top,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇹🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTonga) }

var Tonga = dataTonga
