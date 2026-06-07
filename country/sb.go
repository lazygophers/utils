package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SolomonIslands — Solomon Islands.
var dataSolomonIslands = &Country{
	alpha2:       "SB",
	alpha3:       "SLB",
	numeric:      90,
	callingCodes: []string{"+677"},
	timezones:    []string{"Pacific/Guadalcanal"},
	tlds:         []string{".sb"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Sbd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Melanesia",
	flagEmoji:    "🇸🇧",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSolomonIslands) }
