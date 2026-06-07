//go:build country_all || country_americas || country_central_america || country_mx

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mexico — United Mexican States.
var dataMexico = &Country{
	alpha2:       "MX",
	alpha3:       "MEX",
	numeric:      484,
	callingCodes: []string{"+52"},
	timezones:    []string{
		"America/Mexico_City",
		"America/Cancun",
		"America/Merida",
		"America/Monterrey",
		"America/Mazatlan",
		"America/Chihuahua",
		"America/Hermosillo",
		"America/Tijuana",
	},
	tlds:         []string{".mx"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Mxn,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Central America",
	flagEmoji:    "🇲🇽",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMexico) }

var Mexico = dataMexico
