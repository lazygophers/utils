//go:build country_ag || country_all || country_americas || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// AntiguaAndBarbuda — Antigua and Barbuda.
var dataAntiguaAndBarbuda = &Country{
	alpha2:       "AG",
	alpha3:       "ATG",
	numeric:      28,
	callingCodes: []string{"+1-268"},
	timezones:    []string{"America/Antigua"},
	tlds:         []string{".ag"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Xcd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇦🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAntiguaAndBarbuda) }

var AntiguaAndBarbuda = dataAntiguaAndBarbuda
