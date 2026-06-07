package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// FaroeIslands — Faroe Islands — autonomous region of Denmark.
var dataFaroeIslands = &Country{
	alpha2:       "FO",
	alpha3:       "FRO",
	numeric:      234,
	callingCodes: []string{"+298"},
	timezones:    []string{"Atlantic/Faroe"},
	tlds:         []string{".fo"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("fo"), xlanguage.Danish},
	currency:     currency.Dkk,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇫🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataFaroeIslands) }
