//go:build country_africa || country_all || country_ao || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Angola — Republic of Angola.
var dataAngola = &Country{
	alpha2:       "AO",
	alpha3:       "AGO",
	numeric:      24,
	callingCodes: []string{"+244"},
	timezones:    []string{"Africa/Luanda"},
	tlds:         []string{".ao"},
	officialLanguage:  xlanguage.Portuguese,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.AOA,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇦🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAngola) }

var Angola = dataAngola
