//go:build country_all || country_antarctic || country_bv

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BouvetIsland — Bouvet Island — uninhabited Norwegian dependency.
var dataBouvetIsland = &Country{
	alpha2:       "BV",
	alpha3:       "BVT",
	numeric:      74,
	callingCodes: []string{},
	timezones:    []string{"Europe/Oslo"},
	tlds:         []string{".bv"},
	officialLanguage:  xlanguage.Norwegian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Norwegian},
	currency:     currency.NOK,
	region:       RegionAntarctic,
	flagEmoji:    "🇧🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBouvetIsland) }

var BouvetIsland = dataBouvetIsland
