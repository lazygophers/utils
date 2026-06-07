package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Sudan — Republic of the Sudan.
var dataSudan = &Country{
	alpha2:       "SD",
	alpha3:       "SDN",
	numeric:      729,
	callingCodes: []string{"+249"},
	timezones:    []string{"Africa/Khartoum"},
	tlds:         []string{".sd"},
	languages:    []xlanguage.Tag{xlanguage.Arabic, xlanguage.English},
	currency:     currency.Sdg,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Northern Africa",
	flagEmoji:    "🇸🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSudan) }
