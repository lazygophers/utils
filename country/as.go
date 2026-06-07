//go:build country_all || country_as || country_oceania || country_polynesia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// AmericanSamoa — Territory of American Samoa.
var dataAmericanSamoa = &Country{
	alpha2:       "AS",
	alpha3:       "ASM",
	numeric:      16,
	callingCodes: []string{"+1-684"},
	timezones:    []string{"Pacific/Pago_Pago"},
	tlds:         []string{".as"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("sm")},
	currency:     currency.Usd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇦🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAmericanSamoa) }

var AmericanSamoa = dataAmericanSamoa
