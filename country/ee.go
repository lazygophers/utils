package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Estonia — Republic of Estonia.
var dataEstonia = &Country{
	alpha2:       "EE",
	alpha3:       "EST",
	numeric:      233,
	callingCodes: []string{"+372"},
	timezones:    []string{"Europe/Tallinn"},
	tlds:         []string{".ee"},
	languages:    []xlanguage.Tag{xlanguage.Estonian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇪🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEstonia) }
