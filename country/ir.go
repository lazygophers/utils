//go:build country_all || country_asia || country_ir || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Iran — Islamic Republic of Iran.
var dataIran = &Country{
	alpha2:       "IR",
	alpha3:       "IRN",
	numeric:      364,
	callingCodes: []string{"+98"},
	timezones:    []string{"Asia/Tehran"},
	tlds:         []string{
		".ir",
		".ایران",
	},
	languages:    []xlanguage.Tag{xlanguage.Persian},
	currency:     currency.Irr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇮🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIran) }

var Iran = dataIran
