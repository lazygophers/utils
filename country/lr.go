//go:build country_africa || country_all || country_lr || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Liberia — Republic of Liberia.
var dataLiberia = &Country{
	alpha2:       "LR",
	alpha3:       "LBR",
	numeric:      430,
	callingCodes: []string{"+231"},
	timezones:    []string{"Africa/Monrovia"},
	tlds:         []string{".lr"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Lrd,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇱🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLiberia) }

var Liberia = dataLiberia
