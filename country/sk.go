package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Slovakia — Slovak Republic.
var dataSlovakia = &Country{
	alpha2:       "SK",
	alpha3:       "SVK",
	numeric:      703,
	callingCodes: []string{"+421"},
	timezones:    []string{"Europe/Bratislava"},
	tlds:         []string{".sk"},
	languages:    []xlanguage.Tag{xlanguage.Slovak},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇸🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSlovakia) }
