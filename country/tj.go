package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tajikistan — Republic of Tajikistan.
var dataTajikistan = &Country{
	alpha2:       "TJ",
	alpha3:       "TJK",
	numeric:      762,
	callingCodes: []string{"+992"},
	timezones:    []string{"Asia/Dushanbe"},
	tlds:         []string{".tj"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("tg"), xlanguage.Russian},
	currency:     currency.Tjs,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Central Asia",
	flagEmoji:    "🇹🇯",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTajikistan) }
