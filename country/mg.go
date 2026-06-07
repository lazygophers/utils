//go:build country_africa || country_all || country_eastern_africa || country_mg

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Madagascar — Republic of Madagascar.
var dataMadagascar = &Country{
	alpha2:       "MG",
	alpha3:       "MDG",
	numeric:      450,
	callingCodes: []string{"+261"},
	timezones:    []string{"Indian/Antananarivo"},
	tlds:         []string{".mg"},
	officialLanguage:  xlanguage.MustParse("mg"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("mg"), xlanguage.French},
	currency:     currency.MGA,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇲🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMadagascar) }

var Madagascar = dataMadagascar
