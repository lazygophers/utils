//go:build country_africa || country_all || country_ga || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Gabon — Gabonese Republic.
var dataGabon = &Country{
	alpha2:       "GA",
	alpha3:       "GAB",
	numeric:      266,
	callingCodes: []string{"+241"},
	timezones:    []string{"Africa/Libreville"},
	tlds:         []string{".ga"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XAF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇬🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGabon) }

var Gabon = dataGabon
