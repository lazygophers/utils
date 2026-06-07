//go:build country_africa || country_all || country_bj || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Benin — Republic of Benin.
var dataBenin = &Country{
	alpha2:       "BJ",
	alpha3:       "BEN",
	numeric:      204,
	callingCodes: []string{"+229"},
	timezones:    []string{"Africa/Porto-Novo"},
	tlds:         []string{".bj"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XOF,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇧🇯",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBenin) }

var Benin = dataBenin
