package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guam — Territory of Guam.
var dataGuam = &Country{
	alpha2:       "GU",
	alpha3:       "GUM",
	numeric:      316,
	callingCodes: []string{"+1-671"},
	timezones:    []string{"Pacific/Guam"},
	tlds:         []string{".gu"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Usd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Micronesia",
	flagEmoji:    "🇬🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuam) }
