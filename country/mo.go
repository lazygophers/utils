//go:build country_all || country_asia || country_eastern_asia || country_mo

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Macao — Macao Special Administrative Region of China.
var dataMacao = &Country{
	alpha2:       "MO",
	alpha3:       "MAC",
	numeric:      446,
	callingCodes: []string{"+853"},
	timezones:    []string{"Asia/Macau"},
	tlds:         []string{
		".mo",
		".澳门",
		".澳門",
	},
	officialLanguage:  xlanguage.MustParse("zh-Hant"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("zh-Hant"), xlanguage.Portuguese},
	currency:     currency.MOP,
	region:       RegionEasternAsia,
	flagEmoji:    "🇲🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMacao) }

var Macao = dataMacao
