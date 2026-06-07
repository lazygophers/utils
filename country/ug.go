//go:build country_africa || country_all || country_eastern_africa || country_ug

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Uganda — Republic of Uganda.
var dataUganda = &Country{
	alpha2:       "UG",
	alpha3:       "UGA",
	numeric:      800,
	callingCodes: []string{"+256"},
	timezones:    []string{"Africa/Kampala"},
	tlds:         []string{".ug"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("sw")},
	currency:     currency.Ugx,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇺🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUganda) }

var Uganda = dataUganda
