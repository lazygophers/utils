//go:build country_all || country_asia || country_lb || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Lebanon — Lebanese Republic.
var dataLebanon = &Country{
	alpha2:       "LB",
	alpha3:       "LBN",
	numeric:      422,
	callingCodes: []string{"+961"},
	timezones:    []string{"Asia/Beirut"},
	tlds:         []string{".lb"},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic, xlanguage.French, xlanguage.English},
	currency:     currency.LBP,
	region:       RegionWesternAsia,
	flagEmoji:    "🇱🇧",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataLebanon) }

var Lebanon = dataLebanon
