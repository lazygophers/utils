//go:build country_all || country_asia || country_tr || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Turkey — Republic of Türkiye.
var dataTurkey = &Country{
	alpha2:       "TR",
	alpha3:       "TUR",
	numeric:      792,
	callingCodes: []string{"+90"},
	timezones:    []string{"Europe/Istanbul"},
	tlds:         []string{".tr"},
	officialLanguage:  xlanguage.Turkish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Turkish},
	currency:     currency.TRY,
	region:       RegionWesternAsia,
	flagEmoji:    "🇹🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTurkey) }

var Turkey = dataTurkey
