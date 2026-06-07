package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Curacao — Country of Curaçao.
var dataCuracao = &Country{
	alpha2:       "CW",
	alpha3:       "CUW",
	numeric:      531,
	callingCodes: []string{"+599"},
	timezones:    []string{"America/Curacao"},
	tlds:         []string{".cw"},
	languages:    []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.Ang,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇨🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCuracao) }
