//go:build country_africa || country_all || country_gw || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// GuineaBissau — Republic of Guinea-Bissau.
var dataGuineaBissau = &Country{
	alpha2:       "GW",
	alpha3:       "GNB",
	numeric:      624,
	callingCodes: []string{"+245"},
	timezones:    []string{"Africa/Bissau"},
	tlds:         []string{".gw"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Xof,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇬🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuineaBissau) }

var GuineaBissau = dataGuineaBissau
