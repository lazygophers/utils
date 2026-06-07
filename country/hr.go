//go:build country_all || country_europe || country_hr || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Croatia — Republic of Croatia.
var dataCroatia = &Country{
	alpha2:       "HR",
	alpha3:       "HRV",
	numeric:      191,
	callingCodes: []string{"+385"},
	timezones:    []string{"Europe/Zagreb"},
	tlds:         []string{".hr"},
	languages:    []xlanguage.Tag{xlanguage.Croatian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇭🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCroatia) }

var Croatia = dataCroatia
