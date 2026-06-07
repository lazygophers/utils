package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CookIslands — Cook Islands — self-governing in free association with New Zealand.
var dataCookIslands = &Country{
	alpha2:       "CK",
	alpha3:       "COK",
	numeric:      184,
	callingCodes: []string{"+682"},
	timezones:    []string{"Pacific/Rarotonga"},
	tlds:         []string{".ck"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Nzd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇨🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCookIslands) }
