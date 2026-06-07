//go:build country_africa || country_all || country_eastern_africa || country_ke

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Kenya — Republic of Kenya.
var dataKenya = &Country{
	alpha2:       "KE",
	alpha3:       "KEN",
	numeric:      404,
	callingCodes: []string{"+254"},
	timezones:    []string{"Africa/Nairobi"},
	tlds:         []string{".ke"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("sw"), xlanguage.English},
	currency:     currency.Kes,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇰🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataKenya) }

var Kenya = dataKenya
