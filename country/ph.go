//go:build country_all || country_asia || country_ph || country_south_eastern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Philippines — Republic of the Philippines.
var dataPhilippines = &Country{
	alpha2:       "PH",
	alpha3:       "PHL",
	numeric:      608,
	callingCodes: []string{"+63"},
	timezones:    []string{"Asia/Manila"},
	tlds:         []string{".ph"},
	officialLanguage:  xlanguage.Filipino,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Filipino, xlanguage.English, xlanguage.MustParse("ceb"), xlanguage.MustParse("ilo"), xlanguage.MustParse("hil")},
	currency:     currency.PHP,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇵🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPhilippines) }

var Philippines = dataPhilippines
