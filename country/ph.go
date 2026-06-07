package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Philippines — Republic of the Philippines.
var dataPhilippines = &Country{
	alpha2:       "PH",
	alpha3:       "PHL",
	numeric:      608,
	callingCodes: []string{"+63"},
	timezones:    []string{"Asia/Manila"},
	tlds:         []string{".ph"},
	languages:    []xlanguage.Tag{xlanguage.Filipino, xlanguage.English},
	currency:     currency.Php,
	continent:    "AS",
	region:       "Asia",
	subregion:    "South-eastern Asia",
	flagEmoji:    "🇵🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPhilippines) }
