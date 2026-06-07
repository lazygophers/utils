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
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("tn")},
	currency:     currency.Bwp,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Southern Africa",
	flagEmoji:    "🇧🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBotswana) }
