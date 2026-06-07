//go:build country_all || country_asia || country_ge || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Georgia — Georgia.
var dataGeorgia = &Country{
	alpha2:       "GE",
	alpha3:       "GEO",
	numeric:      268,
	callingCodes: []string{"+995"},
	timezones:    []string{"Asia/Tbilisi"},
	tlds:         []string{
		".ge",
		".გე",
	},
	languages:    []xlanguage.Tag{xlanguage.Georgian},
	currency:     currency.Gel,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇬🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGeorgia) }

var Georgia = dataGeorgia
