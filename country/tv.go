//go:build country_all || country_oceania || country_polynesia || country_tv

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tuvalu — Tuvalu.
var dataTuvalu = &Country{
	alpha2:       "TV",
	alpha3:       "TUV",
	numeric:      798,
	callingCodes: []string{"+688"},
	timezones:    []string{"Pacific/Funafuti"},
	tlds:         []string{".tv"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("tvl")},
	currency:     currency.AUD,
	region:       RegionPolynesia,
	flagEmoji:    "🇹🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTuvalu) }

var Tuvalu = dataTuvalu
