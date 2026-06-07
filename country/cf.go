package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CentralAfricanRepublic — Central African Republic.
var dataCentralAfricanRepublic = &Country{
	alpha2:       "CF",
	alpha3:       "CAF",
	numeric:      140,
	callingCodes: []string{"+236"},
	timezones:    []string{"Africa/Bangui"},
	tlds:         []string{".cf"},
	languages:    []xlanguage.Tag{xlanguage.French, xlanguage.MustParse("sg")},
	currency:     currency.Xaf,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Middle Africa",
	flagEmoji:    "🇨🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCentralAfricanRepublic) }
