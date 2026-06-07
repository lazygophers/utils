//go:build country_all || country_oceania || country_pf || country_polynesia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// FrenchPolynesia — French Polynesia — overseas collectivity of France.
var dataFrenchPolynesia = &Country{
	alpha2:       "PF",
	alpha3:       "PYF",
	numeric:      258,
	callingCodes: []string{"+689"},
	timezones:    []string{
		"Pacific/Tahiti",
		"Pacific/Marquesas",
		"Pacific/Gambier",
	},
	tlds:         []string{".pf"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XPF,
	region:       RegionPolynesia,
	flagEmoji:    "🇵🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFrenchPolynesia) }

var FrenchPolynesia = dataFrenchPolynesia
