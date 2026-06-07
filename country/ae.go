//go:build country_ae || country_all || country_asia || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// UnitedArabEmirates — United Arab Emirates.
var dataUnitedArabEmirates = &Country{
	alpha2:       "AE",
	alpha3:       "ARE",
	numeric:      784,
	callingCodes: []string{"+971"},
	timezones:    []string{"Asia/Dubai"},
	tlds:         []string{
		".ae",
		".امارات",
	},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Aed,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇦🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUnitedArabEmirates) }

var UnitedArabEmirates = dataUnitedArabEmirates
