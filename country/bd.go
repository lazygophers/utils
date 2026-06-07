//go:build country_all || country_asia || country_bd || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bangladesh — People's Republic of Bangladesh.
var dataBangladesh = &Country{
	alpha2:       "BD",
	alpha3:       "BGD",
	numeric:      50,
	callingCodes: []string{"+880"},
	timezones:    []string{"Asia/Dhaka"},
	tlds:         []string{".bd"},
	languages:    []xlanguage.Tag{xlanguage.Bengali},
	currency:     currency.Bdt,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇧🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBangladesh) }

var Bangladesh = dataBangladesh
