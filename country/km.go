//go:build country_africa || country_all || country_eastern_africa || country_km

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Comoros — Union of the Comoros.
var dataComoros = &Country{
	alpha2:       "KM",
	alpha3:       "COM",
	numeric:      174,
	callingCodes: []string{"+269"},
	timezones:    []string{"Indian/Comoro"},
	tlds:         []string{".km"},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic, xlanguage.French},
	currency:     currency.KMF,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇰🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataComoros) }

var Comoros = dataComoros
