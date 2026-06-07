//go:build country_all || country_asia || country_central_asia || country_uz

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Uzbekistan — Republic of Uzbekistan.
var dataUzbekistan = &Country{
	alpha2:       "UZ",
	alpha3:       "UZB",
	numeric:      860,
	callingCodes: []string{"+998"},
	timezones:    []string{
		"Asia/Tashkent",
		"Asia/Samarkand",
	},
	tlds:         []string{".uz"},
	officialLanguage:  xlanguage.MustParse("uz"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("uz"), xlanguage.Russian},
	currency:     currency.UZS,
	region:       RegionCentralAsia,
	flagEmoji:    "🇺🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUzbekistan) }

var Uzbekistan = dataUzbekistan
