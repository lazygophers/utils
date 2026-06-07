//go:build country_africa || country_all || country_southern_africa || country_za

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SouthAfrica — Republic of South Africa.
var dataSouthAfrica = &Country{
	alpha2:       "ZA",
	alpha3:       "ZAF",
	numeric:      710,
	callingCodes: []string{"+27"},
	timezones:    []string{"Africa/Johannesburg"},
	tlds:         []string{".za"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.Afrikaans, xlanguage.MustParse("zu"), xlanguage.MustParse("xh"), xlanguage.MustParse("st"), xlanguage.MustParse("tn"), xlanguage.MustParse("ts"), xlanguage.MustParse("ss"), xlanguage.MustParse("ve"), xlanguage.MustParse("nr")},
	currency:     currency.Zar,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Southern Africa",
	flagEmoji:    "🇿🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSouthAfrica) }

var SouthAfrica = dataSouthAfrica
