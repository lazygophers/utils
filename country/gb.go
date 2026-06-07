package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// UnitedKingdom — United Kingdom of Great Britain and Northern Ireland.
var dataUnitedKingdom = &Country{
	alpha2:       "GB",
	alpha3:       "GBR",
	numeric:      826,
	callingCodes: []string{"+44"},
	timezones:    []string{"Europe/London"},
	tlds:         []string{
		".gb",
		".uk",
	},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("cy"), xlanguage.MustParse("gd")},
	currency:     currency.GBP,
	region:       RegionNorthernEurope,
	flagEmoji:    "🇬🇧",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataUnitedKingdom) }

var UnitedKingdom = dataUnitedKingdom
