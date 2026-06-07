//go:build country_all || country_asia || country_south_eastern_asia || country_tl

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// TimorLeste — Democratic Republic of Timor-Leste.
var dataTimorLeste = &Country{
	alpha2:       "TL",
	alpha3:       "TLS",
	numeric:      626,
	callingCodes: []string{"+670"},
	timezones:    []string{"Asia/Dili"},
	tlds:         []string{".tl"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Usd,
	continent:    "AS",
	region:       "Asia",
	subregion:    "South-eastern Asia",
	flagEmoji:    "🇹🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTimorLeste) }

var TimorLeste = dataTimorLeste
