//go:build country_africa || country_all || country_dz || country_northern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Algeria — People's Democratic Republic of Algeria.
var dataAlgeria = &Country{
	alpha2:       "DZ",
	alpha3:       "DZA",
	numeric:      12,
	callingCodes: []string{"+213"},
	timezones:    []string{"Africa/Algiers"},
	tlds:         []string{
		".dz",
		".الجزائر",
	},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic, xlanguage.French, xlanguage.MustParse("ber")},
	currency:     currency.DZD,
	region:       RegionNorthernAfrica,
	flagEmoji:    "🇩🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAlgeria) }

var Algeria = dataAlgeria
