//go:build country_africa || country_all || country_gh || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ghana — Republic of Ghana.
var dataGhana = &Country{
	alpha2:       "GH",
	alpha3:       "GHA",
	numeric:      288,
	callingCodes: []string{"+233"},
	timezones:    []string{"Africa/Accra"},
	tlds:         []string{".gh"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Ghs,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇬🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGhana) }

var Ghana = dataGhana
