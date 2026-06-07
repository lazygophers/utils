package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Turkmenistan — Turkmenistan.
var dataTurkmenistan = &Country{
	alpha2:       "TM",
	alpha3:       "TKM",
	numeric:      795,
	callingCodes: []string{"+993"},
	timezones:    []string{"Asia/Ashgabat"},
	tlds:         []string{".tm"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("tk")},
	currency:     currency.Tmt,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Central Asia",
	flagEmoji:    "🇹🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTurkmenistan) }
