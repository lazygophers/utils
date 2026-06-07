//go:build country_all || country_gu || country_micronesia || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Guam — Territory of Guam.
var dataGuam = &Country{
	alpha2:       "GU",
	alpha3:       "GUM",
	numeric:      316,
	callingCodes: []string{"+1-671"},
	timezones:    []string{"Pacific/Guam"},
	tlds:         []string{".gu"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.USD,
	region:       RegionMicronesia,
	flagEmoji:    "🇬🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataGuam) }

var Guam = dataGuam
