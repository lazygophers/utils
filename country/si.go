//go:build country_all || country_europe || country_si || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Slovenia — Republic of Slovenia.
var dataSlovenia = &Country{
	alpha2:       "SI",
	alpha3:       "SVN",
	numeric:      705,
	callingCodes: []string{"+386"},
	timezones:    []string{"Europe/Ljubljana"},
	tlds:         []string{".si"},
	officialLanguage:  xlanguage.Slovenian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Slovenian},
	currency:     currency.EUR,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇸🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSlovenia) }

var Slovenia = dataSlovenia
