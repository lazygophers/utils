//go:build country_all || country_europe || country_pt || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Portugal — Portuguese Republic.
var dataPortugal = &Country{
	alpha2:       "PT",
	alpha3:       "PRT",
	numeric:      620,
	callingCodes: []string{"+351"},
	timezones:    []string{
		"Europe/Lisbon",
		"Atlantic/Madeira",
		"Atlantic/Azores",
	},
	tlds:         []string{".pt"},
	officialLanguage:  xlanguage.Portuguese,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.EUR,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇵🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPortugal) }

var Portugal = dataPortugal
