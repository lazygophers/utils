//go:build country_all || country_asia || country_western_asia || country_ye

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Yemen — Republic of Yemen.
var dataYemen = &Country{
	alpha2:       "YE",
	alpha3:       "YEM",
	numeric:      887,
	callingCodes: []string{"+967"},
	timezones:    []string{"Asia/Aden"},
	tlds:         []string{".ye"},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.YER,
	region:       RegionWesternAsia,
	flagEmoji:    "🇾🇪",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataYemen) }

var Yemen = dataYemen
