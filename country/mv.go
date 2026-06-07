//go:build country_all || country_asia || country_mv || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Maldives — Republic of Maldives.
var dataMaldives = &Country{
	alpha2:       "MV",
	alpha3:       "MDV",
	numeric:      462,
	callingCodes: []string{"+960"},
	timezones:    []string{"Indian/Maldives"},
	tlds:         []string{".mv"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("dv")},
	currency:     currency.Mvr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇲🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMaldives) }

var Maldives = dataMaldives
