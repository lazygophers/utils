//go:build country_africa || country_all || country_ng || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Nigeria — Federal Republic of Nigeria.
var dataNigeria = &Country{
	alpha2:       "NG",
	alpha3:       "NGA",
	numeric:      566,
	callingCodes: []string{"+234"},
	timezones:    []string{"Africa/Lagos"},
	tlds:         []string{".ng"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("ha"), xlanguage.MustParse("yo"), xlanguage.MustParse("ig")},
	currency:     currency.NGN,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇳🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNigeria) }

var Nigeria = dataNigeria
