package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NorthMacedonia — Republic of North Macedonia.
var dataNorthMacedonia = &Country{
	alpha2:       "MK",
	alpha3:       "MKD",
	numeric:      807,
	callingCodes: []string{"+389"},
	timezones:    []string{"Europe/Skopje"},
	tlds:         []string{".mk"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("mk")},
	currency:     currency.Mkd,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇲🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNorthMacedonia) }
