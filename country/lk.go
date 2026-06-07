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
	languages:    []xlanguage.Tag{xlanguage.MustParse("si"), xlanguage.MustParse("ta")},
	currency:     currency.Lkr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇱🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSriLanka) }
