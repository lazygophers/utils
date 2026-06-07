//go:build country_all || country_cy || country_europe || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Cyprus — Republic of Cyprus.
var dataCyprus = &Country{
	alpha2:       "CY",
	alpha3:       "CYP",
	numeric:      196,
	callingCodes: []string{"+357"},
	timezones:    []string{"Asia/Nicosia"},
	tlds:         []string{".cy"},
	officialLanguage:  xlanguage.Greek,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Greek, xlanguage.Turkish},
	currency:     currency.EUR,
	region:       RegionWesternAsia,
	flagEmoji:    "🇨🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCyprus) }

var Cyprus = dataCyprus
