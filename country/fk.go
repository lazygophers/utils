//go:build country_all || country_americas || country_fk || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// FalklandIslands — Falkland Islands (Malvinas) — British Overseas Territory.
var dataFalklandIslands = &Country{
	alpha2:       "FK",
	alpha3:       "FLK",
	numeric:      238,
	callingCodes: []string{"+500"},
	timezones:    []string{"Atlantic/Stanley"},
	tlds:         []string{".fk"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Fkp,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇫🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFalklandIslands) }

var FalklandIslands = dataFalklandIslands
