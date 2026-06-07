package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Latvia — Republic of Latvia.
var dataLatvia = &Country{
	alpha2:       "LV",
	alpha3:       "LVA",
	numeric:      428,
	callingCodes: []string{"+371"},
	timezones:    []string{"Europe/Riga"},
	tlds:         []string{".lv"},
	languages:    []xlanguage.Tag{xlanguage.Latvian},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇱🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLatvia) }
