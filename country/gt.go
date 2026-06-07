package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guatemala — Republic of Guatemala.
var dataGuatemala = &Country{
	alpha2:       "GT",
	alpha3:       "GTM",
	numeric:      320,
	callingCodes: []string{"+502"},
	timezones:    []string{"America/Guatemala"},
	tlds:         []string{".gt"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Gtq,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Central America",
	flagEmoji:    "🇬🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuatemala) }
