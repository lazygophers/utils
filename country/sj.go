//go:build country_all || country_europe || country_northern_europe || country_sj

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SvalbardAndJanMayen — Svalbard and Jan Mayen — Norwegian territories.
var dataSvalbardAndJanMayen = &Country{
	alpha2:       "SJ",
	alpha3:       "SJM",
	numeric:      744,
	callingCodes: []string{"+47"},
	timezones:    []string{"Arctic/Longyearbyen"},
	tlds:         []string{".sj"},
	officialLanguage:  xlanguage.Norwegian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Norwegian},
	currency:     currency.NOK,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇸🇯",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSvalbardAndJanMayen) }

var SvalbardAndJanMayen = dataSvalbardAndJanMayen
