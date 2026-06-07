//go:build country_all || country_antarctic || country_aq

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Antarctica — Antarctica — territories south of 60°S.
var dataAntarctica = &Country{
	alpha2:       "AQ",
	alpha3:       "ATA",
	numeric:      10,
	callingCodes: []string{},
	timezones:    []string{
		"Antarctica/McMurdo",
		"Antarctica/Casey",
		"Antarctica/Davis",
		"Antarctica/Mawson",
		"Antarctica/Palmer",
		"Antarctica/Rothera",
		"Antarctica/Syowa",
		"Antarctica/Troll",
		"Antarctica/Vostok",
	},
	tlds:         []string{".aq"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.USD,
	region:       RegionAntarctic,
	flagEmoji:    "🇦🇶",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAntarctica) }

var Antarctica = dataAntarctica
