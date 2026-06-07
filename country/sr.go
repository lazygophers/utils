package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Suriname — Republic of Suriname.
var dataSuriname = &Country{
	alpha2:       "SR",
	alpha3:       "SUR",
	numeric:      740,
	callingCodes: []string{"+597"},
	timezones:    []string{"America/Paramaribo"},
	tlds:         []string{".sr"},
	languages:    []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.Srd,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇸🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSuriname) }
