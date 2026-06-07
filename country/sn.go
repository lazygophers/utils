//go:build country_africa || country_all || country_sn || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Senegal — Republic of Senegal.
var dataSenegal = &Country{
	alpha2:       "SN",
	alpha3:       "SEN",
	numeric:      686,
	callingCodes: []string{"+221"},
	timezones:    []string{"Africa/Dakar"},
	tlds:         []string{".sn"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Xof,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇸🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSenegal) }

var Senegal = dataSenegal
