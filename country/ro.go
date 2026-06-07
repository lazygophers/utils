//go:build country_all || country_eastern_europe || country_europe || country_ro

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Romania — Romania.
var dataRomania = &Country{
	alpha2:       "RO",
	alpha3:       "ROU",
	numeric:      642,
	callingCodes: []string{"+40"},
	timezones:    []string{"Europe/Bucharest"},
	tlds:         []string{".ro"},
	languages:    []xlanguage.Tag{xlanguage.Romanian},
	currency:     currency.Ron,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇷🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataRomania) }

var Romania = dataRomania
