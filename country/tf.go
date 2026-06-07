//go:build country_all || country_antarctic || country_tf

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// FrenchSouthernTerritories — French Southern and Antarctic Lands.
var dataFrenchSouthernTerritories = &Country{
	alpha2:       "TF",
	alpha3:       "ATF",
	numeric:      260,
	callingCodes: []string{},
	timezones:    []string{"Indian/Kerguelen"},
	tlds:         []string{".tf"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionAntarctic,
	flagEmoji:    "🇹🇫",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFrenchSouthernTerritories) }

var FrenchSouthernTerritories = dataFrenchSouthernTerritories
