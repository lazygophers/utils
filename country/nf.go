//go:build country_all || country_australia_and_new_zealand || country_nf || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NorfolkIsland — Norfolk Island — Australian external territory.
var dataNorfolkIsland = &Country{
	alpha2:       "NF",
	alpha3:       "NFK",
	numeric:      574,
	callingCodes: []string{"+672"},
	timezones:    []string{"Pacific/Norfolk"},
	tlds:         []string{".nf"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Aud,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Australia and New Zealand",
	flagEmoji:    "🇳🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNorfolkIsland) }

var NorfolkIsland = dataNorfolkIsland
