//go:build country_africa || country_all || country_eastern_africa || country_mz

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mozambique — Republic of Mozambique.
var dataMozambique = &Country{
	alpha2:       "MZ",
	alpha3:       "MOZ",
	numeric:      508,
	callingCodes: []string{"+258"},
	timezones:    []string{"Africa/Maputo"},
	tlds:         []string{".mz"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Mzn,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇲🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMozambique) }

var Mozambique = dataMozambique
