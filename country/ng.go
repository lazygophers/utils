package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Nigeria — Federal Republic of Nigeria.
var dataNigeria = &Country{
	alpha2:       "NG",
	alpha3:       "NGA",
	numeric:      566,
	callingCodes: []string{"+234"},
	timezones:    []string{"Africa/Lagos"},
	tlds:         []string{".ng"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Ngn,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇳🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNigeria) }
