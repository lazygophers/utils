package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Lesotho — Kingdom of Lesotho.
var dataLesotho = &Country{
	alpha2:       "LS",
	alpha3:       "LSO",
	numeric:      426,
	callingCodes: []string{"+266"},
	timezones:    []string{"Africa/Maseru"},
	tlds:         []string{".ls"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("st")},
	currency:     currency.Lsl,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Southern Africa",
	flagEmoji:    "🇱🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLesotho) }
