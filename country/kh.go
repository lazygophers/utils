package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Cambodia — Kingdom of Cambodia.
var dataCambodia = &Country{
	alpha2:       "KH",
	alpha3:       "KHM",
	numeric:      116,
	callingCodes: []string{"+855"},
	timezones:    []string{"Asia/Phnom_Penh"},
	tlds:         []string{".kh"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("km")},
	currency:     currency.Khr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "South-eastern Asia",
	flagEmoji:    "🇰🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCambodia) }
