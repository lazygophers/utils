//go:build country_all || country_americas || country_central_america || country_ni

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Nicaragua — Republic of Nicaragua.
var dataNicaragua = &Country{
	alpha2:       "NI",
	alpha3:       "NIC",
	numeric:      558,
	callingCodes: []string{"+505"},
	timezones:    []string{"America/Managua"},
	tlds:         []string{".ni"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.NIO,
	region:       RegionCentralAmerica,
	flagEmoji:    "🇳🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNicaragua) }

var Nicaragua = dataNicaragua
