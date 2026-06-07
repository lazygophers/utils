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
	officialLanguage:  xlanguage.Korean,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Korean},
	currency:     currency.KRW,
	region:       RegionEasternAsia,
	flagEmoji:    "🇰🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSouthKorea) }

var SouthKorea = dataSouthKorea
