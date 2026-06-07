package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SouthKorea — Republic of Korea.
var dataSouthKorea = &Country{
	alpha2:       "KR",
	alpha3:       "KOR",
	numeric:      410,
	callingCodes: []string{"+82"},
	timezones:    []string{"Asia/Seoul"},
	tlds:         []string{
		".kr",
		".한국",
	},
	languages:    []xlanguage.Tag{xlanguage.Korean},
	currency:     currency.Krw,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Asia",
	flagEmoji:    "🇰🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSouthKorea) }
