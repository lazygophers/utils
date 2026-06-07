//go:build country_all || country_melanesia || country_oceania || country_sb

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
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.SBD,
	region:       RegionMelanesia,
	flagEmoji:    "🇸🇧",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSolomonIslands) }

var SolomonIslands = dataSolomonIslands
