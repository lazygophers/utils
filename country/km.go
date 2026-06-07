package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Comoros — Union of the Comoros.
var dataComoros = &Country{
	alpha2:       "KM",
	alpha3:       "COM",
	numeric:      174,
	callingCodes: []string{"+269"},
	timezones:    []string{"Indian/Comoro"},
	tlds:         []string{".km"},
	languages:    []xlanguage.Tag{xlanguage.Arabic, xlanguage.French},
	currency:     currency.Kmf,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇰🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataComoros) }
