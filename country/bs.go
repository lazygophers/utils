//go:build country_all || country_americas || country_bs || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bahamas — Commonwealth of the Bahamas.
var dataBahamas = &Country{
	alpha2:       "BS",
	alpha3:       "BHS",
	numeric:      44,
	callingCodes: []string{"+1-242"},
	timezones:    []string{"America/Nassau"},
	tlds:         []string{".bs"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.BSD,
	region:       RegionCaribbean,
	flagEmoji:    "🇧🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBahamas) }

var Bahamas = dataBahamas
