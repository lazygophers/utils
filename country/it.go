package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Italy — Italian Republic.
var dataItaly = &Country{
	alpha2:       "IT",
	alpha3:       "ITA",
	numeric:      380,
	callingCodes: []string{"+39"},
	timezones:    []string{"Europe/Rome"},
	tlds:         []string{".it"},
	languages:    []xlanguage.Tag{xlanguage.Italian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇮🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataItaly) }
