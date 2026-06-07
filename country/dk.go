//go:build country_all || country_dk || country_europe || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Denmark — Kingdom of Denmark.
var dataDenmark = &Country{
	alpha2:       "DK",
	alpha3:       "DNK",
	numeric:      208,
	callingCodes: []string{"+45"},
	timezones:    []string{"Europe/Copenhagen"},
	tlds:         []string{".dk"},
	officialLanguage:  xlanguage.Danish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Danish},
	currency:     currency.DKK,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇩🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataDenmark) }

var Denmark = dataDenmark
