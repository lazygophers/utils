//go:build country_all || country_americas || country_south_america || country_uy

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Uruguay — Oriental Republic of Uruguay.
var dataUruguay = &Country{
	alpha2:       "UY",
	alpha3:       "URY",
	numeric:      858,
	callingCodes: []string{"+598"},
	timezones:    []string{"America/Montevideo"},
	tlds:         []string{".uy"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.UYU,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇺🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUruguay) }

var Uruguay = dataUruguay
