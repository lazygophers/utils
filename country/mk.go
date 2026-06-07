//go:build country_all || country_europe || country_mk || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NorthMacedonia — Republic of North Macedonia.
var dataNorthMacedonia = &Country{
	alpha2:       "MK",
	alpha3:       "MKD",
	numeric:      807,
	callingCodes: []string{"+389"},
	timezones:    []string{"Europe/Skopje"},
	tlds:         []string{".mk"},
	officialLanguage:  xlanguage.MustParse("mk"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("mk")},
	currency:     currency.MKD,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇲🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNorthMacedonia) }

var NorthMacedonia = dataNorthMacedonia
