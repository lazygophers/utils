//go:build country_af || country_all || country_asia || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Afghanistan — Islamic Emirate of Afghanistan.
var dataAfghanistan = &Country{
	alpha2:       "AF",
	alpha3:       "AFG",
	numeric:      4,
	callingCodes: []string{"+93"},
	timezones:    []string{"Asia/Kabul"},
	tlds:         []string{".af"},
	officialLanguage:  xlanguage.MustParse("ps"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("ps"), xlanguage.Persian},
	currency:     currency.AFN,
	region:       RegionSouthernAsia,
	flagEmoji:    "🇦🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAfghanistan) }

var Afghanistan = dataAfghanistan
