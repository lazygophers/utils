package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// United States of America.
var dataUnitedStates = &Country{
	alpha2:  "US",
	alpha3:  "USA",
	numeric: 840,
	callingCodes: []string{"+1"},
	timezones: []string{
		"America/New_York",
		"America/Chicago",
		"America/Denver",
		"America/Los_Angeles",
		"America/Anchorage",
		"Pacific/Honolulu",
	},
	tlds:      []string{".us"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.Spanish},
	currency:  currency.USD,
	region:       RegionNorthernAmerica,
	flagEmoji: "🇺🇸",
	names:     make(map[xlanguage.Tag]string),
	official:  make(map[xlanguage.Tag]string),
	capital:   make(map[xlanguage.Tag]string),
}

func init() { register(dataUnitedStates) }

var UnitedStates = dataUnitedStates
