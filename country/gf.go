//go:build country_all || country_americas || country_gf || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// FrenchGuiana — French Guiana — overseas region of France.
var dataFrenchGuiana = &Country{
	alpha2:       "GF",
	alpha3:       "GUF",
	numeric:      254,
	callingCodes: []string{"+594"},
	timezones:    []string{"America/Cayenne"},
	tlds:         []string{".gf"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇬🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFrenchGuiana) }

var FrenchGuiana = dataFrenchGuiana
