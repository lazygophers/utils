//go:build country_all || country_asia || country_om || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Oman — Sultanate of Oman.
var dataOman = &Country{
	alpha2:       "OM",
	alpha3:       "OMN",
	numeric:      512,
	callingCodes: []string{"+968"},
	timezones:    []string{"Asia/Muscat"},
	tlds:         []string{
		".om",
		".عمان",
	},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.OMR,
	region:       RegionWesternAsia,
	flagEmoji:    "🇴🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataOman) }

var Oman = dataOman
