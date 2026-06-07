//go:build country_all || country_eastern_europe || country_europe || country_ua

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ukraine — Ukraine.
var dataUkraine = &Country{
	alpha2:       "UA",
	alpha3:       "UKR",
	numeric:      804,
	callingCodes: []string{"+380"},
	timezones:    []string{"Europe/Kyiv"},
	tlds:         []string{
		".ua",
		".укр",
	},
	languages:    []xlanguage.Tag{xlanguage.Ukrainian},
	currency:     currency.Uah,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇺🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUkraine) }

var Ukraine = dataUkraine
