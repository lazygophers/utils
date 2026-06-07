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
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Eur,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇷🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataReunion) }
