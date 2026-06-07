//go:build country_all || country_asia || country_la || country_south_eastern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Laos — Lao People's Democratic Republic.
var dataLaos = &Country{
	alpha2:       "LA",
	alpha3:       "LAO",
	numeric:      418,
	callingCodes: []string{"+856"},
	timezones:    []string{"Asia/Vientiane"},
	tlds:         []string{".la"},
	officialLanguage:  xlanguage.MustParse("lo"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("lo")},
	currency:     currency.LAK,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇱🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLaos) }

var Laos = dataLaos
