//go:build country_all || country_melanesia || country_oceania || country_vu

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Vanuatu — Republic of Vanuatu.
var dataVanuatu = &Country{
	alpha2:       "VU",
	alpha3:       "VUT",
	numeric:      548,
	callingCodes: []string{"+678"},
	timezones:    []string{"Pacific/Efate"},
	tlds:         []string{".vu"},
	officialLanguage:  xlanguage.MustParse("bi"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("bi"), xlanguage.English, xlanguage.French},
	currency:     currency.VUV,
	region:       RegionMelanesia,
	flagEmoji:    "🇻🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataVanuatu) }

var Vanuatu = dataVanuatu
