//go:build country_all || country_asia || country_np || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Nepal — Federal Democratic Republic of Nepal.
var dataNepal = &Country{
	alpha2:       "NP",
	alpha3:       "NPL",
	numeric:      524,
	callingCodes: []string{"+977"},
	timezones:    []string{"Asia/Kathmandu"},
	tlds:         []string{".np"},
	officialLanguage:  xlanguage.MustParse("ne"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("ne"), xlanguage.English},
	currency:     currency.NPR,
	region:       RegionSouthernAsia,
	flagEmoji:    "🇳🇵",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNepal) }

var Nepal = dataNepal
