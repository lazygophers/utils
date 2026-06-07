package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// MarshallIslands — Republic of the Marshall Islands.
var dataMarshallIslands = &Country{
	alpha2:       "MH",
	alpha3:       "MHL",
	numeric:      584,
	callingCodes: []string{"+692"},
	timezones:    []string{
		"Pacific/Majuro",
		"Pacific/Kwajalein",
	},
	tlds:         []string{".mh"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("mh")},
	currency:     currency.Usd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Micronesia",
	flagEmoji:    "🇲🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMarshallIslands) }
