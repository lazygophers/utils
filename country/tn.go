//go:build country_africa || country_all || country_northern_africa || country_tn

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tunisia — Tunisian Republic.
var dataTunisia = &Country{
	alpha2:       "TN",
	alpha3:       "TUN",
	numeric:      788,
	callingCodes: []string{"+216"},
	timezones:    []string{"Africa/Tunis"},
	tlds:         []string{
		".tn",
		".تونس",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Tnd,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Northern Africa",
	flagEmoji:    "🇹🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTunisia) }

var Tunisia = dataTunisia
