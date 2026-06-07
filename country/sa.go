//go:build country_all || country_asia || country_sa || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaudiArabia — Kingdom of Saudi Arabia.
var dataSaudiArabia = &Country{
	alpha2:       "SA",
	alpha3:       "SAU",
	numeric:      682,
	callingCodes: []string{"+966"},
	timezones:    []string{"Asia/Riyadh"},
	tlds:         []string{
		".sa",
		".السعودية",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Sar,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇸🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaudiArabia) }

var SaudiArabia = dataSaudiArabia
