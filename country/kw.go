//go:build country_all || country_asia || country_kw || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Kuwait — State of Kuwait.
var dataKuwait = &Country{
	alpha2:       "KW",
	alpha3:       "KWT",
	numeric:      414,
	callingCodes: []string{"+965"},
	timezones:    []string{"Asia/Kuwait"},
	tlds:         []string{".kw"},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Kwd,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇰🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataKuwait) }

var Kuwait = dataKuwait
