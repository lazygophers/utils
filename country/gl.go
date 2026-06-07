package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Greenland — Greenland — autonomous territory of Denmark.
var dataGreenland = &Country{
	alpha2:       "GL",
	alpha3:       "GRL",
	numeric:      304,
	callingCodes: []string{"+299"},
	timezones:    []string{
		"America/Nuuk",
		"America/Danmarkshavn",
		"America/Scoresbysund",
		"America/Thule",
	},
	tlds:         []string{".gl"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("kl"), xlanguage.Danish},
	currency:     currency.Dkk,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Northern America",
	flagEmoji:    "🇬🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGreenland) }
