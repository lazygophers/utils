//go:build country_all || country_americas || country_caribbean || country_gp

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guadeloupe — Guadeloupe — overseas region of France.
var dataGuadeloupe = &Country{
	alpha2:       "GP",
	alpha3:       "GLP",
	numeric:      312,
	callingCodes: []string{"+590"},
	timezones:    []string{"America/Guadeloupe"},
	tlds:         []string{".gp"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionCaribbean,
	flagEmoji:    "🇬🇵",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuadeloupe) }

var Guadeloupe = dataGuadeloupe
