package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bulgaria — Republic of Bulgaria.
var dataBulgaria = &Country{
	alpha2:       "BG",
	alpha3:       "BGR",
	numeric:      100,
	callingCodes: []string{"+359"},
	timezones:    []string{"Europe/Sofia"},
	tlds:         []string{".bg"},
	languages:    []xlanguage.Tag{xlanguage.Bulgarian},
	currency:     currency.Bgn,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Eastern Europe",
	flagEmoji:    "🇧🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBulgaria) }
