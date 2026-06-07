//go:build country_all || country_asia || country_my || country_south_eastern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Malaysia — Malaysia.
var dataMalaysia = &Country{
	alpha2:       "MY",
	alpha3:       "MYS",
	numeric:      458,
	callingCodes: []string{"+60"},
	timezones:    []string{
		"Asia/Kuala_Lumpur",
		"Asia/Kuching",
	},
	tlds:         []string{".my"},
	officialLanguage:  xlanguage.Malay,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Malay, xlanguage.English, xlanguage.Chinese, xlanguage.MustParse("ta")},
	currency:     currency.MYR,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇲🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMalaysia) }

var Malaysia = dataMalaysia
