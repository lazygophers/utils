package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Slovenia — Republic of Slovenia.
var dataSlovenia = &Country{
	alpha2:       "SI",
	alpha3:       "SVN",
	numeric:      705,
	callingCodes: []string{"+386"},
	timezones:    []string{"Europe/Ljubljana"},
	tlds:         []string{".si"},
	languages:    []xlanguage.Tag{xlanguage.Slovenian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇸🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSlovenia) }
