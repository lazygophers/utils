package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Pitcairn — Pitcairn Islands — British Overseas Territory.
var dataPitcairn = &Country{
	alpha2:       "PN",
	alpha3:       "PCN",
	numeric:      612,
	callingCodes: []string{"+64"},
	timezones:    []string{"Pacific/Pitcairn"},
	tlds:         []string{".pn"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Nzd,
	continent:    "OC",
	region:       "Oceania",
	subregion:    "Polynesia",
	flagEmoji:    "🇵🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPitcairn) }
