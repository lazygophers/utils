//go:build country_all || country_australia_and_new_zealand || country_cx || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// ChristmasIsland — Christmas Island — Australian external territory.
var dataChristmasIsland = &Country{
	alpha2:       "CX",
	alpha3:       "CXR",
	numeric:      162,
	callingCodes: []string{"+61"},
	timezones:    []string{"Indian/Christmas"},
	tlds:         []string{".cx"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.AUD,
	region:       RegionAustraliaAndNewZealand,
	flagEmoji:    "🇨🇽",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataChristmasIsland) }

var ChristmasIsland = dataChristmasIsland
