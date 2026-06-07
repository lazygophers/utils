package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// HeardAndMcDonaldIslands — Heard Island and McDonald Islands — Australian external territory.
var dataHeardAndMcDonaldIslands = &Country{
	alpha2:       "HM",
	alpha3:       "HMD",
	numeric:      334,
	callingCodes: []string{},
	timezones:    []string{"Indian/Kerguelen"},
	tlds:         []string{".hm"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Aud,
	continent:    "AN",
	region:       "Antarctic",
	subregion:    "Antarctic",
	flagEmoji:    "🇭🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHeardAndMcDonaldIslands) }
