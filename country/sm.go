//go:build country_all || country_europe || country_sm || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SanMarino — Republic of San Marino.
var dataSanMarino = &Country{
	alpha2:       "SM",
	alpha3:       "SMR",
	numeric:      674,
	callingCodes: []string{"+378"},
	timezones:    []string{"Europe/San_Marino"},
	tlds:         []string{".sm"},
	languages:    []xlanguage.Tag{xlanguage.Italian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇸🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSanMarino) }

var SanMarino = dataSanMarino
