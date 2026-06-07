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
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionWesternEurope,
	flagEmoji:    "🇫🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFrance) }

var France = dataFrance
