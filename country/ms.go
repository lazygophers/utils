//go:build country_all || country_americas || country_caribbean || country_ms

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Montserrat — Montserrat — British Overseas Territory.
var dataMontserrat = &Country{
	alpha2:       "MS",
	alpha3:       "MSR",
	numeric:      500,
	callingCodes: []string{"+1-664"},
	timezones:    []string{"America/Montserrat"},
	tlds:         []string{".ms"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.XCD,
	region:       RegionCaribbean,
	flagEmoji:    "🇲🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMontserrat) }

var Montserrat = dataMontserrat
