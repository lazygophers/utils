package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CocosKeelingIslands — Cocos (Keeling) Islands — Australian external territory.
var dataCocosKeelingIslands = &Country{
	alpha2:       "CC",
	alpha3:       "CCK",
	numeric:      166,
	callingCodes: []string{"+61"},
	timezones:    []string{"Indian/Cocos"},
	tlds:         []string{".cc"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Aud,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Australia and New Zealand",
	flagEmoji:    "🇨🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCocosKeelingIslands) }
