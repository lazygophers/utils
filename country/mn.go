//go:build country_all || country_asia || country_eastern_asia || country_mn

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mongolia — Mongolia.
var dataMongolia = &Country{
	alpha2:       "MN",
	alpha3:       "MNG",
	numeric:      496,
	callingCodes: []string{"+976"},
	timezones:    []string{
		"Asia/Ulaanbaatar",
		"Asia/Hovd",
		"Asia/Choibalsan",
	},
	tlds:         []string{".mn"},
	officialLanguage:  xlanguage.MustParse("mn"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.Mongolian, xlanguage.Russian},
	currency:     currency.MNT,
	region:       RegionEasternAsia,
	flagEmoji:    "🇲🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMongolia) }

var Mongolia = dataMongolia
