//go:build country_africa || country_all || country_ne || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Niger — Republic of the Niger.
var dataNiger = &Country{
	alpha2:       "NE",
	alpha3:       "NER",
	numeric:      562,
	callingCodes: []string{"+227"},
	timezones:    []string{"Africa/Niamey"},
	tlds:         []string{".ne"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XOF,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇳🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNiger) }

var Niger = dataNiger
