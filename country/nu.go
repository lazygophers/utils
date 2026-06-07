//go:build country_all || country_nu || country_oceania || country_polynesia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Niue — Niue — self-governing in free association with New Zealand.
var dataNiue = &Country{
	alpha2:       "NU",
	alpha3:       "NIU",
	numeric:      570,
	callingCodes: []string{"+683"},
	timezones:    []string{"Pacific/Niue"},
	tlds:         []string{".nu"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.NZD,
	region:       RegionPolynesia,
	flagEmoji:    "🇳🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNiue) }

var Niue = dataNiue
