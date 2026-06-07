package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Egypt — Arab Republic of Egypt.
var dataEgypt = &Country{
	alpha2:       "EG",
	alpha3:       "EGY",
	numeric:      818,
	callingCodes: []string{"+20"},
	timezones:    []string{"Africa/Cairo"},
	tlds:         []string{
		".eg",
		".مصر",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Egp,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Northern Africa",
	flagEmoji:    "🇪🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEgypt) }
