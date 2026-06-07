//go:build country_all || country_eastern_europe || country_europe || country_sk

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
	officialLanguage:  xlanguage.Slovak,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Slovak},
	currency:     currency.EUR,
	region:       RegionEasternEurope,
	flagEmoji:    "🇸🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSlovakia) }

var Slovakia = dataSlovakia
