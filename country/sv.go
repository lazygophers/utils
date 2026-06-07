//go:build country_all || country_americas || country_central_america || country_sv

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// ElSalvador — Republic of El Salvador.
var dataElSalvador = &Country{
	alpha2:       "SV",
	alpha3:       "SLV",
	numeric:      222,
	callingCodes: []string{"+503"},
	timezones:    []string{"America/El_Salvador"},
	tlds:         []string{".sv"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.USD,
	region:       RegionCentralAmerica,
	flagEmoji:    "🇸🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataElSalvador) }

var ElSalvador = dataElSalvador
