//go:build country_all || country_asia || country_sy || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Syria — Syrian Arab Republic.
var dataSyria = &Country{
	alpha2:       "SY",
	alpha3:       "SYR",
	numeric:      760,
	callingCodes: []string{"+963"},
	timezones:    []string{"Asia/Damascus"},
	tlds:         []string{
		".sy",
		".سورية",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Syp,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇸🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSyria) }

var Syria = dataSyria
