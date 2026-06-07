//go:build country_all || country_americas || country_central_america || country_pa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Panama — Republic of Panama.
var dataPanama = &Country{
	alpha2:       "PA",
	alpha3:       "PAN",
	numeric:      591,
	callingCodes: []string{"+507"},
	timezones:    []string{"America/Panama"},
	tlds:         []string{".pa"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.PAB,
	region:       RegionCentralAmerica,
	flagEmoji:    "🇵🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPanama) }

var Panama = dataPanama
