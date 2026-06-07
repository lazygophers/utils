package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mauritius — Republic of Mauritius.
var dataMauritius = &Country{
	alpha2:       "MU",
	alpha3:       "MUS",
	numeric:      480,
	callingCodes: []string{"+230"},
	timezones:    []string{"Indian/Mauritius"},
	tlds:         []string{".mu"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.French},
	currency:     currency.Mur,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇲🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMauritius) }
