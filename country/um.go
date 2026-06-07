package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// UsMinorOutlyingIslands — United States Minor Outlying Islands.
var dataUsMinorOutlyingIslands = &Country{
	alpha2:       "UM",
	alpha3:       "UMI",
	numeric:      581,
	callingCodes: []string{"+1"},
	timezones:    []string{
		"Pacific/Midway",
		"Pacific/Wake",
	},
	tlds:         []string{".us"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Micronesia",
	flagEmoji:    "🇺🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUsMinorOutlyingIslands) }
