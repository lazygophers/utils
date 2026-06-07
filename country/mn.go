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
	languages:    []xlanguage.Tag{xlanguage.MustParse("mn")},
	currency:     currency.Mnt,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Asia",
	flagEmoji:    "🇲🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMongolia) }
