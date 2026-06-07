//go:build country_africa || country_all || country_gh || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ghana — Republic of Ghana.
var dataGhana = &Country{
	alpha2:       "GH",
	alpha3:       "GHA",
	numeric:      288,
	callingCodes: []string{"+233"},
	timezones:    []string{"Africa/Accra"},
	tlds:         []string{".gh"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("ak"), xlanguage.MustParse("ee"), xlanguage.MustParse("gaa")},
	currency:     currency.GHS,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇬🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGhana) }

var Ghana = dataGhana
