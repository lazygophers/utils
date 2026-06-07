//go:build country_all || country_americas || country_caribbean || country_ht

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Haiti — Republic of Haiti.
var dataHaiti = &Country{
	alpha2:       "HT",
	alpha3:       "HTI",
	numeric:      332,
	callingCodes: []string{"+509"},
	timezones:    []string{"America/Port-au-Prince"},
	tlds:         []string{".ht"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.HTG,
	region:       RegionCaribbean,
	flagEmoji:    "🇭🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHaiti) }

var Haiti = dataHaiti
