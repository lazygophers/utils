package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Monaco — Principality of Monaco.
var dataMonaco = &Country{
	alpha2:       "MC",
	alpha3:       "MCO",
	numeric:      492,
	callingCodes: []string{"+377"},
	timezones:    []string{"Europe/Monaco"},
	tlds:         []string{".mc"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Europe",
	flagEmoji:    "🇲🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMonaco) }
