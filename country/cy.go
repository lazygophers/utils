package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Cyprus — Republic of Cyprus.
var dataCyprus = &Country{
	alpha2:       "CY",
	alpha3:       "CYP",
	numeric:      196,
	callingCodes: []string{"+357"},
	timezones:    []string{"Asia/Nicosia"},
	tlds:         []string{".cy"},
	languages:    []xlanguage.Tag{xlanguage.Greek, xlanguage.Turkish},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Western Asia",
	flagEmoji:    "🇨🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCyprus) }
