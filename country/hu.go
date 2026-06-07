//go:build country_all || country_eastern_europe || country_europe || country_hu

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Hungary — Hungary.
var dataHungary = &Country{
	alpha2:       "HU",
	alpha3:       "HUN",
	numeric:      348,
	callingCodes: []string{"+36"},
	timezones:    []string{"Europe/Budapest"},
	tlds:         []string{".hu"},
	officialLanguage:  xlanguage.Hungarian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Hungarian},
	currency:     currency.HUF,
	region:       RegionEasternEurope,
	flagEmoji:    "🇭🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHungary) }

var Hungary = dataHungary
