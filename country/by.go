//go:build country_all || country_by || country_eastern_europe || country_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Belarus — Republic of Belarus.
var dataBelarus = &Country{
	alpha2:       "BY",
	alpha3:       "BLR",
	numeric:      112,
	callingCodes: []string{"+375"},
	timezones:    []string{"Europe/Minsk"},
	tlds:         []string{".by"},
	officialLanguage:  xlanguage.Russian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Russian, xlanguage.MustParse("be")},
	currency:     currency.BYN,
	region:       RegionEasternEurope,
	flagEmoji:    "🇧🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBelarus) }

var Belarus = dataBelarus
