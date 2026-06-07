//go:build country_all || country_americas || country_caribbean || country_ky

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
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.KYD,
	region:       RegionCaribbean,
	flagEmoji:    "🇰🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCaymanIslands) }

var CaymanIslands = dataCaymanIslands
