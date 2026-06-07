//go:build country_all || country_at || country_europe || country_western_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Austria — Republic of Austria.
var dataAustria = &Country{
	alpha2:       "AT",
	alpha3:       "AUT",
	numeric:      40,
	callingCodes: []string{"+43"},
	timezones:    []string{"Europe/Vienna"},
	tlds:         []string{".at"},
	officialLanguage:  xlanguage.German,
	spokenLanguages:   []xlanguage.Tag{xlanguage.German},
	currency:     currency.EUR,
	region:       RegionWesternEurope,
	flagEmoji:    "🇦🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAustria) }

var Austria = dataAustria
