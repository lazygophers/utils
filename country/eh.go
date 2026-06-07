//go:build country_africa || country_all || country_eh || country_northern_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// WesternSahara — Sahrawi Arab Democratic Republic / disputed territory.
var dataWesternSahara = &Country{
	alpha2:       "EH",
	alpha3:       "ESH",
	numeric:      732,
	callingCodes: []string{"+212"},
	timezones:    []string{"Africa/El_Aaiun"},
	tlds:         []string{".eh"},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Mad,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Northern Africa",
	flagEmoji:    "🇪🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataWesternSahara) }

var WesternSahara = dataWesternSahara
