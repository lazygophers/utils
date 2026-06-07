package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Qatar — State of Qatar.
var dataQatar = &Country{
	alpha2:       "QA",
	alpha3:       "QAT",
	numeric:      634,
	callingCodes: []string{"+974"},
	timezones:    []string{"Asia/Qatar"},
	tlds:         []string{
		".qa",
		".قطر",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Qar,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇶🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataQatar) }
