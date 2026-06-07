//go:build country_africa || country_all || country_bi || country_eastern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Burundi — Republic of Burundi.
var dataBurundi = &Country{
	alpha2:       "BI",
	alpha3:       "BDI",
	numeric:      108,
	callingCodes: []string{"+257"},
	timezones:    []string{"Africa/Bujumbura"},
	tlds:         []string{".bi"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("rn"), xlanguage.French, xlanguage.English},
	currency:     currency.Bif,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇧🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBurundi) }

var Burundi = dataBurundi
