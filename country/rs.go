package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Serbia — Republic of Serbia.
var dataSerbia = &Country{
	alpha2:       "RS",
	alpha3:       "SRB",
	numeric:      688,
	callingCodes: []string{"+381"},
	timezones:    []string{"Europe/Belgrade"},
	tlds:         []string{
		".rs",
		".срб",
	},
	languages:    []xlanguage.Tag{xlanguage.Serbian},
	currency:     currency.Rsd,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇷🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSerbia) }
