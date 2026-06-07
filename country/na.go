//go:build country_africa || country_all || country_na || country_southern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Namibia — Republic of Namibia.
var dataNamibia = &Country{
	alpha2:       "NA",
	alpha3:       "NAM",
	numeric:      516,
	callingCodes: []string{"+264"},
	timezones:    []string{"Africa/Windhoek"},
	tlds:         []string{".na"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.Afrikaans},
	currency:     currency.NAD,
	region:       RegionSouthernAfrica,
	flagEmoji:    "🇳🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNamibia) }

var Namibia = dataNamibia
