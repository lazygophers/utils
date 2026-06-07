//go:build country_all || country_americas || country_caribbean || country_cu

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Cuba — Republic of Cuba.
var dataCuba = &Country{
	alpha2:       "CU",
	alpha3:       "CUB",
	numeric:      192,
	callingCodes: []string{"+53"},
	timezones:    []string{"America/Havana"},
	tlds:         []string{".cu"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Cup,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇨🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCuba) }

var Cuba = dataCuba
