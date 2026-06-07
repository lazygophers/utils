package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// India — Republic of India.
var dataIndia = &Country{
	alpha2:       "IN",
	alpha3:       "IND",
	numeric:      356,
	callingCodes: []string{"+91"},
	timezones:    []string{"Asia/Kolkata"},
	tlds:         []string{
		".in",
		".भारत",
	},
	languages:    []xlanguage.Tag{xlanguage.Hindi, xlanguage.English},
	currency:     currency.Inr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇮🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIndia) }
