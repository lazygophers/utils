//go:build country_all || country_europe || country_southern_europe || country_va

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// VaticanCity — Vatican City State.
var dataVaticanCity = &Country{
	alpha2:       "VA",
	alpha3:       "VAT",
	numeric:      336,
	callingCodes: []string{
		"+379",
		"+39-06",
	},
	timezones:    []string{"Europe/Vatican"},
	tlds:         []string{".va"},
	officialLanguage:  xlanguage.Italian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Italian, xlanguage.MustParse("la")},
	currency:     currency.EUR,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇻🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataVaticanCity) }

var VaticanCity = dataVaticanCity
