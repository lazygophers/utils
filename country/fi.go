package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Finland — Republic of Finland.
var dataFinland = &Country{
	alpha2:       "FI",
	alpha3:       "FIN",
	numeric:      246,
	callingCodes: []string{"+358"},
	timezones:    []string{"Europe/Helsinki"},
	tlds:         []string{".fi"},
	languages:    []xlanguage.Tag{xlanguage.Finnish, xlanguage.Swedish},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇫🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFinland) }
