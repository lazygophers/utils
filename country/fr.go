package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// France — French Republic.
var dataFrance = &Country{
	alpha2:       "FR",
	alpha3:       "FRA",
	numeric:      250,
	callingCodes: []string{"+33"},
	timezones:    []string{"Europe/Paris"},
	tlds:         []string{".fr"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Europe",
	flagEmoji:    "🇫🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFrance) }
