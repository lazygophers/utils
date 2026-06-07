package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CaymanIslands — Cayman Islands — British Overseas Territory.
var dataCaymanIslands = &Country{
	alpha2:       "KY",
	alpha3:       "CYM",
	numeric:      136,
	callingCodes: []string{"+1-345"},
	timezones:    []string{"America/Cayman"},
	tlds:         []string{".ky"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Kyd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇰🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCaymanIslands) }
