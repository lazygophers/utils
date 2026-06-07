//go:build country_all || country_fj || country_melanesia || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Fiji — Republic of Fiji.
var dataFiji = &Country{
	alpha2:       "FJ",
	alpha3:       "FJI",
	numeric:      242,
	callingCodes: []string{"+679"},
	timezones:    []string{"Pacific/Fiji"},
	tlds:         []string{".fj"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("fj"), xlanguage.Hindi},
	currency:     currency.FJD,
	region:       RegionMelanesia,
	flagEmoji:    "🇫🇯",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFiji) }

var Fiji = dataFiji
