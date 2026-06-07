//go:build country_all || country_melanesia || country_nc || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NewCaledonia — New Caledonia — overseas collectivity of France.
var dataNewCaledonia = &Country{
	alpha2:       "NC",
	alpha3:       "NCL",
	numeric:      540,
	callingCodes: []string{"+687"},
	timezones:    []string{"Pacific/Noumea"},
	tlds:         []string{".nc"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XPF,
	region:       RegionMelanesia,
	flagEmoji:    "🇳🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNewCaledonia) }

var NewCaledonia = dataNewCaledonia
