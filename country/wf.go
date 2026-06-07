//go:build country_all || country_oceania || country_polynesia || country_wf

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// WallisAndFutuna — Wallis and Futuna — overseas collectivity of France.
var dataWallisAndFutuna = &Country{
	alpha2:       "WF",
	alpha3:       "WLF",
	numeric:      876,
	callingCodes: []string{"+681"},
	timezones:    []string{"Pacific/Wallis"},
	tlds:         []string{".wf"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Xpf,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇼🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataWallisAndFutuna) }

var WallisAndFutuna = dataWallisAndFutuna
