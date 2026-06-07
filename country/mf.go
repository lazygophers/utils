//go:build country_all || country_americas || country_caribbean || country_mf

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintMartin — Collectivity of Saint Martin (French part).
var dataSaintMartin = &Country{
	alpha2:       "MF",
	alpha3:       "MAF",
	numeric:      663,
	callingCodes: []string{"+590"},
	timezones:    []string{"America/Marigot"},
	tlds:         []string{".mf"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Eur,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇲🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintMartin) }

var SaintMartin = dataSaintMartin
