//go:build country_africa || country_all || country_eastern_africa || country_re

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Reunion — Réunion — overseas region of France.
var dataReunion = &Country{
	alpha2:       "RE",
	alpha3:       "REU",
	numeric:      638,
	callingCodes: []string{"+262"},
	timezones:    []string{"Indian/Reunion"},
	tlds:         []string{".re"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇷🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataReunion) }

var Reunion = dataReunion
