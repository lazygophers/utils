package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bermuda — British Overseas Territory of Bermuda.
var dataBermuda = &Country{
	alpha2:       "BM",
	alpha3:       "BMU",
	numeric:      60,
	callingCodes: []string{"+1-441"},
	timezones:    []string{"Atlantic/Bermuda"},
	tlds:         []string{".bm"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Bmd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Northern America",
	flagEmoji:    "🇧🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBermuda) }
