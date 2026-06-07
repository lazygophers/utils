//go:build country_africa || country_all || country_eastern_africa || country_rw

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Rwanda — Republic of Rwanda.
var dataRwanda = &Country{
	alpha2:       "RW",
	alpha3:       "RWA",
	numeric:      646,
	callingCodes: []string{"+250"},
	timezones:    []string{"Africa/Kigali"},
	tlds:         []string{".rw"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("rw"), xlanguage.English, xlanguage.French},
	currency:     currency.Rwf,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇷🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataRwanda) }

var Rwanda = dataRwanda
