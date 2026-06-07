//go:build country_all || country_europe || country_gi || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Gibraltar — British Overseas Territory of Gibraltar.
var dataGibraltar = &Country{
	alpha2:       "GI",
	alpha3:       "GIB",
	numeric:      292,
	callingCodes: []string{"+350"},
	timezones:    []string{"Europe/Gibraltar"},
	tlds:         []string{".gi"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Gip,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇬🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGibraltar) }

var Gibraltar = dataGibraltar
