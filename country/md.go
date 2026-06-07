//go:build country_all || country_eastern_europe || country_europe || country_md

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Moldova — Republic of Moldova.
var dataMoldova = &Country{
	alpha2:       "MD",
	alpha3:       "MDA",
	numeric:      498,
	callingCodes: []string{"+373"},
	timezones:    []string{"Europe/Chisinau"},
	tlds:         []string{".md"},
	languages:    []xlanguage.Tag{xlanguage.Romanian},
	currency:     currency.Mdl,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇲🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMoldova) }

var Moldova = dataMoldova
