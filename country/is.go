//go:build country_all || country_europe || country_is || country_northern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Iceland — Iceland.
var dataIceland = &Country{
	alpha2:       "IS",
	alpha3:       "ISL",
	numeric:      352,
	callingCodes: []string{"+354"},
	timezones:    []string{"Atlantic/Reykjavik"},
	tlds:         []string{".is"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("is")},
	currency:     currency.Isk,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇮🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIceland) }

var Iceland = dataIceland
