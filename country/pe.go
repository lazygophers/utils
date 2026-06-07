//go:build country_all || country_americas || country_pe || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Peru — Republic of Peru.
var dataPeru = &Country{
	alpha2:       "PE",
	alpha3:       "PER",
	numeric:      604,
	callingCodes: []string{"+51"},
	timezones:    []string{"America/Lima"},
	tlds:         []string{".pe"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish, xlanguage.MustParse("qu"), xlanguage.MustParse("ay")},
	currency:     currency.PEN,
	region:       RegionSouthAmerica,
	flagEmoji:    "🇵🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPeru) }

var Peru = dataPeru
