//go:build country_all || country_fm || country_micronesia || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Micronesia — Federated States of Micronesia.
var dataMicronesia = &Country{
	alpha2:       "FM",
	alpha3:       "FSM",
	numeric:      583,
	callingCodes: []string{"+691"},
	timezones:    []string{
		"Pacific/Chuuk",
		"Pacific/Pohnpei",
		"Pacific/Kosrae",
	},
	tlds:         []string{".fm"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Micronesia",
	flagEmoji:    "🇫🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMicronesia) }

var Micronesia = dataMicronesia
