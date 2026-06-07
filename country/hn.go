//go:build country_all || country_americas || country_central_america || country_hn

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Honduras — Republic of Honduras.
var dataHonduras = &Country{
	alpha2:       "HN",
	alpha3:       "HND",
	numeric:      340,
	callingCodes: []string{"+504"},
	timezones:    []string{"America/Tegucigalpa"},
	tlds:         []string{".hn"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.HNL,
	region:       RegionCentralAmerica,
	flagEmoji:    "🇭🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHonduras) }

var Honduras = dataHonduras
