package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Denmark — Kingdom of Denmark.
var dataDenmark = &Country{
	alpha2:       "DK",
	alpha3:       "DNK",
	numeric:      208,
	callingCodes: []string{"+45"},
	timezones:    []string{"Europe/Copenhagen"},
	tlds:         []string{".dk"},
	languages:    []xlanguage.Tag{xlanguage.Danish},
	currency:     currency.Dkk,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇩🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataDenmark) }
