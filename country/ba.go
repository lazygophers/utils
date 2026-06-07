//go:build country_all || country_ba || country_europe || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BosniaAndHerzegovina — Bosnia and Herzegovina.
var dataBosniaAndHerzegovina = &Country{
	alpha2:       "BA",
	alpha3:       "BIH",
	numeric:      70,
	callingCodes: []string{"+387"},
	timezones:    []string{"Europe/Sarajevo"},
	tlds:         []string{".ba"},
	officialLanguage:  xlanguage.MustParse("bs"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("bs"), xlanguage.Croatian, xlanguage.Serbian},
	currency:     currency.BAM,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇧🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBosniaAndHerzegovina) }

var BosniaAndHerzegovina = dataBosniaAndHerzegovina
