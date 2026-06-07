//go:build country_all || country_asia || country_central_asia || country_tj

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
	officialLanguage:  xlanguage.MustParse("tg"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("tg"), xlanguage.Russian},
	currency:     currency.TJS,
	region:       RegionCentralAsia,
	flagEmoji:    "🇹🇯",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTajikistan) }

var Tajikistan = dataTajikistan
