//go:build country_all || country_oceania || country_polynesia || country_tk

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Tokelau — Tokelau — dependent territory of New Zealand.
var dataTokelau = &Country{
	alpha2:       "TK",
	alpha3:       "TKL",
	numeric:      772,
	callingCodes: []string{"+690"},
	timezones:    []string{"Pacific/Fakaofo"},
	tlds:         []string{".tk"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.NZD,
	region:       RegionPolynesia,
	flagEmoji:    "🇹🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTokelau) }

var Tokelau = dataTokelau
