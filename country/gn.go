//go:build country_africa || country_all || country_gn || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guinea — Republic of Guinea.
var dataGuinea = &Country{
	alpha2:       "GN",
	alpha3:       "GIN",
	numeric:      324,
	callingCodes: []string{"+224"},
	timezones:    []string{"Africa/Conakry"},
	tlds:         []string{".gn"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.GNF,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇬🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuinea) }

var Guinea = dataGuinea
