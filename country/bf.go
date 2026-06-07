package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BurkinaFaso — Burkina Faso.
var dataBurkinaFaso = &Country{
	alpha2:       "BF",
	alpha3:       "BFA",
	numeric:      854,
	callingCodes: []string{"+226"},
	timezones:    []string{"Africa/Ouagadougou"},
	tlds:         []string{".bf"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Xof,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇧🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBurkinaFaso) }
