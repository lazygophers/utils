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
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.Afrikaans},
	currency:     currency.Nad,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Southern Africa",
	flagEmoji:    "🇳🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNamibia) }
