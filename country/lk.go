//go:build country_all || country_asia || country_lk || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SriLanka — Democratic Socialist Republic of Sri Lanka.
var dataSriLanka = &Country{
	alpha2:       "LK",
	alpha3:       "LKA",
	numeric:      144,
	callingCodes: []string{"+94"},
	timezones:    []string{"Asia/Colombo"},
	tlds:         []string{
		".lk",
		".ලංකා",
	},
	officialLanguage:  xlanguage.MustParse("si"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("si"), xlanguage.MustParse("ta")},
	currency:     currency.LKR,
	region:       RegionSouthernAsia,
	flagEmoji:    "🇱🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSriLanka) }

var SriLanka = dataSriLanka
