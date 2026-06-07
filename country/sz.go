//go:build country_africa || country_all || country_southern_africa || country_sz

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Eswatini — Kingdom of Eswatini.
var dataEswatini = &Country{
	alpha2:       "SZ",
	alpha3:       "SWZ",
	numeric:      748,
	callingCodes: []string{"+268"},
	timezones:    []string{"Africa/Mbabane"},
	tlds:         []string{".sz"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("ss")},
	currency:     currency.Szl,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Southern Africa",
	flagEmoji:    "🇸🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEswatini) }

var Eswatini = dataEswatini
