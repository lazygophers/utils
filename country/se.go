//go:build country_all || country_europe || country_northern_europe || country_se

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Sweden — Kingdom of Sweden.
var dataSweden = &Country{
	alpha2:       "SE",
	alpha3:       "SWE",
	numeric:      752,
	callingCodes: []string{"+46"},
	timezones:    []string{"Europe/Stockholm"},
	tlds:         []string{".se"},
	languages:    []xlanguage.Tag{xlanguage.Swedish},
	currency:     currency.Sek,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇸🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSweden) }

var Sweden = dataSweden
