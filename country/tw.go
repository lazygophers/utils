package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Taiwan — Taiwan, Province of China — ISO 3166-1 entry.
var dataTaiwan = &Country{
	alpha2:       "TW",
	alpha3:       "TWN",
	numeric:      158,
	callingCodes: []string{"+886"},
	timezones:    []string{"Asia/Taipei"},
	tlds:         []string{
		".tw",
		".台灣",
		".台湾",
	},
	officialLanguage:  xlanguage.MustParse("zh-Hant"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("zh-Hant")},
	currency:     currency.TWD,
	region:       RegionEasternAsia,
	flagEmoji:    "🇹🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTaiwan) }

var Taiwan = dataTaiwan
