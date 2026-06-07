//go:build country_africa || country_all || country_bw || country_southern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Botswana — Republic of Botswana.
var dataBotswana = &Country{
	alpha2:       "BW",
	alpha3:       "BWA",
	numeric:      72,
	callingCodes: []string{"+267"},
	timezones:    []string{"Africa/Gaborone"},
	tlds:         []string{".bw"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("tn")},
	currency:     currency.BWP,
	region:       RegionSouthernAfrica,
	flagEmoji:    "🇧🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBotswana) }

var Botswana = dataBotswana
