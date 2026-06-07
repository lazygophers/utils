//go:build country_all || country_australia_and_new_zealand || country_nz || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NewZealand — New Zealand.
var dataNewZealand = &Country{
	alpha2:       "NZ",
	alpha3:       "NZL",
	numeric:      554,
	callingCodes: []string{"+64"},
	timezones:    []string{
		"Pacific/Auckland",
		"Pacific/Chatham",
	},
	tlds:         []string{".nz"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("mi")},
	currency:     currency.NZD,
	region:       RegionAustraliaAndNewZealand,
	flagEmoji:    "🇳🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNewZealand) }

var NewZealand = dataNewZealand
