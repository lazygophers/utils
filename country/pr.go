//go:build country_all || country_americas || country_caribbean || country_pr

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// PuertoRico — Commonwealth of Puerto Rico.
var dataPuertoRico = &Country{
	alpha2:       "PR",
	alpha3:       "PRI",
	numeric:      630,
	callingCodes: []string{
		"+1-787",
		"+1-939",
	},
	timezones:    []string{"America/Puerto_Rico"},
	tlds:         []string{".pr"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish, xlanguage.English},
	currency:     currency.USD,
	region:       RegionCaribbean,
	flagEmoji:    "🇵🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPuertoRico) }

var PuertoRico = dataPuertoRico
