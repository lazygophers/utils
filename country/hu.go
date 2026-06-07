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
	languages:    []xlanguage.Tag{xlanguage.Hungarian},
	currency:     currency.Huf,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇭🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHungary) }
