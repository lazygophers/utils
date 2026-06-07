package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// IvoryCoast — Republic of Côte d'Ivoire.
var dataIvoryCoast = &Country{
	alpha2:       "CI",
	alpha3:       "CIV",
	numeric:      384,
	callingCodes: []string{"+225"},
	timezones:    []string{"Africa/Abidjan"},
	tlds:         []string{".ci"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Xof,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇨🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIvoryCoast) }
