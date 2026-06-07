//go:build country_africa || country_all || country_eastern_africa || country_mw

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Malawi — Republic of Malawi.
var dataMalawi = &Country{
	alpha2:       "MW",
	alpha3:       "MWI",
	numeric:      454,
	callingCodes: []string{"+265"},
	timezones:    []string{"Africa/Blantyre"},
	tlds:         []string{".mw"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("ny")},
	currency:     currency.Mwk,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇲🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMalawi) }

var Malawi = dataMalawi
